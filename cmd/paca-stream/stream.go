package main

import (
  "os"
	"fmt"
	"github/kyp0717/paca-go/db"
	"log"
  "time"
  "bufio"
  "strings"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	// "github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/joho/godotenv"
)

func init() {
  // load the alpaca key and secret env var
	if err := godotenv.Load("env/.env.alpaca"); err != nil {
		log.Fatal("Error in loading .env alpaca file.")
		fmt.Println("Error in loading .env alpaca file.")
	}
  // ensure that postgres is running
	if err := godotenv.Load("env/.env.postgres"); err != nil {
		log.Fatal("Error in loading .env postgres file.")
		fmt.Println("Error in loading .env postgres file.")
	}
  fmt.Println("Connecting to postgres...")
  db.ConnectDB()
}

func main() {
  reader := bufio.NewReader(os.Stdin)

	fmt.Print("** Is Postgres running? (y/n): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Clean and normalize the input
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "n" {
		fmt.Println("...Exiting. Have a great day!")
   	os.Exit(0) // Exit the program successfully
	}

	// API_KEY_ID := os.Getenv("API_KEY_ID")
	// API_SECRET_KEY := os.Getenv("API_SECRET_KEY")

	// client := alpaca.NewClient(alpaca.ClientOpts{
	// 	APIKey:    API_KEY_ID,
	// 	APISecret: API_SECRET_KEY,
	// })

	symbols := []string{"AAPL", "MSFT"}
	// start := time.Date(2024, 11, 20, 30, 0, 0, time.UTC)
	// end := time.Date(2024, 11, 20, 13, 30, 0, 10000000, time.UTC)
  start := time.Date(2021, 8, 9, 13, 30, 0, 0, time.UTC)
  end :=   time.Date(2021, 8, 9, 13, 30, 0, 10000000, time.UTC)


	multiTrades, err := marketdata.GetMultiTrades(symbols, marketdata.GetTradesRequest{
		Start: start,
		End:   end,
	})
	if err != nil {
		panic(err)
	}
	for symbol, trades := range multiTrades {
		fmt.Println(symbol + " trades:")
		for _, trade := range trades {
			fmt.Printf("%+v\n", trade)
		}
	}

}
