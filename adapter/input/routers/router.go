package routers

import (
	"github/LordRadamanthys/orders_with_kafka/adapter/input/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func LoadRouter(controllers controllers.OrdersController) {

	router.GET("/orders/client/:clientId", controllers.FindOrdersByClient)
	router.GET("/numOrders/client/:clientId", controllers.FindLenOrdersByClient)
	router.GET("/orders/client/:clientId/checkout", controllers.Checkout)

	router.Run(":8080")
}
