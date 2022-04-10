package algo

import (
	"log"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)
type Symbol string
type Quote struct {
  // Symbol
  stock string
  // symbol stream.Quote
  in <-chan stream.Quote
  *Dispatcher
}

type Dispatcher struct {
  // fan out to stock channel
  QuoteChannels map[string](chan stream.Quote) // fan out 
  // handler which will fan out
  dispatch func(q stream.Quote)
}

func NewDispatcher() Dispatcher{
  return Dispatcher{
    QuoteChannels: make(map[string]chan stream.Quote),
  }
}

// send each stock to its channel
// need to make outside of initiation because need access to map
// in other words, map need to be created first
func (qh *Dispatcher) Dispatch() func(q stream.Quote){
  handler := func(q stream.Quote) {
		qh.QuoteChannels[q.Symbol] <- q
	}
  return handler
}

func NewQuote(s string, q *Dispatcher) Quote {
  in := make(<-chan stream.Quote)
  return Quote{in: in, Dispatcher: q}
}

func (q Quote) Subscribe(client stream.StocksClient) {
  if err := client.SubscribeToQuotes(q.Dispatch(), q.stock); err != nil {
      log.Fatalf("error during subscribing: %s", err)
 }
}

func (q Quote) Compute(sink chan<- Feature)  {
  var trend Feature
  qq := <-q.in
  tr = Trend{}
  tr.Init(qq)
  for {
    q := <-in
    tr.Update(q)
    tr.GetTrend()
    out<-&tr
  }
}

func (q Quote) Stream() {
  go q.Compute(sink)
}

