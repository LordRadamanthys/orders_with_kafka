package controllers

import (
	"github/LordRadamanthys/orders_with_kafka/application/ports/input"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrdersController struct {
	serviceOrder input.FindOrdersUseCase
}

func NewOrdersController(serviceOrder input.FindOrdersUseCase) *OrdersController {
	return &OrdersController{
		serviceOrder: serviceOrder,
	}
}

func (c *OrdersController) FindOrdersByClient(ctx *gin.Context) {

	idClient := ctx.Param("clientId")

	clientId, err := strconv.Atoi(idClient)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	order := c.serviceOrder.FindOrdersByClient(clientId)

	ctx.JSON(200, order)
}

func (c *OrdersController) FindLenOrdersByClient(ctx *gin.Context) {

	idClient := ctx.Param("clientId")

	clientId, err := strconv.Atoi(idClient)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := c.serviceOrder.GetTotalOrdersByClient(clientId)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, order)
}

func (c *OrdersController) Checkout(ctx *gin.Context) {
	idClient := ctx.Param("clientId")

	clientId, err := strconv.Atoi(idClient)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := c.serviceOrder.CheckoutOrders(clientId)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}
