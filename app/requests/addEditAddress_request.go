package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AddOrEditAddressRequest struct {
	// UserID        uint   `json:"user_id" gorm:"index;not null;" valid:"user_id"`
	ID            uint64 `json:"id,omitempty" valid:"id"`
	ProvinceName  string `json:"province_name" gorm:"size:32" valid:"province_name"`
	CityName      string `json:"city_name" gorm:"size:30" valid:"city_name"`
	RegionName    string `json:"region_name" gorm:"size:32" valid:"region_name"`
	DetailAddress string `json:"detail_address" gorm:"size:64" valid:"detail_address"`
	DefaultFlag   bool   `json:"default_flag" valid:"default_flag"`
}

func AddOrEditAddress(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"province_name":  []string{"required", "max:20"},
		"city_name":      []string{"required", "max:20"},
		"region_name":    []string{"required", "max:20"},
		"detail_address": []string{"required", "max:64"},
		// "default_flag":   []string{"required"},  // panic for bool tag required
		"default_flag": []string{"bool"},
	}

	messages := govalidator.MapData{
		"province_name": []string{
			"required:省名为必填项",
			"max:输入过长,请检查",
		},
		"city_name": []string{
			"required:市名为必填项",
			"max:输入过长,请检查",
		},
		"region_name": []string{
			"required:地区为必填项",
			"max:输入过长,请检查",
		},
		"detail_address": []string{
			"required:详细地址为必填项",
			"max:输入过长,请检查",
		},
		"default_flag": []string{
			"bool:缺少是否默认地址,请填入正确的参数",
			// "max:输入过长,请检查",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}
