package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github/LordRadamanthys/orders_with_kafka/application/domain"
	"github/LordRadamanthys/orders_with_kafka/application/model"
	"github/LordRadamanthys/orders_with_kafka/application/ports/output"
	kafkaconfig "github/LordRadamanthys/orders_with_kafka/configuration/kafka_config"
	"log"

	"github.com/jinzhu/copier"
)

type ConsumeEventStruct struct {
	orderPort output.OrderPort
}

func NewConsumeEventStruct(orderPort output.OrderPort) *ConsumeEventStruct {
	return &ConsumeEventStruct{
		orderPort: orderPort,
	}
}

func (c *ConsumeEventStruct) ConsumeEvent() {
	for {
		m, err := kafkaconfig.KafkaClient.ReadMessage(context.Background())

		if err != nil {
			log.Fatalf("failed to read message: %v", err)
		}
		obj := model.OrderEvent{}
		err = json.Unmarshal(m.Value, &obj)
		if err != nil {
			log.Fatalf("failed to unmarshal message: %v", err)
		}
		// Converta o array de bytes para a struct Pedido
		fmt.Printf("Pedido recebido: %v\n", obj)

		var orderDomain domain.OrderEvent

		if err := copier.Copy(&orderDomain, &obj); err != nil {
			panic(err)
		}

		orderDomain.ClientId = obj.CodeClient
		orderDomain.OrderId = obj.CodeOrder

		c.orderPort.SaveOrder(orderDomain)

	}
}
