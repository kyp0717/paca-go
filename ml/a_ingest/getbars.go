package a_ingest

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

// no need to build ingest script 
// ingest script only if we need to save to db
// Get percent changes of the stock prices over the past 10 minutes.
func getbars(symbols []string ) ([]marketdata.Bar, error) {

	end := time.Now()
	start := end.Add(-10 * time.Minute)
	feed := ""
	// if !hasSipAccess {
	// 	feed = "iex"
	// }

  // this function will start a client with env var 
	multiBars, err := marketdata.GetMultiBars(symbols, marketdata.GetBarsRequest{
		TimeFrame: marketdata.OneMin,
		Start:     start,
		End:       end,
		Feed:      feed,
	})
	if err != nil {
		return (nil, fmt.Errorf("get multi bars: %w", err))
	}
	// Transform the result to our MinuteBar structure
	var result []marketdata.Bar
	for _, bar := range multiBars {
		result = append(result, marketdata.Bar{
			Timestamp: bar.Timestamp,
			Open:      bar.Open,
			High:      bar.High,
			Low:       bar.Low,
			Close:     bar.Close,
			Volume:    bar.Volume,
		})
	}

	return result, nil

}
