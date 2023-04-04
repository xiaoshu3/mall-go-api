package sku

import (
	"mall/app/models"

	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type SKU struct {
	models.BaseModel
	SPUId uint   `json:"spu_id" gorm:"column:spu_id;not null;index;comment:产品ID"`
	Title string `json:"title" gorm:"not null;size:200;"`
	// Images   datatypes.JSON  `json:"images" gorm:"comment:商品图片"`
	Images   string          `json:"images" gorm:"comment:商品图片;size:200"`
	Price    decimal.Decimal `json:"price" gorm:"type:decimal(9,2);comment:商品价格"`
	Param    datatypes.JSON  `json:"param" gorm:"not null;comment:商品参数"`
	Saleable bool            `json:"saleable" gorm:"index;not null;"`
	Valid    bool            `json:"valid" gorm:"index;not null;"`

	models.CommonTimestampsField
}
