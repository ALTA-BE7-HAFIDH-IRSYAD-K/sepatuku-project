package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	Name_product string `json:"name_product" json:"name_product"`
	Description  string `json:"description" form:"description"`
	Price        uint   `json:"price" form:"price"`
	Image        string `json:"image" form:"image"`
	Stock        uint   `json:"stock" form:"stock"`
}
