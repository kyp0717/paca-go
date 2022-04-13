package algo

import (
	"context"
	"time"
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	"github.com/joho/godotenv"
)

func TestTrend(t *testing.T) {
	// err := godotenv.Load("~/projects/paca-go/.env")
	err := godotenv.Load()
	if err != nil {
		t.Log("error: file not found")
	}
	// stream.New
	// Creating a client that connexts to iex
	c := stream.NewStocksClient("iex")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// setting up cancelling upon interrupt
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	go func() {
		<-s
		cancel()
	}()

	if err := c.Connect(ctx); err != nil {
		log.Fatalf("could not establish connection, error: %s", err)
	}


  // create the map
  // add sink channel
  dp := NewQuoteDispatcher()
  appl := dp.GetQuote(c,"APPL")
  tsla := dp.GetQuote(c,"TSLA")

   // Compute trend and send to sink channel
  // sectortrend := make(chan StockTrend)
  sectortrend := make(SectorTrend)
  amd.Compute(sectortrend)
  tsla.Compute(sectortrend)

  // calculate sector status based on stock trends fanning into channel 
  // return a channel of 
  sectorStatus := sectortrend.Compute()

  long_amd := NewTrade("AMD", Long)
  long_amd.BenchMark(sectorStutus)
  long_amd.Enter()

  go func() {
    for {
    v := <- sectortrend
    fmt.Println(v)
    }
  }()
	// and so on...
	time.Sleep(2 * time.Second)
	fmt.Println("we're done")

}
