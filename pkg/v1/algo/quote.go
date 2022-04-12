package algo

import (
	// "fmt"
	"log"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)
type Symbol=string
type Quote struct {
  Symbol
  // *Dispatcher
}
// quotes are fan out to channels
type QuoteDispatcher struct {
  // fan out to stock channel
  Channels map[Symbol](chan stream.Quote) // fan out 
  // handler which will fan out
  Handler func(q stream.Quote)
}

func NewQuoteDispatcher() QuoteDispatcher{
  qc:= make(map[Symbol]chan stream.Quote)
  handler := func(q stream.Quote) {
		qc[q.Symbol] <- q
    // fmt.Println(q)
	}
  return QuoteDispatcher{
    Channels: qc,
    Handler: handler,
  }
}

// func (d *QuoteDispatcher) GetQuote(client stream.StocksClient,s Symbol) (chan stream.Quote){
//   d.Channels[s] = make(chan stream.Quote)
//   if err := client.SubscribeToQuotes(d.Handler, s); err != nil {
//       log.Fatalf("error during subscribing: %s", err)
//   }
//   return d.Channels[s] 
// }

func (d *QuoteDispatcher) GetQuote(client stream.StocksClient,s Symbol) QuoteStream{
  d.Channels[s] = make(chan stream.Quote)
  if err := client.SubscribeToQuotes(d.Handler, s); err != nil {
      log.Fatalf("error during subscribing: %s", err)
  }
  return d.Channels[s] 
}
// pass work now to data structure in another file
// there is no compute here since we are not transforming data
// pure ETL process only
