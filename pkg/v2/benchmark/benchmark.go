package benchmark

import (
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
  )


// type Sector int
// const (
//   Tech Sector = iota
//   Healthcare
//   )
// type StockList []string
type Sector []string

type BenchMark struct {
  Dispatcher
  Sector 
  // Trend  // channel define in trends module
  // StreamClient stream.StocksClient
}

func New(dp Dispatcher) BenchMark{
  return BenchMark{
    Dispatcher: dp,
   }
}

func (b *BenchMark) addTechSector() {
  techList := []string{"AMD","TSLA"}
  b.Sector=techList
}

// stream return channel
func (b BenchMark) Stream(sc stream.StocksClient) TrendChan{
  b.addTechSector()

  sectorTrend := make(TrendChan)
  priceChange := make(PriceChangeChan)
  for _, ticker := range b.Sector {
    quotechan := b.GetQuote(sc, ticker)
    // fan into this channel
    quotechan.Compute(priceChange)
    priceChange.Compute(sectorTrend)
  }
  // calculate sector status based on stock trends fanning into channel 
  // return a channel of 
  return sectorTrend
  // b.benchmarkchan = sectorStatus 
}