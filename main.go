package main

import (
	"encoding/json"
	"fmt"
)

type StockQuote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

func main() {
	// Sample JSON data representing a stock quote
	jsonData := `
	{
		"01. symbol": "AAPL",
		"02. open": "150.50",
		"03. high": "155.00",
		"04. low": "148.50",
		"05. price": "152.75",
		"06. volume": "1000000",
		"07. latest trading day": "2023-08-21",
		"08. previous close": "150.00",
		"09. change": "2.75",
		"10. change percent": "1.83%"
	}
	`

	var quote StockQuote
	err := json.Unmarshal([]byte(jsonData), &quote)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Symbol:", quote.Symbol)
	fmt.Println("Open:", quote.Open)
	fmt.Println("High:", quote.High)
	fmt.Println("Low:", quote.Low)
	fmt.Println("Price:", quote.Price)
	fmt.Println("Volume:", quote.Volume)
	fmt.Println("Latest Trading Day:", quote.LatestTradingDay)
	fmt.Println("Previous Close:", quote.PreviousClose)
	fmt.Println("Change:", quote.Change)
	fmt.Println("Change Percent:", quote.ChangePercent)
}
