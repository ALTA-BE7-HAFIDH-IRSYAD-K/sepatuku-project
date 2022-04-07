package order

import "sepatuku-project/entity/order"

type RepositoryOrderInterface interface {
	CreateOrder(id int, order order.Order) (order.Order, error)
	GetOrderHistory() ([]order.HistoryOrder, error)
}
