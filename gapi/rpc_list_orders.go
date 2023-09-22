package gapi

import (
	"dineflow-order-service/models"
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
		// Convert OrderMenus to protobuf format
		orderMenus := ModelToProtoOrderMenus(order.OrderMenus)

		stream.Send(&pb.Order{
			Id:         order.Id.Hex(),
			Status:     order.Status,
			OrderMenus: orderMenus,
			VendorId:   order.VendorId,
			Price:      order.Price,
			UserId:     order.UserId,
			CreatedAt:  timestamppb.New(order.CreateAt),
			UpdatedAt:  timestamppb.New(order.UpdatedAt),
		})
	}

	return nil
}

func (orderServer *OrderServer) GetOrdersByUserId(req *pb.GetOrdersByUserIdRequest, stream pb.OrderService_GetOrdersByUserIdServer) error {
	userId := req.GetUserId()
	page := req.GetPage()
	limit := req.GetLimit()

	orders, err := orderServer.orderService.FindOrdersByUserId(userId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, order := range orders {
		// Convert OrderMenus to protobuf format
		orderMenus := ModelToProtoOrderMenus(order.OrderMenus)

		stream.Send(&pb.Order{
			Id:         order.Id.Hex(),
			Status:     order.Status,
			OrderMenus: orderMenus,
			VendorId:   order.VendorId,
			Price:      order.Price,
			UserId:     order.UserId,
			CreatedAt:  timestamppb.New(order.CreateAt),
			UpdatedAt:  timestamppb.New(order.UpdatedAt),
		})
	}

	return nil
}

func (orderServer *OrderServer) GetOrdersByVendorId(req *pb.GetOrdersByVendorIdRequest, stream pb.OrderService_GetOrdersByVendorIdServer) error {
	vendorId := req.GetVendorId()
	page := req.GetPage()
	limit := req.GetLimit()

	orders, err := orderServer.orderService.FindOrdersByVendorId(vendorId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, order := range orders {
		// Convert OrderMenus to protobuf format
		orderMenus := ModelToProtoOrderMenus(order.OrderMenus)

		stream.Send(&pb.Order{
			Id:         order.Id.Hex(),
			Status:     order.Status,
			OrderMenus: orderMenus,
			VendorId:   order.VendorId,
			Price:      order.Price,
			UserId:     order.UserId,
			CreatedAt:  timestamppb.New(order.CreateAt),
			UpdatedAt:  timestamppb.New(order.UpdatedAt),
		})
	}

	return nil
}

func ModelToProtoOrderMenus(modelOrderMenus []*models.OrderMenu) []*pb.Order_OrderMenu {
	var protoOrderMenus []*pb.Order_OrderMenu

	for _, modelMenu := range modelOrderMenus {
		protoMenu := &pb.Order_OrderMenu{
			MenuId:  modelMenu.MenuId,
			Price:   modelMenu.Price,
			Amount:  int32(modelMenu.Amount),
			Request: modelMenu.Request,
		}
		protoOrderMenus = append(protoOrderMenus, protoMenu)
	}

	return protoOrderMenus
}
