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
  base Symbol
  dispatcher Dispatcher
  TechTrend []string
  Sector 
  // Trend  // channel define in trends module
}

func New() BenchMark{
  dp := Dispatch()
  return BenchMark{
    dispatcher: dp,
   }
}

func (b *BenchMark) AddTradingStock(s Symbol) {
  b.Sector=append(b.Sector, s)
}


// temporary make benchmark tech for dev
func (b *BenchMark) addTechSector() {
  techList := []string{"AMD","TSLA"}
  b.Sector=techList
}

// stream return channel
func (b BenchMark) StreamTrend(sc stream.StocksClient) TrendChan{
  b.addTechSector()

  sectorTrend := make(TrendChan)
  priceChange := make(PriceChangeChan)
  for _, ticker := range b.Sector {
    if (ticker == b.base) {
      continue
    }
    quotechan := b.dispatcher.GetQuote(sc, ticker)
    // fan into this channel
    quotechan.Compute(priceChange)
    priceChange.Compute(sectorTrend)
  }
  // calculate sector status based on stock trends fanning into channel 
  // return a channel of 
  return sectorTrend
  // b.benchmarkchan = sectorStatus 
}

func (b BenchMark) StreamBase(sc stream.StocksClient) PriceChangeChan{
  priceChange := make(PriceChangeChan)
  quotechan := b.dispatcher.GetQuote(sc, b.base)
    // fan into this channel
  quotechan.Compute(priceChange)
  return priceChange 
  // b.benchmarkchan = sectorStatus 
}
