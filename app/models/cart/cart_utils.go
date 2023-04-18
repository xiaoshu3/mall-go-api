package cart

import "mall/pkg/database"

// func IsRecordExist(userId, spuId uint64) (bool, uint) {
// 	var cartModel Cart
// 	var count int64
// 	database.DB.Model(&Cart{}).Where("user_id = ? and spu_id = ?", userId, spuId).Find(&cartModel).Count(&count)
// 	return count > 0, cartModel.GoodsCount
// }

func GetCartList(userId string, cartList interface{}) {
	database.DB.
		Raw("SELECT a.id,a.goods_count,a.spu_id, b.title,b.price,b.goods_cover_img  from carts a left join spus b on a.spu_id = b.id  where a.user_id = ? order by a.created_at desc", userId).
		Scan(cartList)
}

func GetCartListByPreload(userId string, cartList interface{}) {
	database.DB.Model(&Cart{}).Preload("Spus", "saleable = ? and valid = ?", 1, 1).Where("user_id = ? ", userId).
		Order("created_at desc").Find(cartList)
}
