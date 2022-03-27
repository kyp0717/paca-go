package stream

import "time"

// Quote is a stock quote from the market
type Quote struct {
	Symbol      string
	BidExchange string
	BidPrice    float64
	BidSize     uint32
	AskExchange string
	AskPrice    float64
	AskSize     uint32
	Timestamp   time.Time
	Conditions  []string
	Tape        string

	internal quoteInternal
}

type quoteInternal struct {
	ReceivedAt time.Time
}

// Internal contains internal fields. There aren't any behavioural or backward compatibility
// promises for them: they can be empty or removed in the future. You should not use them at all.
func (q Quote) Internal() quoteInternal {
	return q.internal
}

