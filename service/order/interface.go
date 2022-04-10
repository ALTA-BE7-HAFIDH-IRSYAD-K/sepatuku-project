package order

import (
	"sepatuku-project/entity/order"
)

type ServiceOrderInterface interface {
	CreateOrder(order order.Order) (order.Order, error)
	GetOrderHistory(id int) ([]order.Order, error)
	GetOrderHistoryById(id int) (order.Order, int, error)
	UpdatedHistoryOrder(order order.Order, id int) (order.Order, error)
}
