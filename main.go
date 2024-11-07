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
	} `json:"bpi"`
}

// getPrice fetches the current price of a cryptocurrency from Coindesk API.
func getPrice(c string) (float64, error) {
	url := fmt.Sprintf("https://api.coindesk.com/v1/bpi/currentprice/%s.json", c)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var d PriceData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	// Remove commas from the price string and convert it to float64.
	ps := strings.ReplaceAll(d.Bpi.USD.Rate, ",", "")
	p, err := strconv.ParseFloat(ps, 64)
	if err != nil {
		return 0, err
	}

	return p, nil
}

func main() {
	var l float64 // Store the last fetched price

	for {
		p, err := getPrice("BTC")
		if err != nil {
			fmt.Println("Error fetching price:", err)
			continue
		}

		col := n        // Default color (normal)
		var pct float64 // Percentage change between prices

		if l != 0 { // If we have a previous price to compare with
			diff := p - l
			pct = (diff / l) * 100

			if diff > 0 {
				col = g // Price increased: use green color
			} else if diff < 0 {
				col = r // Price decreased: use red color
			}
		}

		fmt.Printf("BTC Price: %s$%.2f (%.2f%%)%s\n", col, p, pct, n)

		l = p // Update lastPrice for next iteration

		time.Sleep(10 * time.Second) // Wait for 10 seconds before refreshing
	}
}
