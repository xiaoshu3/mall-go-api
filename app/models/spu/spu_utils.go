package spu

import (
	"mall/pkg/database"
)

func Get(idstr string) (spuModel SPU) {
	// database.DB.Select("carousels", "title", "sub_title", "price", "goods_detail_content").
	// 	Where("id", idstr).Where("saleable = ? ", 1).Where("valid = ? ", 1).First(&spuModel)
	// database.DB.Where("id", idstr).First(&spuModel)
	database.DB.Where("id", idstr).Where("saleable = ? ", 1).Where("valid = ? ", 1).
		Select("id", "carousels", "title", "sub_title", "price", "goods_detail_content").First(&spuModel)

	// database.DB.Where("id", idstr).Where("saleable = ? ", 1).Where("valid = ? ", 1).
	// 	Select("title", "id").First(&spuModel)
	return
}

func GetHtml(idstr string) (spuModel SPU) {
	database.DB.Select("id", "goods_detail_content").Where("id", idstr).Where("saleable = ? ", "1").
		Where("valid = ? ", "1").First(&spuModel)
	// spuModel.GoodsDetailContent.UnmarshalJSON()
	return
}

func IsExistById(id uint64) bool {
	var count int64
	database.DB.Model(SPU{}).Where("id = ?", id).Where("saleable = ? ", 1).Where("valid = ? ", 1).Count(&count)
	return count > 0
}
