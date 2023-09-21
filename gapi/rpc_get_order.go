package gapi

import (
	"context"
	"strings"

	"dineflow-order-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (orderServer *OrderServer) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	orderId := req.GetId()

	order, err := orderServer.orderService.FindOrderById(orderId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	orderMenus := ModelToProtoOrderMenus(order.OrderMenus)

	res := &pb.OrderResponse{
		Order: &pb.Order{
			Id:         order.Id.Hex(),
			Status:     order.Status,
			OrderMenus: orderMenus,
			VendorId:   order.VendorId,
			Price:      order.Price,
			UserId:     order.UserId,
			CreatedAt:  timestamppb.New(order.CreateAt),
			UpdatedAt:  timestamppb.New(order.UpdatedAt),
		},
	}
	return res, nil
}
