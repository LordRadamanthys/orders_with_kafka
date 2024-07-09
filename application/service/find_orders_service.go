package service

import (
	"fmt"
	"github/LordRadamanthys/orders_with_kafka/application/domain"
	"github/LordRadamanthys/orders_with_kafka/application/ports/output"
	"math"
)

type FindOrdersServiceStruct struct {
	orderPort output.OrderPort
}

func NewFindOrdersService(orderPort output.OrderPort) *FindOrdersServiceStruct {
	return &FindOrdersServiceStruct{
		orderPort: orderPort,
	}
}

func (fo *FindOrdersServiceStruct) FindOrdersByClient(clientId int) *[]domain.OrderEvent {

	result, err := fo.orderPort.FindByCodeClient(clientId)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	return result
}

func (fo *FindOrdersServiceStruct) GetTotalOrdersByClient(clientId int) (int, error) {

	result, err := fo.orderPort.FindByCodeClient(clientId)

	if err != nil {
		fmt.Println("Error: ", err)
		return 0, nil
	}

	v := len(*result)
	return v, nil
}

func (fo *FindOrdersServiceStruct) CheckoutOrders(clientId int) (float64, error) {
	result, err := fo.orderPort.FindByCodeClient(clientId)

	if err != nil {
		fmt.Println("Error: ", err)
		return 0, nil
	}

	sumOder := 0.0
	for _, v := range *result {

		sumOder += SumOrdersValues(&v.Items)

	}

	return sumOder, nil
}

func SumOrdersValues(orders *[]domain.Item) float64 {

	sum := 0.0
	for _, v := range *orders {
		sum += (v.Price * float64(v.Quanty))
	}
	return roundFloat(sum, 2)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
