package market

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	"github.com/joho/godotenv"
)

// define metric type in another file (see trend.go file)


func TestAlgo(t *testing.T) {
  // Defines 2 types for algo
  // type 1 - what is your metric you derive from quotes
  // type 2 - aggregate the metrics into a decision

  stxlist := []string{"AAPL","TSLA"}
  algo := NewAlgo(stxlist, "AMD")

	algo.HandleQuote = func(in <-chan stream.Quote, out chan<- Metric) {
		q := <-in
		tr := StockTrend{}
		tr.Init(q)
		for {
			q := <-in
			tr.Update(q)
			tr.TransformQuote()
			out <- &tr
		}
	}

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
	qs := NewPacaStream()
	qhandler := func(in <-chan stream.Quote, out chan<- Metric) {
		q := <-in
		tr := StockTrend{}
		tr.Init(q)
		for {
			q := <-in
			tr.Update(q)
			tr.GetTrend()
			out <- &tr
		}
	}
	qs.AddQuoteHandler(qhandler)
	// create a channel to receive stock quote and routine to read from channel
	qs.GetQuote("AAPL")
	qs.GetQuote("TSLA")
	// map (which holds channels) and threads have to exist first before adding handler
	// this handler redirect by fanning out to mulitple channels
	qs.FanOut()

	// subscribe will kick of the stream by connecting to Paca
	qs.Subscribe(c)

	go func() {
		for {
			a := <-qs.Sink
			fmt.Println(a)
		}
	}()

	// and so on...
	time.Sleep(15 * time.Second)
	fmt.Println("we're done")

}
