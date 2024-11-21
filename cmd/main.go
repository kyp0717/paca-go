package main

import (
  "time"
	"fmt"
	"log"
	"os"

	"paca-go/pkg/marketdata"
	"paca-go/pkg/utils"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	API_KEY_ID := os.Getenv("API_KEY_ID")
	API_SECRET_KEY := os.Getenv("API_SECRET_KEY")
	API_BASE_URL := os.Getenv("API_BASE_URL")

	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    API_KEY_ID,
		APISecret: API_SECRET_KEY,
		BaseURL:   API_BASE_URL,
	})
	acct, err := client.GetAccount()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *acct)

  fmt.Printf(utils.GetTimeInDateFormat().String())

	symbols := []string{"AAPL", "MSFT"}
	// start := time.Date(2024, 11, 20, 30, 0, 0, time.UTC)
	// end := time.Date(2024, 11, 20, 13, 30, 0, 10000000, time.UTC)
  start := time.Date(2021, 8, 9, 13, 30, 0, 0, time.UTC)
  end :=   time.Date(2021, 8, 9, 13, 30, 0, 10000000, time.UTC)

  // marketdata.Trades(symbols,start,end)
  marketdata.Trades(symbols,start,end)
  
}
