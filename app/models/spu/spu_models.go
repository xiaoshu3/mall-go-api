//产品表
package spu

import (
	"mall/app/models"
	"mall/pkg/database"
)

type SPU struct {
	models.BaseModel
	// Carousels datatypes.JSON `json:"carousels" gorm:"comment:轮播图"`
	GoodsCoverImg string `json:"goods_cover_img" gorm:"200;comment:商品封面图;"`
	Carousels     string `json:"carousels" gorm:"comment:轮播图;size:2000"`
	Title         string `json:"title" gorm:"not null;size:200;comment:标题"`
	SubTitle      string `json:"subTitle" gorm:"size:200;comment:副标题"`
	// CategoryId uint   `json:"categoryId" gorm:"not null;index;comment:分类ID"`
	BrandId uint `json:"brandId" gorm:"index;comment:品牌ID"`
	// SpecGroupId uint       `json:"spgId" gorm:"not null;index;comment:品类ID"`
	SpecCategoryId uint `json:"categoryId" gorm:"not null;index;comment:分类ID"`
	Price          uint `json:"price" gorm:"size:10;index;comment:商品价格"`
	// GoodsDetailContent datatypes.JSON `json:"goods_detail_content" gorm:"not null;comment:商品参数"`
	GoodsDetailContent string `json:"goods_detail_content" gorm:";comment:商品参数"`
	Saleable           bool   `json:"saleable" gorm:"not null;index;default:1;comment:是否上架"`
	Valid              bool   `json:"valid" gorm:"not null;index;default:1;comment:是否有效"`
	// Skus               []*sku.SKU     `json:"skus"`
	models.CommonTimestampsField
}

func Get(idstr string) (spuModel SPU) {
	// database.DB.Select("carousels", "title", "sub_title", "price", "goods_detail_content").
	// 	Where("id", idstr).Where("saleable = ? ", 1).Where("valid = ? ", 1).First(&spuModel)
	// database.DB.Where("id", idstr).First(&spuModel)
	database.DB.Where("id", idstr).Where("saleable = ? ", 1).Where("valid = ? ", 1).
		Select("id","carousels", "title", "sub_title", "price", "goods_detail_content").First(&spuModel)

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
