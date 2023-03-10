package models

import (
	"time"

	"github.com/spf13/cast"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	// ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index" json:"create_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:update_at;index" json:"update_at,omitempty"`
}

func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
