package cart

import "mall/app/models"

type Cart struct {
	models.BaseModel
	UserID     uint `json:"user_id" gorm:"index;not null;"`
	SkuId      uint `json:"sku_id" gorm:"index;not null"`
	GoodsCount uint `json:"goods_count"`

	models.CommonTimestampsField
}
