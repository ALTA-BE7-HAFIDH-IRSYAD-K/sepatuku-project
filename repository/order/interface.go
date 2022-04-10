package order

import (
	"sepatuku-project/entity/order"
)

type RepositoryOrderInterface interface {
	CreateOrder(order order.Order) (order.Order, error)
	GetOrderById(id int) (order.Order, int, error)
	GetOrderHistory() ([]order.Order, error)
	UpdatedHistoryOrder(order order.Order) (order.Order, error)
}
