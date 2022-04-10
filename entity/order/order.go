package order

import (
	"sepatuku-project/entity/cart"

	"gorm.io/gorm"
)

type Order struct {
	*gorm.Model
	UserId       uint      `json:"user_id" form:"user_id"`
	CartId       uint      `json:"cart_id" form:"cart_id	"`
	Street       string    `json:"street" form:"street"`
	City         string    `json:"city" form:"city"`
	Province     string    `json:"province" form:"province"`
	ZipCode      uint      `json:"zip_code" form:"zip_code"`
	CardType     string    `json:"card_type" form:"card_type"`
	CardName     string    `json:"card_name" form:"card_name"`
	CardNumber   string    `json:"card_number" form:"card_number"`
	Cvv          uint8     `json:"cvv" form:"cvv"`
	ExpiredMonth string    `json:"expired_month" form:"expired_month"`
	ExpiredYear  uint8     `json:"expired_year" form:"expired_year"`
	TotalQty     uint      `json:"total_qty" form:"total_qty"`
	TotalPrice   uint      `json:"total_price" form:"total_price"`
	Status       string    `json:"status" form:"status"`
	Cart         cart.Cart `json:"cart" form:"cart" gorm:"foreignKey:CartId;references:ID"`
}

type Address struct {
	Street   string `json:"street" form:"street"`
	City     string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	ZipCode  uint   `json:"zip_code" form:"zip_code"`
}

type CreditCard struct {
	CardType     string `json:"card_type" form:"card_type"`
	CardName     string `json:"card_name" form:"card_name"`
	CardNumber   uint   `json:"card_number" form:"card_number"`
	Cvv          uint8  `json:"cvv" form:"cvv"`
	ExpiredMonth string `json:"expired_month" form:"expired_month"`
	ExpiredYear  uint8  `json:"expired_year" form:"expired_year"`
}

type FormatOrderRequest struct {
	CartId     []uint     `json:"cart_id" form:"cart_id"`
	Address    Address    `json:"address" form:"address"`
	CreditCard CreditCard `json:"credit_card" form:"credit_card"`
}

type HistoryOrder struct {
	*gorm.Model
	UserId  uint      `json:"user_id" form:"user_id"`
	CartId  uint      `json:"cart_id" form:"cart_id"`
	OrderId uint      `json:"order_id" form:"order_id"`
	Order   Order     `json:"order" form:"order" gorm:"foreignKey:OrderId;references:ID"`
	Cart    cart.Cart `json:"cart" form:"cart" gorm:"foreignKey:CartId;references:ID"`
	Status  string    `json:"status" form:"status"`
}

type Status struct {
	Status string `json:"status" form:"status"`
}
