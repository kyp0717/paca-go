package market 
	
import (
"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
 )

type Metric interface {
  TransformQuote()
  GetTrend() Direction
  GetSymbol() string
}

// type Decision interface {
//   Estimate()
// }
type Decision int64
const (
  Hold Decision = iota
  Buy
  Sell
  )

type Symbol string
type Direction int64
const (
  Up Direction= iota
  Down
  Same
  )
// algo defines 2 features :
// 1. QuoteHandler The function that transforms the stream quotes to metric
// ...
type Algo struct {
  // procedure on how to transform quotes
  HandleQuote func(in <-chan stream.Quote, out chan<- Metric)
  // how to make decision
  // Aggregate func(in <-chan Metric) Decision
  // Feature Metric
  // Feature Map(string)[Decision]
  Features map[string]Decision
  Quotes []string
  TradeType string // long vs short
  TradeChan chan<- Decision
  Trade Symbol
  
}

func NewAlgo(ls []string, s Symbol) Algo {
  t := Algo{Quotes: ls,
            Trade: s,
            TradeChan: make(chan Decision)}
  return t
}

func (a Algo) InitPosition() {
}

func (a Algo) Sell() {
}


func (a Algo) Aggregate(in <-chan Metric) {
  trends := make(map[string]Direction)
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
      case (upPct > 0.75): 
         a.TradeChan <- Hold
      case (upPct < 0.75): 
         a.TradeChan <- Sell
    }
  }
}


