package market

type Symbol string

type Trade struct {
  symbol Symbol 
  long bool
  short bool 
  amout int
}
  

// type Decision interface {
//   Estimate()
// }
type Decision int64
const (
  Hold Decision = iota
  Buy
  Sell
  )

func NewTrade(s Symbol) Trade {
  return Trade{symbol: s}
}
