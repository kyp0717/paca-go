package market
import (
	// "context"
	// "fmt"
	// "fmt"
	"log"
	// "os"
	// "errors"
	// "os/signal"
	// "time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

type LatestQuotes struct {
	// symbol   string
	Latest   float64
	Previous float64
}

func (l *LatestQuotes) update(c <-chan stream.Quote) {
  q := <-c
  l.Previous = l.Latest
  l.Latest = q.AskPrice
}

// // the closure is the function that is returned 
// func makeHandler(c chan stream.Quote) (func(), (chan stream.Quote)) {
//   c:= make(chan stream.Quote)
//   return ((func(q stream.Quote) { c<-q }), c)
// }
//
//
// the closure is the function that is returned 
func makeHandler() (func(q stream.Quote), (chan stream.Quote)) {
  c:= make(chan stream.Quote)
  return func(q stream.Quote) { c<-q }, c
}




// a bunch of side effects
func Qstream(client *stream.StocksClient, s string) (chan stream.Quote) {
  // qstream := make(chan stream.Quote)
  // handler := makeHandler(qstream)
  handler, c := makeHandler()
	err := (*client).SubscribeToQuotes(handler, s)
	// err := client.SubscribeToQuotes(handler, s)
	if err != nil {
		log.Fatalf("error during subscribing: %s", err)
	}
  return c
}

func GetPrices(in <-chan stream.Quote,out chan<- LatestQuotes) {
  l := LatestQuotes{0.0, 0.0}
	for {
    l.update(in)
    out<-l
	}
}

//
// func get_trend(in <-chan stream.Quote,trend chan<- float) {
//   l := LatestQuotes{0.0, 0.0}
// 	for {
//     l.update(in)
//     diff = l.Latest - l.Previous
//     pct = diff/l.Latest
//     if pct < -0.001 {
//     }
// 	}
// }
