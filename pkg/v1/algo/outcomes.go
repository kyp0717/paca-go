package algo

import (
	// "github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

type Decision int64
const (
  Sell Decision= iota
  Buy
  Hold
  )


type Status int64
const (
  Rally Status= iota
  SellOff
  Random
  Unknown
  )

type SectorStatus chan Status
type MarketStatus chan Status

func (ss SectorStatus) Trade() {
  go func() {
    for {
    status<-ss
    switch status {
      case Rally: 
      case SellOff:
      }
    }
  }()

}
