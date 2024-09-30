package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
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

func FetchUSDTRate() (Rate, error) {
	url := "https://garantex.org/api/v2/depth?market=usdtrub"
	resp, err := http.Get(url)
	if err != nil {
		return Rate{}, err
	}
	defer resp.Body.Close()

	// unmarshall the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Rate{}, err
	}

	var data GarantexResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
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
