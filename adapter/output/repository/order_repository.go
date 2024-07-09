package repository

import (
	"context"
	"github/LordRadamanthys/orders_with_kafka/application/domain"
	mongodbconfig "github/LordRadamanthys/orders_with_kafka/configuration/mongodb_config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type OrderRepositoryStruct struct{}

func NewOrderRepository() *OrderRepositoryStruct {
	return &OrderRepositoryStruct{}
}

func (or *OrderRepositoryStruct) SaveOrder(order domain.OrderEvent) error {

	_, err := mongodbconfig.MongoClient.InsertOne(context.Background(), order)

	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	return nil
}

func (or *OrderRepositoryStruct) FindByCodeClient(clientId int) (*[]domain.OrderEvent, error) {

	var orders []domain.OrderEvent
	cursor, err := mongodbconfig.MongoClient.Find(context.Background(), bson.M{"clientId": clientId})
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var order domain.OrderEvent
		err := cursor.Decode(&order)
		if err != nil {
			log.Println("Error: ", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return &orders, nil
}
