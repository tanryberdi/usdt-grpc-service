package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type GarantexResponse struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

type Rate struct {
	Ask       float64
	Bid       float64
	Timestamp time.Time
}

func FetchUSDTRate() (Rate, error) {
	url := "https://garantex.org/api/v2/depth?market={usdtrub}"
	resp, err := http.Get(url)
	if err != nil {
		return Rate{}, err
	}
	defer resp.Body.Close()

	var data GarantexResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Rate{}, err
	}

	askPrice := data.Asks[0][0]
	bidPrice := data.Bids[0][0]

	// Convert string to float64
	ask, _ := strconv.ParseFloat(askPrice, 64)
	bid, _ := strconv.ParseFloat(bidPrice, 64)

	return Rate{
		Ask:       ask,
		Bid:       bid,
		Timestamp: time.Now(),
	}, nil
}
