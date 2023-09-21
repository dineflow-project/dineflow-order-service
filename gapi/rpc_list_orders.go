package gapi

import (
	"dineflow-order-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (orderServer *OrderServer) GetOrders(req *pb.GetOrdersRequest, stream pb.OrderService_GetOrdersServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	orders, err := orderServer.orderService.FindOrders(int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, order := range orders {
		stream.Send(&pb.Order{
			Id:        order.Id.Hex(),
			Status:    order.Status,
			MenuId:    order.MenuId,
			VenderId:  order.VenderId,
			Price:     order.Price,
			Request:   order.Request,
			UserId:    order.UserId,
			CreatedAt: timestamppb.New(order.CreateAt),
			UpdatedAt: timestamppb.New(order.UpdatedAt),
		})
	}

	return nil
}