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

func TestStream(t *testing.T) {
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

	// apple := NewStream(&c, "AAPL")
	// apple.Stream(out)

	tsla := NewStream(&c, "TSLA")
  go tsla.Extract()
	// tsla.Stream2()
 
	apple := NewStream(&c, "AAPL")
  go apple.Extract()
	// apple.Stream2()

	// and so on...
	time.Sleep(15 * time.Second)
	fmt.Println("we're done")

}
