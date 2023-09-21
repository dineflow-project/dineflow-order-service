package gapi

import (
	"context"

	"dineflow-order-service/models"
	"dineflow-order-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (orderServer *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {

	// Convert the protobuf request to your model
	order := &models.CreateOrderRequest{
		Status:     req.GetStatus(),
		VendorId:   req.GetVendorId(),
		UserId:     req.GetUserId(),
		OrderMenus: ProtoToModelOrderMenu(req.GetOrderMenus()), // Convert order menus
	}

	// Calculate the total price by summing the prices of OrderMenu items
	order.Price = CalculateTotalPrice(order.OrderMenus)

	newOrder, err := orderServer.orderService.CreateOrder(order)

	if err != nil {
		// Handle the error accordingly
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Convert the created order to protobuf response
	res := &pb.OrderResponse{
		Order: &pb.Order{
			Id:        newOrder.Id.Hex(),
			Status:    newOrder.Status,
			VendorId:  newOrder.VendorId,
			Price:     newOrder.Price, // Use the calculated total price
			UserId:    newOrder.UserId,
			CreatedAt: timestamppb.New(newOrder.CreateAt),
			UpdatedAt: timestamppb.New(newOrder.UpdatedAt),
		},
	}
	return res, nil
}

func ProtoToModelOrderMenu(protoOrderMenu []*pb.CreateOrderRequest_OrderMenu) []*models.OrderMenu {
	var modelOrderMenus []*models.OrderMenu

	for _, protoMenu := range protoOrderMenu {
		modelMenu := &models.OrderMenu{
			MenuId:  protoMenu.MenuId,
			Request: protoMenu.Request,
			Price:   protoMenu.Price, // Set the Price field from the protobuf message
		}
		modelOrderMenus = append(modelOrderMenus, modelMenu)
	}

	return modelOrderMenus
}

func CalculateTotalPrice(orderMenus []*models.OrderMenu) float32 {
	totalPrice := float32(0)

	for _, menu := range orderMenus {
		totalPrice += menu.Price
	}

	return totalPrice
}
