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

type StockList struct {
  Symbols []string
  LatestPrice []float64 
  PrevPrice []float64 
}

type StockQuote struct {
  Symbol string
  Latestquote float64
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
// type Prices map[string]float64

// var qChan chan stream.Quote
func main() {
  qChan := make(chan StockQuote)
  pricesChan := make(chan map[string]float64)
  MyStockList := []string{"TSLA", "AAPL"}

  prices := make(map[string]float64)
  go func(cc chan StockQuote) {
    for {
    if len(prices) == 2 {
        pricesChan <- prices
        prices = make(map[string]float64)
        // fmt.Println(prices)
      }
    q := <-cc
    // fmt.Println(q)
    if contains(MyStockList, q.Symbol) {
        prices[q.Symbol] = q.Latestquote
      }
    // fmt.Println(prices)
     }
  }(qChan)

  go func(cc chan map[string]float64) {
    for {
      prices := <-cc
      fmt.Println(prices)
      }
  }(pricesChan)

	quoteHandler := func(q stream.Quote) {
      stk := StockQuote{q.Symbol, q.AskPrice}
      qChan <- stk
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

	time.Sleep(3 * time.Second)
	// Unsubscribing from AAPL quotes
	if err := c.UnsubscribeFromQuotes("TSLA","AAPL"); err != nil {
		log.Fatalf("error during unsubscribing: %s", err)
	}
	fmt.Println("unsubscribed from quotes")

	// and so on...
	time.Sleep(100 * time.Second)
	fmt.Println("we're done")
}

