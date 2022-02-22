package market

import (
	// "context"
	"fmt"
	"log"
	// "os"
	// "errors"
	// "os/signal"
	"time"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

// Todo - test whether channels in struct is active when this object is created
type QuoteStream struct {
	conn     *stream.StocksClient
	stock    string
	Tap  chan stream.Quote  // source data - one per quote
}

type LatestQuotes struct {
	symbol    string
	Latest   float64
	Previous float64
}

// a bunch of side effects
func NewStream(client *stream.StocksClient, s string) *QuoteStream {
  qs := QuoteStream{}
  qs.stock = s
  qs.Tap = make(chan stream.Quote)

  h := func(q stream.Quote) {
	  qs.Tap <- q
  }
  err := (*client).SubscribeToQuotes(h, qs.stock)
	if  err != nil {
		log.Fatalf("error during subscribing: %s", err)
	}
  // fmt.Printf("-- subscribed to quote: %s\n", s)
  return &qs
}

// func extract(qchan <-chan stream.Quote, sink chan<- LatestQuotes) {
func (t *QuoteStream) Extract() {
	q := <-t.Tap
	lq := &LatestQuotes{q.Symbol, q.AskPrice, 0.0}
	for {
		q := <-t.Tap
		// fmt.Println(q)
		lq.Previous = lq.Latest
		lq.Latest= q.AskPrice
		fmt.Printf("%v\n", t)
		// time.Sleep(0.2 * time.Second)
	}
}
func (t *QuoteStream) Stream(sink chan LatestQuotes) {
  go func() {
    q := <- t.Tap
    lq := &LatestQuotes{t.stock, q.AskPrice, 0.0}
    for {
      q := <-t.Tap
      // fmt.Println(q)
      lq.Previous = lq.Latest
      lq.Latest = q.AskPrice
      // fmt.Printf("%v\n", lq)
      // calculate price diff
      time.Sleep(1 * time.Second)
      sink <- *lq
	}
  }()
}


// func (t *QuoteStream) Stream2(sink chan LatestQuotes) {
//   go extract(t.Tap, sink)
// }
func (t *QuoteStream) Stream2() {
  go t.Extract()
}



