package market 

import (
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)


type StockTrend struct {
	Symbol    string
	Latestq   float64
	Previousq float64
  Trend 
}
// type Stock string

// type StockTrend struct {
//   PriceHistory
// }

func (st *StockTrend) Init(q stream.Quote) {
  st.Symbol = q.Symbol
  st.Latestq= q.AskPrice
  st.Previousq= 0.0
}

func (st *StockTrend) Update(q stream.Quote) {
  st.Latestq= q.AskPrice
  st.Previousq= st.Latestq
}

func (st *StockTrend) GetSymbol() string{
  return st.Symbol
}


func (st *StockTrend) GetTrend() Trend{
  return st.Trend
}

func (st *StockTrend) TransformQuote() {
  diff := st.Latestq - st.Previousq
  if diff > 0 {
      st.Trend = Up
  } else {
    st.Trend = Down
  }
}


