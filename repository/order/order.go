package order

import (
	"gorm.io/gorm"
	_order "sepatuku-project/entity/order"
)

type RepositoryOrder struct {
	database *gorm.DB
}

func NewRepositoryOrder(database *gorm.DB) *RepositoryOrder {
	return &RepositoryOrder{
		database: database,
	}
}

func (ro *RepositoryOrder) GetOrderHistory() ([]_order.HistoryOrder, error) {
	var historyOrder []_order.HistoryOrder
	tx := ro.database.Find(&historyOrder)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return historyOrder, nil
}

func (ro *RepositoryOrder) CreateOrder(id int, order _order.Order) (_order.Order, error) {
	order.UserId = uint(id)
	tx := ro.database.Save(&order)

	if tx.Error != nil {
		return order, tx.Error
	}
	if tx.RowsAffected == 0 {
		return order, tx.Error
	}

	return order, nil
}
