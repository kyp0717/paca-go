package market

import (
	// "context"
	"fmt"
	// "log"
	// "os"

	// "errors"
	// "os/signal"
	// "sync/atomic"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	// "github.com/joho/godotenv"
)

type ISector interface {
  track() 
}


type Sector struct{
  quotes []string
}

type Track struct{
  stock string
  tap chan stream.Quote // source data
}

type MarketStatus int
const (
  Random MarketStatus = iota
  Rally 
  SellOff
  Unpredictable
  )

func (s Sector) track() MarketStatus {
  for i := range s.quotes {
    trends

  }


  return Unpredictable
} 

