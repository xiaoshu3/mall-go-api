package home

import (
	"mall/app/models"
	"mall/pkg/database"
)

type HomeCarousel struct {
	models.BaseModel
	CarouselUrl string `json:"carouselUrl" gorm:"size:200"`
	RedirectUrl string `json:"redirectUrl" gorm:"size:200"`
	Name        string `json:"name" gorm:"size:20"`
	Sort        uint   `json:"sort"`
}

func AllCarousel() (carousels []HomeCarousel) {
	// database.DB.Where("name = ?", nil).Find(&carousels)
	database.DB.Where(map[string]interface{}{"name": nil}).Find(&carousels).Order("sort")
	return
}

func AllGrid() (grids []HomeCarousel) {
	database.DB.Where("name <> ?", "").Find(&grids).Order("sort")
	return
}
