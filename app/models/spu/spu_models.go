//产品表
package spu

import (
	"mall/app/models"
	"mall/app/models/sku"
)

type SPU struct {
	models.BaseModel
	Title    string `json:"title" gorm:"not null;size:200;comment:标题"`
	SubTitle string `json:"subTitle" gorm:"size:200;comment:副标题"`
	// CategoryId uint   `json:"categoryId" gorm:"not null;index;comment:分类ID"`
	BrandId     uint       `json:"brandId" gorm:"index;comment:品牌ID"`
	SpecGroupId uint       `json:"spgId" gorm:"not null;index;comment:品类ID"`
	Saleable    bool       `json:"saleable" gorm:"not null;index;comment:是否上架"`
	Valid       bool       `json:"valid" gorm:"not null;index;comment:是否有效"`
	Skus        []*sku.SKU `json:"skus"`
	models.CommonTimestampsField
}
