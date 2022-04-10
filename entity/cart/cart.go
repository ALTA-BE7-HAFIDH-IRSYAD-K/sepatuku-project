package cart

import (
	"sepatuku-project/entity/product"

	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	UserId       uint            `json:"user_id" form:"user_id"`
	ProductId    uint            `json:"product_id" form:"product_id"`
	Quantity     uint            `json:"quantity" form:"quantity"`
	TotalPrice   uint            `json:"total_price" form:"total_price"`
	Status       string          `json:"status" form:"status"`
	Name_product string          `json:"name_product" form:"name_product"`
	Description  string          `json:"description" form:"description"`
	Price        uint            `json:"price" form:"price"`
	Image        string          `json:"image" form:"image"`
	Stock        uint            `json:"stock" form:"stock"`
	Product      product.Product `gorm:"foreignKey:ProductId;references:ID"`
}

type CartResponse struct {
	UserId       uint   `json:"user_id" form:"user_id"`
	ProductId    uint   `json:"product_id" form:"product_id"`
	Quantity     uint   `json:"quantity" form:"quantity"`
	TotalPrice   uint   `json:"total_price" form:"total_price"`
	Status       string `json:"status" form:"status"`
	Name_product string `json:"name_product" form:"name_product"`
	Description  string `json:"description" form:"description"`
	Price        uint   `json:"price" form:"price"`
	Image        string `json:"image" form:"image"`
	Stock        uint   `json:"stock" form:"stock"`
}
type CartResponseCreate struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	ProductId uint   `json:"product_id" form:"product_id"`
	Status    string `json:"status" form:"status"`
}
