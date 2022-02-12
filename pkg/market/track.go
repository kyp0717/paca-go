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

type Track struct{
  paca stream.StocksClient
  stock string
  TapChan chan stream.Quote // source data - one per quote
  SinkChan chan<- Last2Quotes // send to channel - one total 
  // Last Last2Quotes
}

type Last2Quotes struct {
	// symbol    string
	Latestq   float64
	Previousq float64
}

func (t Track) createTap() {
  cc := make(chan stream.Quote)
  t.TapChan = cc
}



func (t Track) createSink() (chan Last2Quotes){
  cc := make(chan Last2Quotes)
	return cc
}

func (t Track) generate(c stream.StocksClient) {
    c.SubscribeToQuotes(t.track, t.stock)
		go t.track()
}


// extract stock quote from channel and insert into data type
func (t Track) track(q stream.Quote) {
  // t.stock = s
	t.TapChan
  l := &Last2Quotes{q.AskPrice, 0.0}
  // t.Last = l
	for {
		q := <-t.TapChan
		fmt.Println(q)
		l.Previousq = l.Latestq
		l.Latestq = q.AskPrice
		fmt.Printf("%v\n", t)
		time.Sleep(3 * time.Second)
    t.SinkChan <- *l
	}
}


