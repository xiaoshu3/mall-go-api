package spu

import (
	// "encoding/json"

	// "log"
	v1 "mall/app/http/controllers/api/v1"
	"mall/app/requests"
	"mall/pkg/response"

	"mall/app/models/cycle"
	"mall/app/models/spu"

	"github.com/gin-gonic/gin"
)

type SpuController struct {
	v1.BaseAPIController
}

func (sc *SpuController) GetSpuById(c *gin.Context) {
	id := c.Param("id")
	data := spu.Get(id)
	// logger.Dump

	// str := data.GoodsDetailContent
	// strList := strings.Split(str, " ")
	// fmt.Printf("len = %d\n",len(strList))
	if data.ID > 0 {
		response.Data(c, data)
	} else {
		response.Abort404(c)
		// response.Data(c, data.ID)
	}
}

func (sc *SpuController) GetSpus(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := cycle.Paginate(c, 8)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
