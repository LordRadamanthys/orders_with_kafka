package output

import "github/LordRadamanthys/orders_with_kafka/application/model"

type ConsumerPort interface {
	ConsumeEvent(orderEvent model.OrderEvent) error
}
