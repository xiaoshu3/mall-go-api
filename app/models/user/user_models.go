package user

import (
	"mall/app/models"
	"mall/app/models/address"
	"mall/app/models/cart"
	"mall/pkg/database"
	"mall/pkg/hash"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	Carts   []*cart.Cart       `json:"carts"`
	Address []*address.Address `json:"address"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
