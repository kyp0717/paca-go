package benchmark

import (
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
  "paca-go/pkg/v1/quote"
)

type Trend int64
const (
  Rally Trend= iota
  SellOff
  Random
  Unknown
  )

type SectorTrendChan chan Trend 
type MarketTrendChan chan Trend 


// determine  status for each stock
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




