package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	// "sync/atomic"
	"time"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)


func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

type History struct {
  Current Prices
  Previous Prices
}
type Prices map[string]float64
type PriceList map[string][2]float64

// var qChan chan stream.Quote
func main() {
  pricesChan := make(chan map[string]float64)
  // pricesChan := make(chan Prices)
  MyStockList := []string{"TSLA", "AAPL"}

  // prices := make(map[string]float64)
  stkprices := make(Prices)

  go func(cc chan map[string]float64) {
    h := History{}
    for {
      prices := <-cc
      // fmt.Println(prices)
      h.Previous = h.Current
      h.Current = prices
      p := PriceList

      for k, pprice := range h.Previous {
        for j, cprice := range h.Current {
          if k == j {
             p[k] = [2]float64{pprice, cprice}
          }
      }

      fmt.Println(h)
      }
      }
    }(pricesChan)

	quoteHandler := func(q stream.Quote) {
      // stk := StockQuote{q.Symbol, q.AskPrice}
      // qChan <- stk
      if len(stkprices) == 2 {
          pricesChan <-stkprices
          stkprices = make(map[string]float64)
          time.Sleep(10 * time.Second)
          // fmt.Println(prices)
        }
      // q := <-cc
      if contains(MyStockList, q.Symbol) {
          stkprices[q.Symbol] = q.AskPrice
        }
      }



	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setting up cancelling upon interrupt
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	go func() {
		<-s
		cancel()
	}()

	// Creating a client that connexts to iex
	c := stream.NewStocksClient("iex")

	if err := c.Connect(ctx); err != nil {
		log.Fatalf("could not establish connection, error: %s", err)
	}
	fmt.Println("established connection")

	// starting a goroutine that checks whether the client has terminated
	go func() {
		err := <-c.Terminated()
		if err != nil {
			log.Fatalf("terminated with error: %s", err)
		}
		fmt.Println("exiting")
		os.Exit(0)
	}()

	// time.Sleep(3 * time.Second)
	// Adding TSLA trade subscription
	if err := c.SubscribeToQuotes(quoteHandler, "TSLA", "AAPL"); err != nil {
		log.Fatalf("error during subscribing: %s", err)
	}
	fmt.Println("subscribed to quotes")

	// and so on...
	time.Sleep(150 * time.Second)
	fmt.Println("we're done")

	// Unsubscribing from AAPL quotes
	if err := c.UnsubscribeFromQuotes("TSLA","AAPL"); err != nil {
		log.Fatalf("error during unsubscribing: %s", err)
	}
	fmt.Println("unsubscribed from quotes")
}

