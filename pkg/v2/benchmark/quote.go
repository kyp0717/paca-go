package benchmark

import (
	"log"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)
type Symbol=string

// quotes are fan out to channels
type Dispatcher struct{
  // fan out to stock channel
  Channels map[Symbol]chan stream.Quote // fan out 
  // handler which will fan out
  Handler func(q stream.Quote)
}

type PriceMove int64
const (
  Up PriceMove = iota
  Down
  None
  )


type PriceHistory struct {
	Symbol   
	Latest   float64
	Previous float64
}

type PriceChange struct {
	Symbol   
  PriceMove
}

func Dispatch() Dispatcher {
  qc:= make(map[Symbol](chan stream.Quote))
  handler := func(q stream.Quote) {
		qc[q.Symbol] <- q
    // fmt.Println(q)
	}
  return Dispatcher{
    Channels: qc,
    Handler: handler,
  }
}

func (d *Dispatcher) GetQuote(client stream.StocksClient,s Symbol) QuoteChan{
  d.Channels[s] = make(chan stream.Quote)
  if err := client.SubscribeToQuotes(d.Handler, s); err != nil {
      log.Fatalf("error during subscribing: %s", err)
  }
  return d.Channels[s]
}

func (t PriceMove) String() string {
    // declare an array of strings
    // ... operator counts how many
    // items in the array (7)
    names := [...]string{
        "Up", 
        "Down", 
        "None"}
    return names[t]
}

// determine price change for each stock and fan in to channel
type QuoteChan chan stream.Quote
type PriceChangeChan chan PriceChange
func (qs QuoteChan) Compute(sink PriceChangeChan)  {
  go func() {
  ph:=PriceHistory{}
  q:=<-qs
  ph.Init(q)
  for {
    q:=<-qs
    ph.Update(q)
    st := ph.Transform()
    sink<-st
   }
  }()
}

func (st *PriceHistory) Init(q stream.Quote) {
	st.Symbol = q.Symbol
	st.Latest = q.AskPrice
	st.Previous = 0.0
}

func (st *PriceHistory) Update(q stream.Quote) {
	st.Latest = q.AskPrice
	st.Previous = st.Latest
}

func (st *PriceHistory) Transform() PriceChange {
	diff := st.Latest - st.Previous
	switch {
    case diff > 0: return PriceChange{st.Symbol,Up} 
    case diff < 0: return PriceChange{st.Symbol,Down}
    default: return PriceChange{st.Symbol,None} 
	}
}
