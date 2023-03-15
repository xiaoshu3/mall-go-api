package routes

import (
	"mall/app/http/controllers/api/v1/auth"
	"mall/app/http/controllers/api/v1/category"
	"mall/app/http/controllers/api/v1/home"
	"mall/app/http/controllers/api/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里

	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")

		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)

			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)

			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.GuestJWT(), lgc.RefreshToken)
		}
	}

	// 首页
	{
		hc := new(home.HomeController)
		homeGroup := v1.Group("/home")
		homeGroup.GET("/carousels", hc.GetCarousels)
		homeGroup.GET("/grids", hc.GetGrids)
	}

	// 分类页
	{
		cc := new(category.CategoryController)
		// categoryGroup := v1.Group("/categorys")
		// categoryGroup.GET("/spec", cc.GetALlCategorys)
		v1.GET("/category", cc.GetALlCategorys)
	}
}
