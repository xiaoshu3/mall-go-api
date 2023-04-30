package address

import (
	"errors"
	v1 "mall/app/http/controllers/api/v1"
	"mall/app/models"
	"mall/app/models/address"
	"mall/app/requests"
	"mall/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type AddressController struct {
	v1.BaseAPIController
}

func (ac *AddressController) AddAddress(c *gin.Context) {
	userID, ok := c.Get("current_user_id")
	if !ok {
		response.Unauthorized(c)
		return
	}
	id := cast.ToUint(userID)

	// 1. 验证表单
	request := requests.AddOrEditAddressRequest{}
	if ok := requests.Validate(c, &request, requests.AddOrEditAddress); !ok {
		return
	}

	// 2. 验证成功，创建数据
	address := address.Address{
		UserID:        id,
		Name:          request.Name,
		Phone:         request.Phone,
		ProvinceName:  request.ProvinceName,
		CityName:      request.CityName,
		RegionName:    request.RegionName,
		DetailAddress: request.DetailAddress,
		DefaultFlag:   request.DefaultFlag,
	}

	address.Create()

	if address.ID > 0 {
		response.Created(c, address)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}

}

func (ac *AddressController) DeleteAddress(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Abort404(c)
	}
	data := address.GetAddressById(id)

	// if data.ID > 0 {
	// 	response.Data(c, data)
	// } else {
	// 	response.Abort404(c)
	// }

	if data.ID == 0 {
		response.Abort404(c)
	} else {
		if err := data.Delete(); err != nil {
			response.Error(c, err)
		} else {
			response.Success(c)
		}
	}

}

func (ac *AddressController) EditAddress(c *gin.Context) {
	userID, ok := c.Get("current_user_id")
	if !ok {
		response.Unauthorized(c)
	}
	id := cast.ToUint(userID)

	// 1. 验证表单
	request := requests.AddOrEditAddressRequest{}
	if ok := requests.Validate(c, &request, requests.AddOrEditAddress); !ok {
		return
	}

	if request.ID == 0 {
		response.Error(c, errors.New("参数错误"))
		return
	}

	address := address.Address{
		BaseModel:     models.BaseModel{ID: request.ID},
		UserID:        id,
		Name:          request.Name,
		Phone:         request.Phone,
		ProvinceName:  request.ProvinceName,
		CityName:      request.CityName,
		RegionName:    request.RegionName,
		DetailAddress: request.DetailAddress,
		DefaultFlag:   request.DefaultFlag,
	}

	affectedRpw := address.Save()
	if affectedRpw <= 0 {
		response.Error(c, errors.New("参数错误"))
		return
	} else {
		response.Data(c, address)
	}
}

func (ac *AddressController) GetDefaultAddress(c *gin.Context) {
	userID, ok := c.Get("current_user_id")
	if !ok {
		response.Unauthorized(c)
	}
	userId := cast.ToString(userID)

	data := address.GetDefaultAddress(userId)
	// if data.ID == 0 {
	// 	response.Abort404(c)
	// } else {
	// 	response.Data(c, data)
	// }
	response.Data(c, data)

}

func (ac *AddressController) GetAddressList(c *gin.Context) {
	userID, ok := c.Get("current_user_id")
	if !ok {
		response.Unauthorized(c)
	}
	userId := cast.ToString(userID)

	datas := address.GetAddressList(userId)
	response.Data(c, datas)
}

func (ac *AddressController) GetAddressByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Abort404(c)
	}
	data := address.GetAddressById(id)

	if data.ID == 0 {
		response.Abort404(c)
	} else {
		response.Data(c, data)
	}
}
