package gapi

import (
	"context"
	"strings"

	"dineflow-order-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (orderServer *OrderServer) DeleteOrder(ctx context.Context, req *pb.OrderRequest) (*pb.DeleteOrderResponse, error) {
	orderId := req.GetId()

	if err := orderServer.orderService.DeleteOrder(orderId); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteOrderResponse{
		Success: true,
	}

	return res, nil
}
