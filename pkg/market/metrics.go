package market 
	
import (
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)
// What is a metric?
// Metric can be a trend, etc
type Metric interface {
  GetTrend()  Trend
  HandleQuote(q stream.Quote) Trend
  Aggregate(chan<- Trend) (<-chan Trend)
}

type Trend int64
const (
  Up Trend = iota
  Down
  Same
  )

// type Analysis interface {
//   TransformQuote()
//   GetTrend()  Trend
//   GetSymbol() string
// }

type Analysis struct {
  PacaStream
  StockHistory
  Stock
  Sector
  Market
}

type StockHistory struct {
  Symbol string
	Latestq   float64
	Previousq float64
}
// type Stock string
// type Symbol string
type Stock Symbol
type Sector map[Stock]Trend
type Market []Sector

func (st *StockHistory) Init(q stream.Quote) {
  st.Symbol = q.Symbol
  st.Latestq= q.AskPrice
  st.Previousq= 0.0
}

func (st *StockHistory) Update(q stream.Quote) {
  st.Latestq= q.AskPrice
  st.Previousq= st.Latestq
}

func (st *StockHistory) GetSymbol() string{
  return st.Symbol
}

func (st *StockHistory) GetTrend() Trend {
  diff := st.Latestq - st.Previousq
  switch { 
  case diff > 0 :  return Up
  case diff < 0 :  return  Down
  default: return Same
  }
}

func (s Stock) HandleQuote(q stream.Quote) {
}


func NewAnalysis(stockList []string) Metric {
  a := Analysis{}
  return a
} 

func (s Sector) Aggregate(in <-chan Trend) (out <-chan Trend){
  trends := make(map[string]Trend)
  for {
    m := <-in
    stock := m.GetSymbol()
    trends[stock] = m.GetTrend()
    count :=0
    ups:=0
    downs:=0
    sames:=0
    for _, trend := range trends {
      count++
      switch trend {
        case Up: ups++
        case Down: downs++
        case Same: sames++
      }
    }
    // var upPct float64
    upPct := float64(ups/count)
    switch {
      case (upPct > 0.75): out <- Hold
      case (upPct < 0.75): out <- Sell
    }
  }
}
