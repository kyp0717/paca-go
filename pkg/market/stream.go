package market 

import (
	"log"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)


// type quote string

// PacaQuote data structure stores a list of channels
// Direct incoming message to the proper channel
type PacaStream struct {
  // fan out to stock channel
  QuoteChannels map[string](chan stream.Quote) // fan out 
  QuoteHandlers map[string](func(in <-chan stream.Quote, out chan<- Metric))
  QuoteHandler func(in <-chan stream.Quote, out chan<- Metric)
  // handler which will fan out
  fanOutHandler func(q stream.Quote)

  // process the stock quote in the routine
  // send the metric to the sink channel
  Sink chan Metric// fan in this channel
  // StockMetric Metric // type of calculation or analysis
  // MarketMetric Metric // type of calculation or analysis
  
}

func NewPacaStream() PacaStream{
  m := make(map[string](chan stream.Quote))
  s := make(chan Metric)
  return PacaStream{
    QuoteChannels: m,
    Sink: s,
  }
}

func (qs *PacaStream) exist(q string) bool {
  // check if channel to stream one stock quote exist
  if _, ok := qs.QuoteChannels[q]; ok {
    return true
  } else {
    return false }
}

// create channel for each stock (quote)
// spawn thread to receive (pull) data from channel
func (qs *PacaStream) GetQuote(quote string) {
  // make the channel for the quote
	qs.QuoteChannels[quote] = make(chan stream.Quote)
  // make the function for the quote
  // go processQuote(qs.QuoteChannels[quote], qs.Sink, quoteHandler)
	go qs.QuoteHandler(qs.QuoteChannels[quote],qs.Sink)
}

func (qs *PacaStream) AddQuoteHandler(f func(in <-chan stream.Quote, out chan<- Metric)) {
  qs.QuoteHandler = f
}

// send each stock to its channel
// need to make outside of initiation because need access to map
// in other words, map need to be created first
func (qs *PacaStream) FanOut() {
	qs.fanOutHandler = func(q stream.Quote) {
		qs.QuoteChannels[q.Symbol] <- q
	}
}

func (qs *PacaStream) Subscribe(client stream.StocksClient ) {
  for q := range qs.QuoteChannels {
    if err := client.SubscribeToQuotes(qs.fanOutHandler, q); err != nil {
      log.Fatalf("error during subscribing: %s", err)
    }
  }
}

