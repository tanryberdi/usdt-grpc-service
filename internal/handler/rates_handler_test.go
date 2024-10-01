package handler

import (
	"context"
	"testing"
	"time"

	"usdt-grpc-service/internal/service"
	"usdt-grpc-service/proto"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRateService_GetRates(t *testing.T) {
	// Mock the database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the FetchUSDTRate function
	mockFetchUSDTRate := func(url string) (service.Rate, error) {
		return service.Rate{
			Ask:       75.0,
			Bid:       74.0,
			Timestamp: time.Now(),
		}, nil
	}

	// Mock the SaveRate function
	mock.ExpectExec("INSERT INTO rates").WithArgs(75.0, 74.0, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	// Create the RateService
	rateService := &RateService{
		DB:            db,
		FetchUSDTRate: mockFetchUSDTRate,
	}

	// Create a context
	ctx := context.Background()

	// Call the GetRates method
	req := &proto.GetRatesRequest{}
	res, err := rateService.GetRates(ctx, req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 75.0, res.Ask)
	assert.Equal(t, 74.0, res.Bid)
}

func TestRateService_HealthCheck(t *testing.T) {
	// Mock the database
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)
	defer db.Close()

	// Mock the Ping function
	mock.ExpectPing().WillReturnError(nil)

	// Create the RateService
	rateService := &RateService{DB: db}

	// Create a context
	ctx := context.Background()

	// Call the HealthCheck method
	req := &proto.HealthCheckRequest{}
	res, err := rateService.HealthCheck(ctx, req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "Healthy", res.Status)
}
