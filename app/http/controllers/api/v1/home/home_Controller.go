package home

import (
	v1 "mall/app/http/controllers/api/v1"
	"mall/app/models/home"
	"mall/pkg/response"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
	v1.BaseAPIController
}

func (hc *HomeController) GetCarousels(c *gin.Context) {
	data := home.AllCarousel()
	response.Data(c, data)
	// response.JSON(c, gin.H{
	// 	"data": data,
	// })
}

func (hc *HomeController) GetGrids(c *gin.Context) {
	data := home.AllGrid()
	response.Data(c, data)
}
