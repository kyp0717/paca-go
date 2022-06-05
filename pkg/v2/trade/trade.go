package trade

import (
	"fmt"
	bm "paca-go/pkg/v2/benchmark"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/shopspring/decimal"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

// Trade type is used to track target asset, init and submit order
type Trade struct {
  bm.Symbol
  TradeType
  RestClient alpaca.Client
  // benchmark bm.BenchMark
  sectorTrend bm.TrendChan
  stockTrend bm.PriceChangeChan
}

type TradeType int
const (
  Long TradeType = iota
  Short
)

func New(s bm.Symbol, rc alpaca.Client ) Trade {
  return Trade{Symbol: s, 
    RestClient : rc,
  }
}

// Start the benchmark stream so we can perform evaulation
// The trading stock is also part of the stream
func (t Trade) BenchMark(sc stream.StocksClient) {
  benchmark := bm.New()
  benchmark.AddTradingStock(t.Symbol)
  // benchmark.Stream(sc)
  t.sectorTrend = benchmark.StreamTrend(sc)
  t.stockTrend = benchmark.StreamBase(sc)
}

// func (t Trade) init(bm benchmark) bool {
//   t.Stream()
//   t.Benchmark(sector)
//   ok, _ := t.Evaluate()
//   return ok
// }

func (t Trade) evalTrade()  (bool, error){
  // if trade is Long and sector Rally and Position is Gain - Hold Position
  // else Sell Position
  sector :=<-t.sectorTrend
  stock :=<-t.stockTrend
  mktFavorable := false
  switch {
  case sector == bm.Rally && stock.PriceMove == bm.Up && t.TradeType == Long: mktFavorable = true
  case sector == bm.SellOff && stock.PriceMove== bm.Down && t.TradeType == Short: mktFavorable= true
  case sector == bm.Random && t.TradeType == Short: mktFavorable= true
  case sector == bm.Random && t.TradeType == Long: mktFavorable= true
  case sector == bm.Unknown: mktFavorable=false
  default: mktFavorable=false
  }
  if !mktFavorable {
      t.RestClient.ClosePosition(t.Symbol)
  } 
  return mktFavorable, nil
}

func (t Trade) Enter() (chan int ) {
  done := make(chan int)
  t.submitOrder(100, t.Symbol, t.TradeType)
  for {
    okpos, errpos := t.evalPosition()
    oktrade, errtrade := t.evalTrade()
    if !okpos || !oktrade {
      done<-1 
    } 
   }
  return done
}


func (t Trade) init () error  {
  ok, err := t.evalTrade()
  if err == nil {
    if ok {
      t.submitOrder(100, t.Symbol, t.TradeType)
    }
  } else {
    return err
  }
  return nil
}

func (t Trade) evalPosition() (bool, error) {
    pos, _ := t.RestClient.GetPosition(t.Symbol)
    pnl := pos.UnrealizedPLPC
    threshold := decimal.NewFromFloat(0.05)
    // t := *&threshold
    if pnl.GreaterThan(threshold) {
      t.RestClient.ClosePosition(t.Symbol)
      return false, nil
    }
    return true, nil
}

// Submit an order if quantity is above 0.
func (t Trade) submitOrder(qty int, symbol string, tradeside TradeType) error {
	account, err :=t.RestClient.GetAccount()
	if err != nil {
		return fmt.Errorf("get account: %w", err)
  }
  var side string
  switch tradeside {
  case Long: side = "Long"
  case Short: side = "Short"
  }
	if qty > 0 {
		adjSide := alpaca.Side(side)
		decimalQty := decimal.NewFromInt(int64(qty))
		_, err := t.RestClient.PlaceOrder(alpaca.PlaceOrderRequest{
			AccountID:   account.ID,
			AssetKey:    &symbol,
			Qty:         &decimalQty,
			Side:        adjSide,
			Type:        "market",
			TimeInForce: "day",
		})
		if err == nil {
			fmt.Printf("Market order of | %d %s %s | completed\n", qty, symbol, side)
		} else {
			fmt.Printf("Order of | %d %s %s | did not go through: %s\n", qty, symbol, side, err)
		}
		return err
	}
	fmt.Printf("Quantity is <= 0, order of | %d %s %s | not sent\n", qty, symbol, side)
	return nil
}
