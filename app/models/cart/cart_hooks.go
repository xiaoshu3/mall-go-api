package cart

import "gorm.io/gorm"

func (cartModel *Cart) AfterCreate(tx *gorm.DB) (err error) {
	tx.Table("cart_spu").Create(map[string]interface{}{
		"cart_id": cartModel.ID,
		"sp_uid":  cartModel.SpuId,
	})
	return
}

func (cartModel *Cart) BeforeDelete(tx *gorm.DB) (err error) {
	// tx.Table("cart_spu").Delete(map[string]interface{}{
	// 	"cart_id": cartModel.ID,
	// 	// "sp_uid":  cartModel.SpuId,
	// })
	tx.Exec("delete from cart_spu where cart_id = ?", cartModel.ID)
	return
}
