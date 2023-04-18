package routes

import (
	"mall/app/http/controllers/api/v1/auth"
	"mall/app/http/controllers/api/v1/cart"
	"mall/app/http/controllers/api/v1/category"
	"mall/app/http/controllers/api/v1/home"
	"mall/app/http/controllers/api/v1/middlewares"
	"mall/app/http/controllers/api/v1/spu"
	"mall/app/http/controllers/api/v1/user"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里

	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("300-H"))

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

	// 商品接口
	{
		sc := new(spu.SpuController)
		v1.GET("/products/:id", sc.GetSpuById)

		v1.GET("/spus", sc.GetSpus)
	}

	// 用户接口
	{
		uc := new(user.UsersController)

		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)

		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", middlewares.AuthJWT(), uc.Index)
		}

	}

	// 购物车相关接口
	{
		cc := new(cart.CartController)
		v1.POST("/shop-cart", middlewares.AuthJWT(), cc.AddGoodsToCart)

		v1.GET("/shop-cart", middlewares.AuthJWT(), cc.GetCartList)
		v1.GET("/shop-carts", middlewares.AuthJWT(), cc.GetCartListByPreload)

		v1.DELETE("/shop-cart", middlewares.AuthJWT(), cc.DeleteCartItem)
		v1.PUT("/shop-cart", middlewares.AuthJWT(), cc.EditCartItem)
	}

}
