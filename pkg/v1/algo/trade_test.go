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
)

func TestTrade(t *testing.T) {
  techSector := []string{"AAPL","TSLA"}
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
  dp := NewDispatcher()
  long_amd := NewTrade("AMD", Long, dp)
  long_amd.AddClient(restclient, streamclient)
  // check AMD stock movment
  sectorStatusChan, _ := long_amd.Benchmark(techSector)
  ok, err := long_amd.Evaluate(sectorStatusChan)
    if ok {
      // Initiate the trade by entering the position 
      long_amd.Enter(sectorStatusChan)
    }

}
