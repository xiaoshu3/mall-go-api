package address

import (
	"mall/app/models"
	"mall/pkg/database"
)

type Address struct {
	models.BaseModel
	UserID uint `json:"user_id" gorm:"index;not null;"`

	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`

	ProvinceName  string `json:"province_name" gorm:"size:32"`
	CityName      string `json:"city_name" gorm:"size:30"`
	RegionName    string `json:"region_name" gorm:"size:32"`
	DetailAddress string `json:"detail_address" gorm:"size:64"`
	DefaultFlag   bool   `json:"default_flag"`

	models.CommonTimestampsField
}

func (addreddModel *Address) Create() {
	database.DB.Create(&addreddModel)
}

func (addreddModel *Address) Save() (rowsAffected int64) {
	result := database.DB.Save(&addreddModel)
	return result.RowsAffected
}

func (addreddModel *Address) Delete() error {
	return database.DB.Delete(&addreddModel).Error
}
