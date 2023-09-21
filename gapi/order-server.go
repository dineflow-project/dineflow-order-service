package gapi

import (
	"dineflow-order-service/pb"
	"dineflow-order-service/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	orderCollection *mongo.Collection
	orderService    services.OrderService
}

func NewGrpcOrderServer(orderCollection *mongo.Collection, orderService services.OrderService) (*OrderServer, error) {
	orderServer := &OrderServer{
		orderCollection: orderCollection,
		orderService:    orderService,
	}

	return orderServer, nil
}
