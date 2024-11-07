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
	n = "\033[0m"  // Normal text color
	r = "\033[31m" // Red color for price decrease
	g = "\033[32m" // Green color for price increase
)

// PriceData represents the structure of the JSON response from the API.
type PriceData struct {
	Bpi struct {
		USD struct {
			Rate string `json:"rate"`
		} `json:"USD"`
		EUR struct {
			Rate string `json:"rate"`
		} `json:"EUR"`
	} `json:"bpi"`
}

// getPrice fetches the current price of Bitcoin in USD and EUR from Coindesk API.
func getPrice() (float64, float64, error) {
	url := "https://api.coindesk.com/v1/bpi/currentprice.json"
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var d PriceData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, 0, err
	}

	// Remove commas from the price strings and convert them to float64.
	psUSD := strings.ReplaceAll(d.Bpi.USD.Rate, ",", "")
	pUSD, err := strconv.ParseFloat(psUSD, 64)
	if err != nil {
		return 0, 0, err
	}

	psEUR := strings.ReplaceAll(d.Bpi.EUR.Rate, ",", "")
	pEUR, err := strconv.ParseFloat(psEUR, 64)
	if err != nil {
		return 0, 0, err
	}

	return pUSD, pEUR, nil
}

func main() {
	var lastUSD float64 // Store the last fetched USD price

	for {
		pUSD, pEUR, err := getPrice()
		if err != nil {
			fmt.Println("Error fetching price:", err)
			continue
		}

		col := n        // Default color (normal)
		var pct float64 // Percentage change between prices

		if lastUSD != 0 { // If we have a previous USD price to compare with
			diff := pUSD - lastUSD
			pct = (diff / lastUSD) * 100

			if diff > 0 {
				col = g // Price increased: use green color
			} else if diff < 0 {
				col = r // Price decreased: use red color
			}
		}

		fmt.Printf("BTC Price: %s$%.2f (%.2f%%) | €%.2f%s", col, pUSD, pct, pEUR, n)

		lastUSD = pUSD // Update last USD price for next iteration

		time.Sleep(10 * time.Second) // Wait for 10 seconds before refreshing
	}
}
