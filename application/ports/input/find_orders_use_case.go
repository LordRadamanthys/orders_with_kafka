package input

import "github/LordRadamanthys/orders_with_kafka/application/domain"

type FindOrdersUseCase interface {
	FindOrdersByClient(clientId int) *[]domain.OrderEvent
	GetTotalOrdersByClient(clientId int) (int, error)
	CheckoutOrders(clientId int) (float64, error)
}
