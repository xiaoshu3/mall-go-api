package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AddCartGoodsRequest struct {
	SpuId  uint64 `valid:"spu_id" json:"spu_id"`
	AddNum uint   `valid:"add_num" json:"add_num"`
}

func AddCartGoods(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"spu_id":  []string{"required", "min:0"},
		"add_num": []string{"required", "min:0"},
	}
	messages := govalidator.MapData{
		"spu_id": []string{
			"required:商品Id为必填项",
			"min:用户Id错误,请重新登录",
		},
		"add_num": []string{
			"required:添加数目为必填项",
			"min:添加数目需大于等于1",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}
