package benchmark

import (
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
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

func (pc PriceChangeChan ) Compute() SectorTrendChan  {
  sink:= make(SectorTrendChan)
  go func() {
  pricechanges := make(map[string]PriceMove)
  for {
    pricechange := <-pc
    pricechanges[pricechange.Symbol] = pricechange.PriceMove
    count :=0
    ups:=0
    downs:=0
    none:=0
    for _, trend := range pricechanges{
      count++
      switch trend {
        case Up: ups++
        case Down: downs++
        case None: none++
      }
    }
    // var upPct float64
    upPct := float64(ups/count)
    switch {
      case (upPct > 0.75): sink<-Rally
      case (upPct < 0.25): sink<-SellOff
    }
   }
  }()
  return sink
}



