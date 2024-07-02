package routes

import (
	controller "github.com/PoudelAmrit123/go-rsm/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/foods", controller.GetFoods())

	incomingRoutes.GET("/food/:food_id", controller.GetFood())

	incomingRoutes.POST("/foods", controller.CreateFood())

	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())

}
