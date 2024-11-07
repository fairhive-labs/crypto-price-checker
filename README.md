# Crypto Price Checker

A simple command-line tool to fetch and display the current price of Bitcoin (BTC) in both USD and EUR. The tool also includes a progress bar that updates every 10 seconds, reflecting the wait period before fetching new prices.

> Feel free to contribute or report issues!

## Features

- Fetches real-time Bitcoin (BTC) prices from the CoinDesk API.
- Displays BTC prices in both **USD** and **EUR**.
- Color-coded price changes:
  - **Green**: Price increased.
  - **Red**: Price decreased.
  - **Normal**: No change.
- Progress bar that updates every second during the 10-second wait period between price refreshes.

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/fairhive-labs/crypto-price-checker.git
```

### 2. Navigate to the project directory
```bash
cd crypto-price-checker
```

### 3. Run the program
```bash
go run main.go
```

## Usage
Once you run the program, it will continuously fetch and display Bitcoin prices in USD and EUR every 10 seconds. The progress bar will show how much time is left until the next price update.

Example output:
```text
PRICE CHECKER started...
BTC Price: $76445.55 (0.08%) | 70904.17€ [in green if price increased]
BTC Price: $76404.69 (-0.05%) | 70866.26€ [in red if price decreased]
[##########]
```

## Dependencies
This project uses Go's standard library, so no external dependencies are required.

## API Used
[CoinDesk API](https://www.coindesk.com/) for fetching real-time Bitcoin prices.

## License
This project is licensed under the MIT License - see the LICENSE file for details. 
