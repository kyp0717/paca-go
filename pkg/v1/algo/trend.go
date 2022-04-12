package algo

import (
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

type Trend int64
const (
  Up Trend = iota
  Down
  None
  )

type StockTrend struct {
  Symbol
  Trend
}

func (t Trend) String() string {
    // declare an array of strings
    // ... operator counts how many
    // items in the array (7)
    names := [...]string{
        "Up", 
        "Down", 
        "None"}
    return names[t]
}

// type StockTrend chan<- Feature
type PriceHistory struct {
	Symbol   
	Latest   float64
	Previous float64
}

type QuoteStream chan stream.Quote

func (qs QuoteStream) Compute(sink chan<- StockTrend)  {
  go func() {
  sh:=PriceHistory{}
  q:=<-qs
  sh.Init(q)
  for {
    q:=<-qs
    sh.Update(q)
    st := sh.Transform()
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

func (st *PriceHistory) Transform() StockTrend {
	diff := st.Latest - st.Previous
	switch {
    case diff > 0: return StockTrend{st.Symbol,Up} 
    case diff < 0: return StockTrend{st.Symbol,Down}
    default: return StockTrend{st.Symbol,None} 
	}
}




