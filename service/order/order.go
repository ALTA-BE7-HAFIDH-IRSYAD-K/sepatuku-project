package order

import (
	"sepatuku-project/entity/order"
	_orderService "sepatuku-project/repository/order"
)

type ServiceOrder struct {
	orderService _orderService.RepositoryOrderInterface
}

func (o *ServiceOrder) GetOrderHistoryById(id int) (order.Order, int, error) {
	//TODO implement me
	orderId, rows, err := o.orderService.GetOrderById(id)

	return orderId, rows, err
}

func NewOrderService(orderRepo _orderService.RepositoryOrderInterface) ServiceOrderInterface {
	return &ServiceOrder{
		orderService: orderRepo,
	}
}

func (o *ServiceOrder) UpdatedHistoryOrder(order order.Order, id int) (order.Order, error) {
	//TODO implement me
	orderId, _, err := o.orderService.GetOrderById(id)

	orderId.Status = order.Status

	if err != nil {
		return orderId, err
	}

	orderHistory, err := o.orderService.UpdatedHistoryOrder(orderId)
	return orderHistory, err
}

func (o *ServiceOrder) CreateOrder(order order.Order) (order.Order, error) {
	//TODO implement me
	newOrder, err := o.orderService.CreateOrder(order)
	return newOrder, err
}

func (o *ServiceOrder) GetOrderHistory(id int) ([]order.Order, error) {
	//TODO implement me
	getOrder, err := o.orderService.GetOrderHistory(id)
	return getOrder, err
}
