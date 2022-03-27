package market

import (
	"context"
	"fmt"
	"time"
	// "fmt"
	"log"
	"os"
	"os/signal"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	"github.com/joho/godotenv"
)

func TestGetPrices(t *testing.T) {
  sink := make(chan LatestQuotes)
	go func() {
    for {
    l := <-sink
    fmt.Println(l)
    }
	}()


	// err := godotenv.Load("~/projects/paca-go/.env")
	err := godotenv.Load()
	if err != nil {
		t.Log("error: file not found")
	}
	// stream.New
	// Creating a client that connexts to iex
	c1 := stream.NewStocksClient("iex")
	// c2 := stream.NewStocksClient("iex")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setting up cancelling upon interrupt
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	go func() {
		<-s
		cancel()
	}()

	if err := c1.Connect(ctx); err != nil {
		log.Fatalf("could not establish connection, error: %s", err)
	}

	// if err := c2.Connect(ctx); err != nil {
	// 	log.Fatalf("could not establish connection, error: %s", err)
	// }
  apple_chan := Qstream(&c1, "AAPL")
  go GetPrices(apple_chan, sink)

  tsla_chan := Qstream(&c1, "TSLA")
  go GetPrices(tsla_chan, sink)

	// and so on...
	time.Sleep(15 * time.Second)
	fmt.Println("we're done")

}
