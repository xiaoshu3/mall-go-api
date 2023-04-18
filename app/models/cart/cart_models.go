package cart

import (
	"mall/app/models"
	"mall/app/models/spu"
	"mall/pkg/database"
)

type Cart struct {
	models.BaseModel
	UserID     uint64    `json:"user_id" gorm:"index:idx_user_spu;not null;"`
	SpuId      uint64    `json:"spu_id" gorm:"index:idx_user_spu;not null"`
	GoodsCount uint      `json:"goods_count"`
	Spus       []spu.SPU `gorm:"many2many:cart_spu;"`

	models.CommonTimestampsField
}

type CartList struct {
	models.BaseModel
	SpuID         uint64
	GoodsCoverImg string `json:"goods_cover_img,omitempty" gorm:"200;comment:商品封面图;"`
	Title         string `json:"title,,omitempty" gorm:"not null;size:200;comment:标题"`
	Price         uint   `json:"price,omitempty" gorm:"size:10;index;comment:商品价格"`
	GoodsCount    uint   `json:"goods_count" `
}

func (cartModel *Cart) Create() {
	database.DB.Create(&cartModel)
}

func (cartModel *Cart) Delete() error {
	return database.DB.Delete(&cartModel).Error
}

func (cartModel *Cart) Save() (rowsAffected int64) {
	result := database.DB.Save(&cartModel)
	return result.RowsAffected
}

func (cartModel *Cart) IsRecordExist() bool {
	database.DB.Model(&Cart{}).Where(&cartModel).Find(&cartModel)
	return cartModel.ID > 0
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(&Cart{}).Where(field+" = ?", value).Count(&count)
	// logger.Dump(count,"Count")
	return count > 0
}

func UpdateGoodsCount(id uint64, goodsCount uint) {
	database.DB.Model(&Cart{}).Where("id = ?", id).Update("goods_count", goodsCount)
}
