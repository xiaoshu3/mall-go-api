package spu

import (
	// "encoding/json"

	// "log"
	v1 "mall/app/http/controllers/api/v1"
	"mall/pkg/response"

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
