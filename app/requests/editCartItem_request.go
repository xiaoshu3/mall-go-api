package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type EditCartItemRequest struct {
	CartId     uint64 `valid:"cart_id" json:"cart_id"`
	GoodsCount uint   `valid:"goods_count" json:"goods_count"`
}

func EditCartItem(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"cart_id":     []string{"required", "min:0"},
		"goods_count": []string{"required", "min:0", "max:99"},
	}
	messages := govalidator.MapData{
		"cart_id": []string{
			"required:购物车Id为必填项",
			"min:购物车Id错误,请重新登录",
		},
		"goods_count": []string{
			"required:添加数目为必填项",
			"min:添加数目需大于等于1",
			"max:添加数目超过最大值99",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}
