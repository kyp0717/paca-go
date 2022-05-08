package quote

import (
	"log"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)
type Symbol=string

// quotes are fan out to channels
type Dispatcher struct{
  // fan out to stock channel
  Channels map[Symbol]chan stream.Quote // fan out 
  // handler which will fan out
  Handler func(q stream.Quote)
}

type PriceChange int64
const (
  Up PriceChange = iota
  Down
  None
  )

func Dispatch() Dispatcher {

  qc:= make(map[Symbol](chan stream.Quote))
  handler := func(q stream.Quote) {
		qc[q.Symbol] <- q
    // fmt.Println(q)
	}
  return Dispatcher{
    Channels: qc,
    Handler: handler,
  }
}

type QuoteChan chan stream.Quote
func (d *Dispatcher) GetQuote(client stream.StocksClient,s Symbol) QuoteChan{
  d.Channels[s] = make(chan stream.Quote)
  if err := client.SubscribeToQuotes(d.Handler, s); err != nil {
      log.Fatalf("error during subscribing: %s", err)
  }
  return d.Channels[s]
}
// pass work now to data structure in another file
// there is no compute here since we are not transforming data
// pure ETL process only




func (t PriceChange) String() string {
    // declare an array of strings
    // ... operator counts how many
    // items in the array (7)
    names := [...]string{
        "Up", 
        "Down", 
        "None"}
    return names[t]
}

// type StockQuote struct {
//   quote.Symbol
//   PriceChange
//   
// }
// func (t StockQuote) GetSymbol() string {
//     return t.Symbol
// }
// type StockTrend chan<- Feature
type PriceHistory struct {
	quote.Symbol   
	Latest   float64
	Previous float64
}

// type SectorTrend chan StockTrend

// 
func (st StockQuote) Compute() (SectorTrendChan)  {
  sink:= make(SectorTrendChan)
  go func() {
  trends := make(map[string]PriceChange)
  for {
    strend := <-st
    stock := strend.GetSymbol()
    trends[stock] = strend.Trend
    count :=0
    ups:=0
    downs:=0
    none:=0
    for _, trend := range trends {
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
