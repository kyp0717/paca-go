package algo

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	"github.com/joho/godotenv"

  "paca-go/pkg/v1/trade"
  "paca-go/pkg/v1/quote"
  "paca-go/pkg/v1/benchmark"
)

func TestAlgo(t *testing.T) {
	// err := godotenv.Load("~/projects/paca-go/.env")
	err := godotenv.Load()
	if err != nil {
		t.Log("error: file not found")
	}
	// stream.New
	// Creating a client that connexts to iex
	streamclient := stream.NewStocksClient("iex")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// setting up cancelling upon interrupt
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	go func() {
		<-s
		cancel()
	}()

	if err := streamclient.Connect(ctx); err != nil {
		log.Fatalf("could not establish connection, error: %s", err)
	}
	restclient := alpaca.NewClient(alpaca.ClientOpts{
		// Alternatively you can set your key and secret using the
		// APCA_API_KEY_ID and APCA_API_SECRET_KEY environment variables
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
		BaseURL:   "https://paper-api.alpaca.markets",
	})

  // create the map
  // add sink channel
  dispatcher := quote.Dispatch()
  tech_bm := benchmark.New(dispatcher,)
  long_amd := NewTrade("AMD", Long, restclient)
  done, err := long_amd.Enter(techbm)
  if err!=nil  {
    // Initiate the trade by entering the position 
    fmt.Println("Bad trade.  Don't trade")
  }
  <-done

}
