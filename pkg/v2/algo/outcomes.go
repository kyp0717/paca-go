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



// func (ss SectorTrendChan) Trade(t Trade) {
//   go func() {
//
//     for {
//       status:=<-ss
//     switch {
//       case status == Rally && t.TradeType==Long: 
//       case status == SellOff && t.TradeType==Long: 
//       }
//     }
//   }()
//
// }
