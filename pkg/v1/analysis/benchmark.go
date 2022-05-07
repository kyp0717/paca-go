package algo

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
  SectorTrend  // channel define in trends module
  StreamClient stream.StocksClient
}

func NewBenchMark(dp Dispatcher, sc stream.StocksClient) BenchMark{
  return BenchMark{
    Dispatcher: dp,
    StreamClient: sc,
   }
}

func (b *BenchMark) addTechSector() {
  techList := []string{"AMD","TSLA"}
  b.Sector=techList
}

// stream return channel
func (b BenchMark) Stream() SectorStatus{
  b.addTechSector()

  b.SectorTrend = make(SectorTrend)
  for _, stock := range b.Sector {
    quotechan := b.GetQuote(b.StreamClient, stock)
    // fan into this channel
    quotechan.Compute(b.SectorTrend)
  }
  // calculate sector status based on stock trends fanning into channel 
  // return a channel of 
  return b.SectorTrend.Compute()
  // b.benchmarkchan = sectorStatus 
}
