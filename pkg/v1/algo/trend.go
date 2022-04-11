package algo

import "github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"


type Trend int64
const (
  Up Trend = iota
  Down
  Same
  )
// type Trend struct {}

// type StockTrend chan<- Feature
type StockHistory struct {
	Symbol    string
	Latestq   float64
	Previousq float64
}

type StockTrend=Trend
// type StockTrend struct {
//   Trend
//   sink chan Trend
// }


// type Stock string

// type StockTrend struct {
//   PriceHistory
// }

func (st *StockHistory) Init(q stream.Quote) {
	st.Symbol = q.Symbol
	st.Latestq = q.AskPrice
	st.Previousq = 0.0
}

func (st *StockHistory) Update(q stream.Quote) {
	st.Latestq = q.AskPrice
	st.Previousq = st.Latestq
}

func (st *StockHistory) GetSymbol() string {
	return st.Symbol
}


func (st *StockHistory) Compute() StockTrend {
	diff := st.Latestq - st.Previousq
	switch {
    case diff > 0: return Up 
    case diff < 0: return Down
    default: return Same
	}
}


func (st StockTrend) Sink(c chan<- Trend)  {
   c<-st
}
