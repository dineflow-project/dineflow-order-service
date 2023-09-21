package routes

import (
	"dineflow-order-service/controllers"

	"github.com/gin-gonic/gin"
)

type OrderRouteController struct {
	orderController controllers.OrderController
}

func NewOrderControllerRoute(orderController controllers.OrderController) OrderRouteController {
	return OrderRouteController{orderController}
}

func (r *OrderRouteController) OrderRoute(rg *gin.RouterGroup) {
	router := rg.Group("/orders")

	router.GET("/", r.orderController.FindOrders)
	router.GET("/:orderId", r.orderController.FindOrderById)
	router.POST("/", r.orderController.CreateOrder)
	router.PATCH("/:orderId", r.orderController.UpdateOrder)
	router.DELETE("/:orderId", r.orderController.DeleteOrder)
}
