package address

import "mall/pkg/database"

func GetAddressById(idstr string) (addressModel Address) {
	database.DB.Where("id = ?", idstr).Find(&addressModel)
	return
}

func GetDefaultAddress(userID string) (addressModel Address) {
	database.DB.Where("user_id = ?", userID).Where("default_flag =?", 1).First(&addressModel)
	return
}

func GetAddressList(userID string) (addressModels []Address) {
	database.DB.Where("user_id = ?", userID).Find(&addressModels)
	return
}
