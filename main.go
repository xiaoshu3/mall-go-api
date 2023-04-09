package main

import (
	"fmt"
	"mall/app/cmd"
	"mall/bootstrap"
	"mall/pkg/config"
	"mall/pkg/console"
	"os"

	"github.com/spf13/cobra"

	btsConfig "mall/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	// var env string
	// flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	// flag.Parse()
	// config.InitConfig(env)

	// gin.SetMode(gin.ReleaseMode)

	// // 初始化 Logger
	// bootstrap.SetupLogger()
	// // new 一个 Gin Engine 实例
	// router := gin.New()

	// bootstrap.SetupDB()

	// // 初始化 redis
	// bootstrap.SetupRedis()

	// // 初始化路由绑定
	// bootstrap.SetupRoute(router)

	// redis.Redis.Set("key", "ABCD", time.Minute*10)

	// logger.DebugString("redis_test", "key", redis.Redis.Get("key"))
	// 运行服务

	// router.GET("/test_auth", middlewares.AuthJWT(), func(c *gin.Context) {
	// 	userModel := auth.CurrentUser(c)
	// 	response.Data(c, userModel)
	// })

	// err := router.Run(":8010")
	// if err != nil {
	// 	// 错误处理，端口被占用了或者其他错误
	// 	fmt.Println(err.Error())
	// }

	var rootCmd = &cobra.Command{
		Use:   "Mall",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
