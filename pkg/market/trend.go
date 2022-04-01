package market 

import (
	// "context"
	// "time"
	// "fmt"
	// "log"
	// "os"
	// "os/signal"
	// "testing"
	//
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	// "github.com/joho/godotenv"
)

type Metric interface {
  GetTrend() 
}

type StockTrend struct {
	symbol    string
	latestq   float64
	previousq float64
  trend  string
}

// type StockTrend struct {
//   PriceHistory
// }

func (st *StockTrend) Init(q stream.Quote) {
  st.symbol = q.Symbol
  st.latestq= q.AskPrice
  st.previousq= 0.0
}

func (st *StockTrend) Update(q stream.Quote) {
  st.latestq= q.AskPrice
  st.previousq= st.latestq
}

func (st *StockTrend) GetTrend() {
  diff := st.latestq - st.previousq
  if diff > 0 {
    st.trend="Up"
  } else {
    st.trend="Down"
  }
}


