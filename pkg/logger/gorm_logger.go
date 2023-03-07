package logger

import (
	"context"
	"errors"
	"mall/pkg/helpers"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger:     Logger,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Debugf(str, args)
}

func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Warnf(str, args)

}

func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Errorf(str, args)

}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elasped := time.Since(begin)

	sql, rows := fc()

	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", helpers.MicrosecondStr(elasped)),
		zap.Int64("row", rows),
	}

	if err != nil {
		if errors.Is(err, gormlogger.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound", logFields...)
		} else {
			logFields = append(logFields, zap.Error(err))
			// l.logger.Error("Database error", logFields...)
			l.logger().Error("Database error", logFields...)
		}
	}

	if l.SlowThreshold != 0 && elasped > l.SlowThreshold {
		l.logger().Warn("Database Slow Log", logFields...)
	}

	l.logger().Debug("Database Query", logFields...)
}

// logger 内用的辅助方法，确保 Zap 内置信息 Caller 的准确性（如 paginator/paginator.go:148）
func (l GormLogger) logger() *zap.Logger {
	// 跳过 gorm 内置的调用
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	gormPackage = "gorm.io/gorm"
	zapgormPackage = "moul.io/zapgorm2"
	// 减去一次封装，以及一次在 logger 初始化里添加 zap.AddCallerSkip(1)
	// clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-1))
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		// if ok {
		// 	//  filename    {"file":
		// 	//"C:/Users/xiaoshu/go/pkg/mod/gorm.io/gorm@v1.24.2/migrator/migrator.go"}
		// 	Info("filename", zap.String("file", file))
		// }
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// 返回一个附带跳过行号的新的 zap logger
			// Info("OK!! " + file)
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}
