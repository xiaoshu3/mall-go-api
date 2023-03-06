package user

import "mall/app/models"

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	Phone string `json:"-"`
	Password string 

	models.CommonTimestampsField
}

