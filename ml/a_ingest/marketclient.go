package a_ingest

import (
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)
// Global Alpaca market data client
var marketClient *marketdata.Client

// Initialize the Alpaca client
func init() {
	marketClient = marketdata.NewClient(marketdata.ClientOpts{
	})
}
