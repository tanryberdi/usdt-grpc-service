package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"usdt-grpc-service/internal"

	"go.uber.org/zap"
)

type GarantexResponse struct {
	Timestamp int64 `json:"timestamp"`
	Asks      []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Amount string `json:"amount"`
		Factor string `json:"factor"`
		Type   string `json:"type"`
	} `json:"asks"`

	Bids []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Amount string `json:"amount"`
		Factor string `json:"factor"`
		Type   string `json:"type"`
	} `json:"bids"`
}

type Rate struct {
	Ask       float64
	Bid       float64
	Timestamp time.Time
}

// FetchUSDTRate fetches the USDT to RUB rate from Garantex
func FetchUSDTRate() (Rate, error) {
	url := "https://garantex.org/api/v2/depth?market=usdtrub"
	resp, err := http.Get(url)
	if err != nil {
		internal.Logger.Error("failed to fetch USDT rate", zap.Error(err))
		return Rate{}, err
	}
	defer resp.Body.Close()

	// unmarshall the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.Logger.Error("failed to read response body", zap.Error(err))
		return Rate{}, err
	}

	var data GarantexResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		internal.Logger.Error("failed to unmarshal response", zap.Error(err))
		return Rate{}, err
	}

	askPrice := data.Asks[0].Price
	bidPrice := data.Bids[0].Price

	// Convert string to float64
	ask, _ := strconv.ParseFloat(askPrice, 64)
	bid, _ := strconv.ParseFloat(bidPrice, 64)

	return Rate{
		Ask:       ask,
		Bid:       bid,
		Timestamp: time.Now(),
	}, nil
}
