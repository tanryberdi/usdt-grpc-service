package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"usdt-grpc-service/internal/db"
	"usdt-grpc-service/internal/handler"
	"usdt-grpc-service/proto"

	"google.golang.org/grpc"
)

func main() {
	connStr := "postgres://tanryberdi:tanryberdi@localhost:5432/test?sslmode=disable"
	dbConn, err := db.ConnectToDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	rateService := &handler.RateService{DB: dbConn}
	proto.RegisterRateServiceServer(grpcServer, rateService)

	// Gracefully shutdown the server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Capture signals to shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}
