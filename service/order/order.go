package order

import (
	"sepatuku-project/entity/order"
	_orderService "sepatuku-project/repository/order"
)

type ServiceOrder struct {
	orderService _orderService.RepositoryOrderInterface
}

func NewOrderService(orderRepo _orderService.RepositoryOrderInterface) ServiceOrderInterface {
	return &ServiceOrder{
		orderService: orderRepo,
	}
}

func (o *ServiceOrder) CreateOrder(id int, order order.Order) (order.Order, error) {
	//TODO implement me
	newOrder, err := o.orderService.CreateOrder(id, order)
	return newOrder, err
}

func (o *ServiceOrder) GetOrderHistory() ([]order.HistoryOrder, error) {
	//TODO implement me
	getOrder, err := o.orderService.GetOrderHistory()
	return getOrder, err
}
