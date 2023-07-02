package router

import (
	"context"
	"encoding/json"
	"firebond/model"
	"math"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func ApplyRoutes(r *mux.Router) {

	r.HandleFunc("/api/{cryptocurrency}/{fiat}", getExchangeRate).Methods("GET")
	r.HandleFunc("/api/{cryptocurrency}", getAllExchangeRates).Methods("GET")
	r.HandleFunc("/api", getExchangeRateHistory).Methods("GET")
	r.HandleFunc("/api/history/{cryptocurrency}/{fiat}", getEthereumBalance).Methods("GET")

}

var db = client.Database("crypto_data")
var collection = db.Collection("exchange_rates")

func getAllExchangeRates(w http.ResponseWriter, r *http.Request) {
	// Retrieve all exchange rates from the database based on cryptocurrency
	params := mux.Vars(r)
	cryptocurrency := params["cryptocurrency"]

	// Query the database
	var rates []model.ExchangeRateHistory
	cursor, err := collection.Find(context.Background(), map[string]string{"cryptocurrency": cryptocurrency})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var rate model.ExchangeRateHistory
		cursor.Decode(&rate)
		rates = append(rates, rate)
	}

	if len(rates) == 0 {
		http.Error(w, "Exchange rates not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(rates)
}

func getExchangeRate(w http.ResponseWriter, r *http.Request) {
	// Retrieve exchange rate from the database based on cryptocurrency and fiat currency
	params := mux.Vars(r)
	cryptocurrency := params["cryptocurrency"]
	fiatCurrency := params["fiat"]

	// Query the database
	var rate model.ExchangeRateHistory
	err := collection.FindOne(context.Background(), map[string]string{
		"cryptocurrency": cryptocurrency,
		"fiat_currency":  fiatCurrency,
	}).Decode(&rate)

	if err != nil {
		http.Error(w, "Exchange rate not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(rate)
}

func getExchangeRateHistory(w http.ResponseWriter, r *http.Request) {
	// Retrieve exchange rate history from the database based on cryptocurrency and fiat currency
	params := mux.Vars(r)
	cryptocurrency := params["cryptocurrency"]
	fiatCurrency := params["fiat"]

	// Query the database
	var rates []model.ExchangeRateHistory
	cursor, err := collection.Find(context.Background(), map[string]interface{}{
		"cryptocurrency": cryptocurrency,
		"fiat_currency":  fiatCurrency,
		"timestamp":      map[string]int64{"$gt": time.Now().Add(-24 * time.Hour).Unix()},
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var rate model.ExchangeRateHistory
		cursor.Decode(&rate)
		rates = append(rates, rate)
	}

	if len(rates) == 0 {
		http.Error(w, "Exchange rate history not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(rates)
}

func getEthereumBalance(w http.ResponseWriter, r *http.Request) {
	// Retrieve the balance of a specific Ethereum address
	params := mux.Vars(r)
	address := params["address"]

	// Connect to Ethereum node
	client, err := ethclient.Dial(ethereumNode)
	if err != nil {
		http.Error(w, "Error connecting to Ethereum node", http.StatusInternalServerError)
		return
	}

	// Get balance
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		http.Error(w, "Error retrieving balance", http.StatusInternalServerError)
		return
	}

	// Convert balance from wei to ether
	etherBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	json.NewEncoder(w).Encode(model.BalanceResponse{
		Address: address,
		Balance: etherBalance.String(),
	})
}
