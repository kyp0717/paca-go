package b_transform

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	"github.com/shopspring/decimal"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

// no need to build ingest script 
// ingest script only if we need to save to db
// Get percent changes of the stock prices over the past 10 minutes.
func getbars(symbols: []string ) ([]marketdata.Bar, error) {

	end := time.Now()
	start := end.Add(-10 * time.Minute)
	feed := ""
	if !hasSipAccess {
		feed = "iex"
	}
	multiBars, err := algo.dataClient.GetMultiBars(symbols, marketdata.GetBarsRequest{
		TimeFrame: marketdata.OneMin,
		Start:     start,
		End:       end,
		Feed:      feed,
	})
	if err != nil {
		return fmt.Errorf("get multi bars: %w", err)
	}


}
