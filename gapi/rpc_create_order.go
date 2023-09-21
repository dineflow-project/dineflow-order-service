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

	order := &models.CreateOrderRequest{
		MenuId:   req.GetMenuId(),
		Status:   req.GetStatus(),
		VendorId: req.GetVendorId(),
		Price:    req.GetPrice(),
		Request:  req.GetRequest(),
		UserId:   req.GetUserId(),
	}

	newOrder, err := orderServer.orderService.CreateOrder(order)

	if err != nil {
		// if strings.Contains(err.Error(), "title already exists") {
		// 	return nil, status.Errorf(codes.AlreadyExists, err.Error())
		// }

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.OrderResponse{
		Order: &pb.Order{
			Id:        newOrder.Id.Hex(),
			Status:    order.Status,
			MenuId:    order.MenuId,
			VendorId:  order.VendorId,
			Price:     order.Price,
			Request:   order.Request,
			UserId:    order.UserId,
			CreatedAt: timestamppb.New(newOrder.CreateAt),
			UpdatedAt: timestamppb.New(newOrder.UpdatedAt),
		},
	}
	return res, nil
}
