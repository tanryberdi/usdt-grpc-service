package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"usdt-grpc-service/internal"
	"usdt-grpc-service/internal/db"
	"usdt-grpc-service/internal/handler"
	"usdt-grpc-service/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	internal.InitLogger()
	defer internal.Logger.Sync()

	host := os.Getenv("DB_HOST")
	connStr := "postgres://tanryberdi:tanryberdi@" + host + ":5432/testdb?sslmode=disable"
	dbConn, err := db.ConnectToDB(connStr)
	if err != nil {
		internal.Logger.Fatal("failed to connect to the database", zap.Error(err))
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		internal.Logger.Fatal("failed to listen", zap.Error(err))
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	rateService := &handler.RateService{DB: dbConn}
	proto.RegisterRateServiceServer(grpcServer, rateService)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		internal.Logger.Fatal("failed to serve", zap.Error(err))
	}

	// Gracefully shutdown the server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			internal.Logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	// Capture signals got shut down the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}
