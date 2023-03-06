package user

import "mall/pkg/database"

func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?",phone).Count(&count)
	return count > 0
}