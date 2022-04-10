package cart

import (
	"sepatuku-project/entity/product"

	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	UserId     uint            `json:"user_id" form:"user_id"`
	ProductId  uint            `json:"product_id" form:"product_id"`
	Quantity   uint            `json:"quantity" form:"quantity"`
	TotalPrice uint            `json:"total_price" form:"total_price"`
	Status     string          `json:"status" form:"status"`
	Product    product.Product `json:"product" form:"product" gorm:"foreignKey:ProductId;references:ID"`
}
