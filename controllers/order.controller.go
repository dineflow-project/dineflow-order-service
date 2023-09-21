package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"dineflow-order-service/models"
	"dineflow-order-service/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return OrderController{orderService}
}

func (pc *OrderController) CreateOrder(ctx *gin.Context) {
	var order *models.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newOrder, err := pc.orderService.CreateOrder(order)

	if err != nil {
		// if strings.Contains(err.Error(), "title already exists") {
		// 	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
		// 	return
		// }

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newOrder})
}

func (pc *OrderController) UpdateOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	var order *models.UpdateOrder
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedOrder, err := pc.orderService.UpdateOrder(orderId, order)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedOrder})
}

func (pc *OrderController) FindOrderById(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	order, err := pc.orderService.FindOrderById(orderId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": order})
}

func (pc *OrderController) FindOrders(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	orders, err := pc.orderService.FindOrders(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(orders), "data": orders})
}

func (pc *OrderController) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	err := pc.orderService.DeleteOrder(orderId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
