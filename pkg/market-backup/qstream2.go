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
type QuoteStream2 struct {
	conn  *stream.StocksClient
	stock string
  // prices LatestQuotes
	// Tap   chan stream.Quote // source data - one per quote
}

type LatestQuotes2 struct {
	// symbol   string
	Latest   float64
	Previous float64
}

func (l *LatestQuotes2) update(q stream.Quote) {
  l.Previous = l.Latest
  l.Latest = q.AskPrice
}

// a bunch of side effects
func New2(s string) *QuoteStream2 {
	qs := QuoteStream2{}
  qs.stock = s
  return &qs
}

func (qs QuoteStream2) Stream(client *stream.StocksClient) (chan stream.Quote) {
  // qs.prices.Latest = 0.0
  // qs.prices.Previous= 0.0
	// qs.stock = s
  // qs.createTap()
  cc := make(chan stream.Quote)

	h := func(q stream.Quote) {
    // fmt.Println(q)
    cc <- q
	}
	err := (*client).SubscribeToQuotes(h, qs.stock)
	if err != nil {
		log.Fatalf("error during subscribing: %s", err)
	}
  return cc 
}

// func (t *QuoteStream2) createTap() {
//   t.Tap = make(chan stream.Quote)
// }

func (t *QuoteStream2) Extract(c <-chan stream.Quote) {
	q := <-c
  l :=LatestQuotes{Latest: q.AskPrice, Previous: 0.0}
	for {
		q := <-c
    l.update(q)
		// fmt.Println(t.prices)
		fmt.Println(l)
	}
}




