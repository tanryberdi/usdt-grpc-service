package handler

import (
	"context"
	"database/sql"

	"usdt-grpc-service/internal/db"
	"usdt-grpc-service/internal/service"
	"usdt-grpc-service/proto"
)

type RateService struct {
	proto.UnimplementedRateServiceServer
	DB *sql.DB
}

// GetRates Implement the method from the proto definition to get the rates
func (s *RateService) GetRates(
	ctx context.Context,
	req *proto.GetRatesRequest,
) (*proto.GetRatesResponse, error) {
	rate, err := service.FetchUSDTRate()
	if err != nil {
		return nil, err
	}

	err = db.SaveRate(s.DB, rate.Ask, rate.Bid, rate.Timestamp)
	if err != nil {
		return nil, err
	}

	return &proto.GetRatesResponse{
		Ask:       rate.Ask,
		Bid:       rate.Bid,
		Timestamp: rate.Timestamp.String(),
	}, nil
}

// HealthCheck Implement the method
func (s *RateService) HealthCheck(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	// Simple health check logic: You can check the DB connection or return a simple status
	if err := s.DB.Ping(); err != nil {
		return &proto.HealthCheckResponse{Status: "Unhealthy"}, err
	}
	return &proto.HealthCheckResponse{Status: "Healthy"}, nil
}
