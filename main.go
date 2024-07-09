package main

import (
	"fmt"
	"github/LordRadamanthys/orders_with_kafka/adapter/input/controllers"
	"github/LordRadamanthys/orders_with_kafka/adapter/input/routers"
	"github/LordRadamanthys/orders_with_kafka/adapter/output/repository"
	"github/LordRadamanthys/orders_with_kafka/application/service"
	kafkaconfig "github/LordRadamanthys/orders_with_kafka/configuration/kafka_config"
	mongodbconfig "github/LordRadamanthys/orders_with_kafka/configuration/mongodb_config"
)

func main() {

	kafkaconfig.InitKafkaConfig()
	fmt.Println("Hello, World!")
	mongodbconfig.InitMongoDBConfig()

	repository := repository.NewOrderRepository()
	consumeService := service.NewConsumeEventStruct(repository)

	orderService := service.NewFindOrdersService(repository)
	controllerOder := controllers.NewOrdersController(orderService)

	go consumeService.ConsumeEvent()

	routers.LoadRouter(*controllerOder)

}
