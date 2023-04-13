//产品表
package spu

import (
	"mall/app/models"
	
)

type SPU struct {
	models.BaseModel
	// Carousels datatypes.JSON `json:"carousels" gorm:"comment:轮播图"`
	GoodsCoverImg string `json:"goods_cover_img,omitempty" gorm:"200;comment:商品封面图;"`
	Carousels     string `json:"carousels,omitempty" gorm:"comment:轮播图;size:2000"`
	Title         string `json:"title,,omitempty" gorm:"not null;size:200;comment:标题"`
	SubTitle      string `json:"subTitle,omitempty" gorm:"size:200;comment:副标题"`
	// CategoryId uint   `json:"categoryId" gorm:"not null;index;comment:分类ID"`
	BrandId uint `json:"brandId,omitempty" gorm:"index;comment:品牌ID"`
	// SpecGroupId uint       `json:"spgId" gorm:"not null;index;comment:品类ID"`
	SpecCategoryId uint `json:"categoryId,omitempty" gorm:"not null;index;comment:分类ID"`
	Price          uint `json:"price,omitempty" gorm:"size:10;index;comment:商品价格"`
	// GoodsDetailContent datatypes.JSON `json:"goods_detail_content" gorm:"not null;comment:商品参数"`
	GoodsDetailContent string `json:"goods_detail_content,omitempty" gorm:";comment:商品参数"`
	Saleable           bool   `json:"saleable,omitempty" gorm:"not null;index;default:1;comment:是否上架"`
	Valid              bool   `json:"valid,omitempty" gorm:"not null;index;default:1;comment:是否有效"`
	// Skus               []*sku.SKU     `json:"skus"`
	models.CommonTimestampsField
}


