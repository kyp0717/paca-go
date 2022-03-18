package market

import (
	// "context"
	// "fmt"
	"fmt"
	"log"
	// "os"
	// "errors"
	// "os/signal"
	// "time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

// Todo - test whether channels in struct is active when this object is created
type QuoteStream struct {
	conn  *stream.StocksClient
	stock string
  prices LatestQuotes
	Tap   chan stream.Quote // source data - one per quote
}

type LatestQuotes struct {
	// symbol   string
	Latest   float64
	Previous float64
}

func (l *LatestQuotes) update(q stream.Quote) {
  l.Previous = l.Latest
  l.Latest = q.AskPrice
}

// a bunch of side effects
func New(client *stream.StocksClient, s string) *QuoteStream {
	qs := QuoteStream{}
  qs.prices.Latest = 0.0
  qs.prices.Previous= 0.0
	qs.stock = s
  qs.createTap()

	h := func(q stream.Quote) {
    // fmt.Println(q)
		qs.Tap <- q
    // cc <- q
	}
	err := (*client).SubscribeToQuotes(h, qs.stock)
	if err != nil {
		log.Fatalf("error during subscribing: %s", err)
	}
  return &qs 
}

func (t *QuoteStream) createTap() {
  t.Tap = make(chan stream.Quote)
}

func (t *QuoteStream) Extract() {
	for {
		q := <-t.Tap
    t.prices.update(q)
    fmt.Println(t.stock, ": ", t.prices)
		// fmt.Println(q)
	}
}




