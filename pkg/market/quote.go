package market

import "github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"

type QuoteHandler interface {
	HandleQuote()
}

// type Quote struct {
// 	stream.Quote
// }

// type Quote string
type Quote stream.Quote



func (q Quote) HandleQuote() {
}

