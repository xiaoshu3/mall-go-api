package specgroup

import (
	"mall/app/models"
	"mall/app/models/spu"
	"mall/pkg/database"
)

// 品类表
type SpecGroup struct {
	models.BaseModel

	SpgID     uint64          `json:"spgId,omitempty" gorm:"uniqueIndex"`
	Name      string          `json:"name,omitempty" gorm:"uniqueIndex;size:256"`
	Sort      uint            `json:"sort" gorm:"not null;comment:排名指数"`
	Brands    []*Brand        `json:"brands" gorm:"many2many:spec_brand_relation"`
	Categorys []*SpecCategory `json:"categorys"`
}

// 品牌表
type Brand struct {
	models.BaseModel
	Name        string `json:"name,omitempty" gorm:"size:200;not null;"`
	Image       string `json:"image,omitempty" gorm:"size:500;comment:图片地址"`
	RedirectUrl string `json:"redirectUrl" gorm:"size:200"`
	// Letter byte   `json:"letter,omitempty" gorm:"not null;comment:品牌首字母"`
	Spus []*spu.SPU `json:"spus"`
}

// 品类下细分
type SpecCategory struct {
	models.BaseModel
	Name        string     `json:"name,omitempty" gorm:"size:200;not null;"`
	Image       string     `json:"image,omitempty" gorm:"size:500;comment:图片地址"`
	Sort        uint       `json:"sort" gorm:"not null;default:100;comment:排名指数"`
	RedirectUrl string     `json:"redirectUrl" gorm:"size:200"`
	SpecGroupId uint       `json:"spg_id"`
	Spus        []*spu.SPU `json:"spus"`
}

func All() (categorys []SpecGroup) {
	database.DB.Model(&SpecGroup{}).Preload("Categorys").Preload("Brands").Order("sort").Find(&categorys)
	return
}
