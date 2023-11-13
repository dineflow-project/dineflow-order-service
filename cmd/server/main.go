package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"dineflow-order-service/config"
	"dineflow-order-service/gapi"
	"dineflow-order-service/pb"
	"dineflow-order-service/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	authCollection *mongo.Collection

	// ? Create the Order Variables
	orderService services.OrderService
	// OrderController      controllers.OrderController
	orderCollection *mongo.Collection
	// OrderRouteController routes.OrderRouteController
)

func init() {

	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.EnvMongoDBURI())
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// ? Instantiate the Constructors
	orderCollection = mongoclient.Database("order_db").Collection("orders")
	orderService = services.NewOrderService(orderCollection, ctx)
	// OrderController = controllers.NewOrderController(orderService)
	// OrderRouteController = routes.NewOrderControllerRoute(OrderController)

	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	// startGinServer(config)
	startGrpcServer()
}

func startGrpcServer() {

	orderServer, err := gapi.NewGrpcOrderServer(orderCollection, orderService)
	if err != nil {
		log.Fatal("cannot create grpc orderServer: ", err)
	}

	grpcServer := grpc.NewServer()

	// ? Register the Order gRPC service
	pb.RegisterOrderServiceServer(grpcServer, orderServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.EnvGRPCServerAddress())
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}

// func startGinServer(config config.Config) {

// 	corsConfig := cors.DefaultConfig()
// 	corsConfig.AllowOrigins = []string{config.Origin}
// 	corsConfig.AllowCredentials = true

// 	server.Use(cors.New(corsConfig))

// 	router := server.Group("/api")
// 	router.GET("/healthchecker", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok"})
// 	})

// 	// ? Order Route
// 	OrderRouteController.OrderRoute(router)
// 	log.Fatal(server.Run(":" + config.Port))
// }
