package order

import "sepatuku-project/entity/order"

type ServiceOrderInterface interface {
	CreateOrder(id int, order order.Order) (order.Order, error)
	GetOrderHistory() ([]order.HistoryOrder, error)
}
