package order

import (
	"fmt"
	_order "sepatuku-project/entity/order"

	"gorm.io/gorm"
)

type RepositoryOrder struct {
	database *gorm.DB
}

func NewRepositoryOrder(database *gorm.DB) *RepositoryOrder {
	return &RepositoryOrder{
		database: database,
	}
}

func (ro *RepositoryOrder) GetOrderById(id int) (_order.Order, int, error) {
	var order _order.Order
	tx := ro.database.Find(&order, id)

	if tx.Error != nil {
		return order, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return order, 0, tx.Error
	}
	return order, int(tx.RowsAffected), nil
}

func (ro *RepositoryOrder) GetOrderHistory(id int) ([]_order.Order, error) {
	var historyOrder []_order.Order

	tx := ro.database.Preload("Cart").Preload("Cart.Product").Where("user_id", id).Find(&historyOrder)

	fmt.Println("historyOrder", historyOrder)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return historyOrder, nil
}

func (ro *RepositoryOrder) CreateOrder(order _order.Order) (_order.Order, error) {
	fmt.Println("order", order)
	tx := ro.database.Save(&order)

	if tx.Error != nil {
		return order, tx.Error
	}
	if tx.RowsAffected == 0 {
		return order, tx.Error
	}

	return order, nil
}

func (ro *RepositoryOrder) UpdatedHistoryOrder(order _order.Order) (_order.Order, error) {
	tx := ro.database.Save(&order)
	if tx.Error != nil {
		return order, tx.Error
	}
	if tx.RowsAffected == 0 {
		return order, tx.Error
	}
	return order, nil
}
