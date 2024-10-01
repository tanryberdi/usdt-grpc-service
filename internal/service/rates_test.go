package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"usdt-grpc-service/internal"

	"github.com/stretchr/testify/assert"
)

func TestFetchUSDTRate(t *testing.T) {
	// Initialize zap logger
	internal.InitLogger()
	defer internal.Logger.Sync()

	// Mock server to simulate Garantex API response
	mockResponse := `{
        "timestamp": 1638316800,
        "asks": [{"price": "75.0", "volume": "100", "amount": "7500", "factor": "1", "type": "ask"}],
        "bids": [{"price": "74.0", "volume": "100", "amount": "7400", "factor": "1", "type": "bid"}]
    }`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Override the URL to point to the mock server
	url := server.URL

	rate, err := FetchUSDTRate(url)
	assert.NoError(t, err)
	assert.Equal(t, 75.0, rate.Ask)
	assert.Equal(t, 74.0, rate.Bid)
	assert.WithinDuration(t, time.Now(), rate.Timestamp, time.Second)
}
