package cart

import (
	"gorm.io/gorm"
	"sepatuku-project/entity/product"
)

type Cart struct {
	*gorm.Model
	UserId     uint              `json:"user_id" form:"user_id"`
	ProductId  uint              `json:"product_id" form:"product_id"`
	Quantity   uint              `json:"quantity" form:"quantity"`
	TotalPrice uint              `json:"total_price" form:"total_price"`
	Product    []product.Product `json:"product" form:"product" gorm:"foreignKey:ID;references:ProductId"`
}
