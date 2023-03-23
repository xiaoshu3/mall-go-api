package address

import "mall/app/models"

type Address struct {
	models.BaseModel
	UserID        uint `json:"user_id" gorm:"index;not null;"`
	ProvinceName  string `json:"province_name" gorm:"size:32"`
	CityName      string `json:"city_name" gorm:"size:30"`   
	RegionName    string `json:"region_name" gorm:"size:32"`
	DetailAddress string `json:"detail_address" gorm:"size:64"`
	DefaultFlag bool `json:"default_flag"`

	models.CommonTimestampsField
}
