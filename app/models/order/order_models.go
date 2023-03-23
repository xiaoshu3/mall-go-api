package order

import (
	"mall/app/models"
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	models.BaseModel

	OrderNO    string          `json:"order_no" gorm:"not null;size:20;index"`
	UserID     uint            `json:"user_id" gorm:"index"`
	TotalPrice decimal.Decimal `json:"total_price" gorm:"not null;type:decimal(9,2)"`
	PayStatus  int             `json:"pay_status" gorm:"not null;comment:0.未支付 1.支付成功 2.支付失败"`
	PayType    int             `json:"pay_type" gorm:"not null;comment:0.无 1微信支付 2.支付宝支付"`
	PayTime    time.Time       `json:"pay_time" gorm:"not null"`
	Status     int             `json:"status" gorm:"index;not null;comment: 0.待支付 1.已支付 2.配货完成 3.出库成功4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭"`
	Info       string          `json:"info" gorm:"not null"`

	OrderDetails []OrderDetail `json:"order_detail,omitempty"`
	OrderAddress OrderAddress  `json:"order_address,omitempty"`
	models.CommonTimestampsField
}

type OrderDetail struct {
	models.BaseModel
	OrderID      uint64          `json:"order_id"`
	SkuID        uint            `json:"sku_id" gorm:""`
	SkuName      string          `json:"sku_name" gorm:"size:200"`
	SkuCount     uint            `json:"sku_count" gorm:"size:20"`
	SellingPrice decimal.Decimal `json:"selling_price" gorm:"type:decimal(9,2)"`
	SkuCoverImg  string          `json:"sku_cover_img" gorm:"size:200"`
	models.CommonTimestampsField
}

type OrderAddress struct {
	OrderID       uint64 `json:"order_id,omitempty" gorm:"primaryKey;autoIncrement"`
	UserName      string `json:"user_name,omitempty" gorm:"size:200;"`
	UserPhone     string `json:"user_phone,omitempty" gorm:"size:20;"`
	ProvinceName  string `json:"province_name,omitempty" gorm:"size:32"`
	CityName      string `json:"city_name,omitempty" gorm:"size:30"`
	RegionName    string `json:"region_name,omitempty" gorm:"size:32"`
	DetailAddress string `json:"detail_address,omitempty" gorm:"size:64"`
}
