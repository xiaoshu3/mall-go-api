package cycle

import (
	"mall/app/models/specgroup"
	"mall/app/models/spu"
	"mall/pkg/app"
	"mall/pkg/database"
	"mall/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Paginate(c *gin.Context, perPage int) (spus []spu.SPU, paging paginator.Paging) {

	categoryId := c.Query("categoryId")
	brandId := c.Query("brandId")
	specGroupId := c.Query("specGroupId")
	keyword := c.Query("keyword")

	// if categoryId == "" && brandId == "" && keyword == "" {
	// 	return
	// }

	db := database.DB.Model(spu.SPU{}).Select("id", "goods_cover_img", "title", "sub_title", "price").
		Where("saleable = ? ", "1").Where("valid = ? ", "1")

	if categoryId != "" {
		db = db.Where("spec_category_id = ?", categoryId)
	}

	if keyword != "" {
		db = db.Where("title LIKE ?", "%"+keyword+"%")
	}

	if brandId != "" {
		specCategoryIds := specgroup.GetIdsBySpecGroupId(specGroupId)
		db = db.Where("brand_id = ? AND spec_category_id in ?", brandId, specCategoryIds)
	}

	paging = paginator.Paginate(
		c,
		db,
		&spus,
		app.V1URL(database.TableName(&spu.SPU{})),
		perPage,
	)
	return
}
