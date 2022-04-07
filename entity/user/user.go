package user

import (
	"sepatuku-project/entity/cart"
	_prodcut "sepatuku-project/entity/product"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string             `gorm:"unique" json:"username" form:"username"`
	Email    string             `gorm:"unique" json:"email" form:"email"`
	Password string             `json:"password" form:"password"`
	Address  string             `json:"address" form:"password"`
	Phone    string             `json:"phone" form:"phone"`
	Cart     []cart.Cart        `gorm:"foreignKey:UserId;references:ID" json:"cart" form:"cart"`
	Product  []_prodcut.Product `gorm:"foreignKey:UserID;references:ID"`
}
