package market 

import (
	// "context"
	// "fmt"
	"log"
	// "os"

	// "errors"
	// "os/signal"
	// "sync/atomic"
	// "time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	// "github.com/joho/godotenv"
)

type History struct {
	Current  Prices
	Previous Prices
}
type Prices map[string]float64
type PriceList map[string][2]float64

type Last struct {
	symbol    string
	latestq   float64
	previousq float64
}
type Metric interface{
  calculate() 
}

// type quote string

// PacaQuote data structure stores a list of channels
// Direct incoming message to the proper channel
type PacaStream struct {
  QuoteChannels map[string](chan stream.Quote) // fan out 
  handler func(q stream.Quote)
  Sink chan Metric// fan in this channel
  StockMetric Metric // type of calculation or analysis
  MarketMetric Metric // type of calculation or analysis
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

func (qs *PacaStream) GetQuote(quote string) {
		qs.QuoteChannels[quote] = make(chan stream.Quote)
		go qs.Analyze(quote, qs.Metric)
}

// need to make outside of initiation because need access to map
// in other words, map need to be created first
func (qs *PacaStream) FanOut() {
	qs.handler = func(q stream.Quote) {
		qs.QuoteChannels[q.Symbol] <- q
	}
}

func (qs *PacaStream) Subscribe(client stream.StocksClient ) {
  for q := range qs.QuoteChannels {
    if err := client.SubscribeToQuotes(qs.handler, q); err != nil {
      log.Fatalf("error during subscribing: %s", err)
    }
  }
}

// runs forever
// return Metrics which is interface since metrics can be custom
func analyze (symbol string) {
	q := <-qs.QuoteChannels[symbol]
	m := &Last{q.Symbol, q.AskPrice, 0.0}
	for {
	  q := <-qs.QuoteChannels[symbol]
		t.previousq = t.latestq
		t.latestq = q.AskPrice
    qs.Sink<-t
		// fmt.Printf("%v\n", t)
		// time.Sleep(1 * time.Second)
	}
}


func (qs *PacaStream) Process(symbol string, m Metric)  {
	q := <-qs.QuoteChannels[symbol]
	t := &Last{q.Symbol, q.AskPrice, 0.0}
	for {
	  q := <-qs.QuoteChannels[symbol]
		t.previousq = t.latestq
		t.latestq = q.AskPrice
    qs.Sink<-t
		// fmt.Printf("%v\n", t)
		// time.Sleep(1 * time.Second)
	}
}
