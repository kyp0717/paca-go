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

type Tracker struct{
  userCancel 
  

}

type Last2Quotes struct {
	symbol    string
	latestq   float64
	previousq float64
}

type QuoteChanMap map[string](chan stream.Quote)
type QuoteChanSink chan<- Last2Quotes

func creatQuoteChanMap(ls []string) QuoteChanMap {
	qMap := make(map[string](chan stream.Quote))
	for _, v := range ls {
		qMap[v] = make(chan stream.Quote)
	}
	return qMap
}

func spawnQuoteRoutines(qchans QuoteChanMap, sinkchan QuoteChanSink) {
	for _, qchan := range qchans {
		go extractQuote(qchan, sinkchan)
	}
}

// extract stock quote from channel and insert into data type
func extractQuote(qchan <-chan stream.Quote, sinkchan QuoteChanSink) {
	q := <-qchan
	t := &Last2Quotes{q.Symbol, q.AskPrice, 0.0}
	for {
		q := <-qchan
		fmt.Println(q)
		t.previousq = t.latestq
		t.latestq = q.AskPrice
		fmt.Printf("%v\n", t)
		time.Sleep(3 * time.Second)
		sinkchan <- *t
	}
}


