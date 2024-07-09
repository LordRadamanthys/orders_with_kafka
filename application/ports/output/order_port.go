package output

import (
	"github/LordRadamanthys/orders_with_kafka/application/domain"
)

type OrderPort interface {
	SaveOrder(order domain.OrderEvent) error
	FindByCodeClient(clientId int) (*[]domain.OrderEvent, error)
}
