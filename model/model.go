package model

import "go.mongodb.org/mongo-driver/mongo"

type ExchangeRateHistory struct {
	ID             int64   `json:"id"`
	Cryptocurrency string  `json:"cryptocurrency"`
	FiatCurrency   string  `json:"fiat_currency"`
	Rate           float64 `json:"rate"`
	Timestamp      int64   `json:"timestamp"`
}

type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

var (
	client       *mongo.Client
	db           *mongo.Database
	collection   *mongo.Collection
	ethereumNode = "http://localhost:8545" // Change this to your Ethereum node URL
)
