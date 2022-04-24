package algo
import (
  "github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
  )
type Trade struct {
  Symbol
  TradeType
  Dispatcher
  RestClient alpaca.Client
  StreamClient stream.StocksClient
  strendchan chan StockTrend
  benchmarkchan chan Status
}


type TradeType int
const (
  Long TradeType = iota
  Short
)

func NewTrade(s Symbol, t TradeType, d Dispatcher ) Trade {
  return Trade{Symbol: s, TradeType: t, Dispatcher: d}
}

func (t Trade) AddClient(rc alpaca.Client, sc stream.StocksClient) {
  t.RestClient = rc
  t.StreamClient = sc
}

func (t Trade) Stream()   {
  pricetrend := make(chan StockTrend)
  go func(){
    q := t.GetQuote(t.StreamClient, t.Symbol)
    q.Compute(pricetrend)
  }()
  t.strendchan = pricetrend 
}

func (t Trade) Init(sector []string) bool {
  t.Stream()
  t.Benchmark(sector)
  ok, _ := t.Evaluate()
  return ok
}

func (t Trade) Benchmark(sector []Symbol) {
  sectorTrend := make(SectorTrend)
  for _, s := range sector {
    q := t.GetQuote(t.StreamClient, s)
    q.Compute(sectorTrend)
  }
  // calculate sector status based on stock trends fanning into channel 
  // return a channel of 
  sectorStatus := sectorTrend.Compute()
  t.benchmarkchan = sectorStatus 
}

func (t Trade) Evaluate() (bool, error) {
  // if trade is Long and sector Rally and Position is Gain - Hold Position
  // else Sell Position
  sectorStat :=<-t.benchmarkchan
  stock :=<-t.strendchan
  mktFavorable := false
  switch {
  case sectorStat == Rally && stock.Trend == Up && t.TradeType == Long: mktFavorable = true
  case sectorStat == SellOff && stock.Trend == Down && t.TradeType == Short: mktFavorable= true
  case sectorStat == Random && t.TradeType == Short: mktFavorable= true
  case sectorStat == Random && t.TradeType == Long: mktFavorable= true
  case sectorStat == Unknown: mktFavorable=false
  default: mktFavorable=false
  }
  return mktFavorable, nil
}


func (t Trade) Enter(ls []string) {
  ok := t.Init(ls)
  if !ok {
    panic("bad")
  }
  
  
  t.RestClient.PlaceOrder(req alpaca.PlaceOrderRequest)
  go func() {
    for {
    pos, _ := t.RestClient.GetPosition(t.Symbol)
    pnl := pos.UnrealizedPLPC
    if pnl > decimal(0.05) {
      t.RestClient.ClosePosition(t.Symbol)
    }
      ok, _ := t.Evaluate()
      if ok {
        continue
      } else {
        t.RestClient.Sell(t.Symbol)
      }
      
    }
  }()
}



