package algo

type Trade struct {
  Symbol
  TradeType
  Position
  BenchMark chan Status
}

type Position int 
const (
  Loss Position = iota
  Gain 
  NoChange
)

type TradeType int
const (
  Long TradeType = iota
  Short
)

func NewTrade(s Symbol, t TradeType) Trade {
  return Trade{Symbol: s, TradeType: t}
}

func (t Trade) BenchMark(s Status) bool {
  // if trade is Long and sector Rally and Position is Gain - Hold Position
  // else Sell Position
}

func (t Trade) Enter() {
  go func() {
    switch s {
      case Rally: alpaca.Buy(t.Symbol)
      default: ??
    }
    for {
      ss :=<-s
    }
  }
}


func (p Position) Sell() {


