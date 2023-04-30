package cart

import (
	"errors"
	v1 "mall/app/http/controllers/api/v1"
	"mall/app/models"
	"mall/app/models/cart"
	"mall/app/models/spu"
	"mall/app/requests"
	"mall/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type CartController struct {
	v1.BaseAPIController
}

func (cc *CartController) AddGoodsToCart(c *gin.Context) {
	userID, ok := c.Get("current_user_id")
	if !ok {
		response.Unauthorized(c)
	}

	// 1. 验证表单
	request := requests.AddCartGoodsRequest{}
	if ok := requests.Validate(c, &request, requests.AddCartGoods); !ok {
		return
	}

	//判断商品
	if ok = spu.IsExistById(request.SpuId); !ok {
		response.BadRequest(c, errors.New("商品不存在"))
		return
	}

	// 2. 验证成功，创建数据
	cartModel := cart.Cart{
		UserID: cast.ToUint64(userID),
		SpuId:  cast.ToUint64(request.SpuId),
		// GoodsCount: cast.ToUint(request.AddNum),
	}

	// logger.Dump(cartModel)

	// 3.查看记录是否存在
	ok = cartModel.IsRecordExist()

	// logger.Dump(cartModel)
	if ok {
		cartModel.GoodsCount += request.AddNum
		if cartModel.GoodsCount > 99 {
			response.BadRequest(c, errors.New("商品添加数量超过最大值"))
			return
		}
		rowAffect := cartModel.Save()
		if rowAffect > 0 {
			response.Data(c, cartModel)
		} else {
			response.Abort500(c, "添加失败，请稍后尝试~")
		}
	} else {
		cartModel.GoodsCount = request.AddNum
		cartModel.Create()
		if cartModel.ID > 0 {
			response.CreatedJSON(c, gin.H{
				"data": cartModel,
			})
		} else {
			response.Abort500(c, "添加失败，请稍后尝试~")
		}
	}
}

func (cc *CartController) GetCartList(c *gin.Context) {
	userID, ok := c.Get("current_user_id")
	if !ok {
		response.Unauthorized(c)
	}

	id := cast.ToString(userID)
	var cartList []cart.CartList
	cart.GetCartList(id, &cartList)

	response.Data(c, cartList)
}

func (cc *CartController) GetCartListByPreload(c *gin.Context) {
	userID, ok := c.Get("current_user_id")
	if !ok {
		response.Unauthorized(c)
	}

	id := cast.ToString(userID)
	var cartList []cart.Cart
	cart.GetCartListByPreload(id, &cartList)
	response.Data(c, cartList)
}

type DeleteCartItemRequset struct {
	CartItem uint64 `json:"cart_item" binding:"required"`
}

func (cc *CartController) DeleteCartItem(c *gin.Context) {
	var request DeleteCartItemRequset
	if err := c.ShouldBind(&request); err != nil {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		// 	"error":   err.Error(),
		// })
		// fmt.Println(err.Error())
		response.Error(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}
	cartModel := cart.Cart{
		BaseModel: models.BaseModel{ID: request.CartItem},
	}

	// 查找记录是否存在
	if ok := cart.IsExist("id", cast.ToString(request.CartItem)); !ok {
		response.Abort404(c)
		return
	}
	if err := cartModel.Delete(); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c)
	}

}

func (cc *CartController) EditCartItem(c *gin.Context) {
	// 1. 验证表单
	request := requests.EditCartItemRequest{}
	if ok := requests.Validate(c, &request, requests.EditCartItem); !ok {
		return
	}

	// 2 查找是否存在记录
	if ok := cart.IsExist("id", cast.ToString(request.CartId)); !ok {
		response.Abort404(c)
		return
	}

	// 3 更新记录
	cart.UpdateGoodsCount(request.CartId, request.GoodsCount)

	response.Success(c)
}

type GetCartListByIdsRequest struct {
	CartItemIds string `json:"cartItemIds" binding:"required"`
}

func (cc *CartController) GetCartListByIds(c *gin.Context) {
	// 1. 验证表单

	request := GetCartListByIdsRequest{}
	if err := c.ShouldBind(&request); err != nil {
		response.Error(c, err)
		return
	}
	ids := strings.Split(request.CartItemIds, ",")

	// err := c.ShouldBindQuery(&request)
	// if err != nil {
	// 	response.Error(c, err)
	// 	return
	// }

	// ids := strings.Split(request.CartItemIds, ",")
	// logger.Dump(ids, "IDS===")
	// logger.Dump(idsParam, "IDSParm")

	var cartList []cart.Cart
	cart.GetCartListByIds(ids, &cartList)
	response.Data(c, cartList)
}
