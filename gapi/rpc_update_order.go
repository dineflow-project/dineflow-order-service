package gapi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"dineflow-order-service/models"
	"dineflow-order-service/pb"
	"dineflow-order-service/rabbitmq"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (orderServer *OrderServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.OrderResponse, error) {
	orderId := req.GetId()

	// Convert the protobuf request to your model
	order := &models.UpdateOrder{
		Status:     req.GetStatus(),
		VendorId:   req.GetVendorId(),
		UserId:     req.GetUserId(),
		OrderMenus: ProtoToModelUpdateOrderMenu(req.GetOrderMenus()), // Convert order menus
		UpdatedAt:  time.Now(),
	}

	// Calculate the total price by summing the prices of OrderMenu items
	order.Price = CalculateTotalPrice(order.OrderMenus)

	updatedOrder, err := orderServer.orderService.UpdateOrder(orderId, order)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Convert the updated order to protobuf response
	orderMenus := ModelToProtoOrderMenus(order.OrderMenus)
	res := &pb.OrderResponse{
		Order: &pb.Order{
			Id:         updatedOrder.Id.Hex(),
			Status:     updatedOrder.Status,
			VendorId:   updatedOrder.VendorId,
			OrderMenus: orderMenus,
			Price:      updatedOrder.Price, // Use the calculated total price
			UserId:     updatedOrder.UserId,
			CreatedAt:  timestamppb.New(updatedOrder.CreateAt),
			UpdatedAt:  timestamppb.New(updatedOrder.UpdatedAt),
		},
	}

	notiType := ""
	if updatedOrder.Status == "cooking" {
		notiType = "cooking order"
	} else if updatedOrder.Status == "finish" {
		notiType = "finish order"
	}
	err = rabbitmq.PushNotification(updatedOrder.VendorId, updatedOrder.Id.Hex(), notiType)
	if err != nil {
		fmt.Println("Error RabbitMQ: ", err)
	}

	return res, err
}

func ProtoToModelUpdateOrderMenu(protoOrderMenu []*pb.UpdateOrderRequest_OrderMenu) []*models.OrderMenu {
	var modelOrderMenus []*models.OrderMenu

	for _, protoMenu := range protoOrderMenu {
		modelMenu := &models.OrderMenu{
			MenuId:  protoMenu.MenuId,
			Amount:  int(protoMenu.Amount),
			Price:   protoMenu.Price,
			Request: protoMenu.Request,
		}
		modelOrderMenus = append(modelOrderMenus, modelMenu)
	}

	return modelOrderMenus
}
