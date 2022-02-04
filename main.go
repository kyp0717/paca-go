package main

import (
	"context"
	"fmt"
	"log"
	"os"
  // "errors"
	"os/signal"
	// "sync/atomic"
	"time"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)


type History struct {
  Current Prices
  Previous Prices
}
type Prices map[string]float64
type PriceList map[string][2]float64

type Last struct {
  symbol string
  latestq float64
  previousq float64
}
type QChanMap map[string](chan stream.Quote)

func creatQChanMap(ls []string) QChanMap {
  qMap := make(map[string](chan stream.Quote))
  for _, v := range ls {
    qMap[v] = make(chan stream.Quote)
  }
  return qMap
}

func creatQRoutines (QChans QChanMap ) {
  for _, cc := range QChans {
    go track(cc)
 }
}

func track(qchan <-chan stream.Quote) {
  q := <- qchan
  t := &Last{q.Symbol, q.AskPrice, 0.0}
  for {
    q := <-qchan
    fmt.Println(q)
    t.previousq = t.latestq
    t.latestq = q.AskPrice
    fmt.Printf("%v\n", t)
    time.Sleep(3*time.Second)
  }
}


func main() {
  MyStockList := []string{"TSLA", "AAPL"}

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

  qmap := creatQChanMap(MyStockList)
  creatQRoutines(qmap)

  qHandler := func(q stream.Quote) {
    qmap[q.Symbol] <- q
  }
	// time.Sleep(3 * time.Second)
	// Adding TSLA trade subscription
	if err := c.SubscribeToQuotes(qHandler, "TSLA", "AAPL"); err != nil {
		log.Fatalf("error during subscribing: %s", err)
	}
	fmt.Println("subscribed to quotes")

	// and so on...
	time.Sleep(40 * time.Second)
	fmt.Println("we're done")

	// Unsubscribing from AAPL quotes
	if err := c.UnsubscribeFromQuotes("TSLA","AAPL"); err != nil {
		log.Fatalf("error during unsubscribing: %s", err)
	}
	fmt.Println("unsubscribed from quotes")
}

