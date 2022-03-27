package algo

import (
	"paca-go/pkg/market"
)

type stockTrend struct {
	symbol    string
	latestq   float64
	previousq float64
}

// // algebraic type 
// const (
//   Up MarketTrend = iota
//   Down
//   )

func (st StockTrend) Calculate (symbol string) {
	q := <-qs.QuoteChannels[symbol]
	t := &Last{q.Symbol, q.AskPrice, 0.0}
	for {
	  q := <-qs.QuoteChannels[symbol]
		t.previousq = t.latestq
		t.latestq = q.AskPrice
    qs.Sink<-t
		// fmt.Printf("%v\n", t)
		// time.Sleep(1 * time.Second)
	}
}
func (s stockTrend) processQuote() {

}

func run_model() {
  // parameters := 

  // metrics := 
	trend := market.NewPacaStream()
  trend.GetQuote("AAPL")
  trend.GetQuote("TSLA")
  // initiate the fan out -- add handler to trend
  trend.FanOut()

  for {
    mktstatus := trend.Process()
    if mktstatus == 

  }
  

  
  
}
