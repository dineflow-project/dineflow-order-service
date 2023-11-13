package services

import "dineflow-order-service/models"

type OrderService interface {
	CreateOrder(*models.CreateOrderRequest) (*models.DBOrder, error)
	UpdateOrder(string, *models.UpdateOrder) (*models.DBOrder, error)
	FindOrderById(string) (*models.DBOrder, error)
	FindOrders(page int, limit int) ([]*models.DBOrder, error)
	FindOrdersByUserId(UserId string, page int, limit int) ([]*models.DBOrder, error)
	FindOrdersByVendorId(VendorId string, page int, limit int) ([]*models.DBOrder, error)
	DeleteOrder(string) error
}
