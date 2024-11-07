package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	norm      = "\033[0m"   // Normal text color
	red       = "\033[31m"  // Red color for price decrease
	green     = "\033[32m"  // Green color for price increase
	clearLine = "\033[2K\r" // Clears the current line and moves the cursor to the start
)

// PriceData represents the structure of the JSON response from the API.
type PriceData struct {
	Bpi struct {
		USD struct {
			Rate string `json:"rate"`
		} `json:"USD"`
	} `json:"bpi"`
}

// getPrice fetches the current price of a cryptocurrency from Coindesk API.
func getPrice(crypto string) (float64, error) {
	url := fmt.Sprintf("https://api.coindesk.com/v1/bpi/currentprice/%s.json", crypto)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data PriceData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	// Remove commas from the price string and convert it to float64.
	priceStr := strings.ReplaceAll(data.Bpi.USD.Rate, ",", "")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func main() {
	var lastPrice float64 // Store the last fetched price

	for {
		price, err := getPrice("BTC")
		if err != nil {
			fmt.Println("Error fetching price:", err)
			continue
		}

		color := norm         // Default color (normal)
		var pctChange float64 // Percentage change between prices

		if lastPrice != 0 { // If we have a previous price to compare with
			change := price - lastPrice
			pctChange = (change / lastPrice) * 100

			if change > 0 {
				color = green // Price increased: use green color
			} else if change < 0 {
				color = red // Price decreased: use red color
			}
		}

		fmt.Printf("%s%sBTC Price: $%.2f (%.2f%%)%s", clearLine, color, price, pctChange, norm)

		lastPrice = price // Update lastPrice for next iteration

		time.Sleep(30 * time.Second) // Wait for 30 seconds before refreshing
	}
}
