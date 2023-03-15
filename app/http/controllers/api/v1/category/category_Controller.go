package category

import (
	v1 "mall/app/http/controllers/api/v1"
	"mall/app/models/specgroup"
	"mall/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	v1.BaseAPIController
}

func (cc *CategoryController) GetALlCategorys(c *gin.Context) {
	data := specgroup.All()
	response.Data(c, data)
}
