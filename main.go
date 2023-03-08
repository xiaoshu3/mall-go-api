package main

import (
	"flag"
	"fmt"
	"mall/app/http/controllers/api/v1/middlewares"
	"mall/bootstrap"
	"mall/pkg/config"
	"mall/pkg/response"

	"mall/pkg/auth"

	"github.com/gin-gonic/gin"

	btsConfig "mall/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	gin.SetMode(gin.ReleaseMode)

	// 初始化 Logger
	bootstrap.SetupLogger()
	// new 一个 Gin Engine 实例
	router := gin.New()

	bootstrap.SetupDB()

	// 初始化 redis
	bootstrap.SetupRedis()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// redis.Redis.Set("key", "ABCD", time.Minute*10)

	// logger.DebugString("redis_test", "key", redis.Redis.Get("key"))
	// 运行服务

	router.GET("/test_auth", middlewares.AuthJWT(), func(c *gin.Context) {
		userModel := auth.CurrentUser(c)
		response.Data(c, userModel)
	})
	err := router.Run(":8010")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
