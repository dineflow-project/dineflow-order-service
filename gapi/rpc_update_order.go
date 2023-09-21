package gapi

import (
	"context"
	"strings"
	"time"

	"dineflow-order-service/models"
	"dineflow-order-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (orderServer *OrderServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.OrderResponse, error) {
	orderId := req.GetId()

	order := &models.UpdateOrder{
		MenuId:    req.GetMenuId(),
		Status:    req.GetStatus(),
		VendorId:  req.GetVendorId(),
		Price:     req.GetPrice(),
		Request:   req.GetRequest(),
		UserId:    req.GetUserId(),
		UpdatedAt: time.Now(),
	}

	updatedOrder, err := orderServer.orderService.UpdateOrder(orderId, order)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.OrderResponse{
		Order: &pb.Order{
			Id:        updatedOrder.Id.Hex(),
			Status:    order.Status,
			MenuId:    order.MenuId,
			VendorId:  order.VendorId,
			Price:     order.Price,
			Request:   order.Request,
			UserId:    order.UserId,
			CreatedAt: timestamppb.New(updatedOrder.CreateAt),
			UpdatedAt: timestamppb.New(updatedOrder.UpdatedAt),
		},
	}
	return res, nil
}
