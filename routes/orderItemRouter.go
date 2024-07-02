package routes

import (
	controller "github.com/PoudelAmrit123/go-rsm/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controller.GetOrderItems())

	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())

	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())

	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

	incomingRoutes.GET("/orderItem-order/:order_id", controller.GetOrderItemsByOrder())

}
