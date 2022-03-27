package stream

import (
	"context"
	"net/url"
	"os"
	"time"
)

// StockOption is a configuration option for the StockClient
type StockOption interface {
	applyStock(*stockOptions)
}


// Option is a configuration option that can be used for both StockClient and CryptoClient
type Option interface {
	StockOption
}

type options struct {
	baseURL        string
	key            string
	secret         string
	reconnectLimit int
	reconnectDelay time.Duration
	processorCount int
	bufferSize     int
	sub            subscriptions

	// for testing only
	connCreator func(ctx context.Context, u url.URL) (conn, error)
}

type funcOption struct {
	f func(*options)
}


func (fo *funcOption) applyStock(o *stockOptions) {
	fo.f(&o.options)
}


func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}


// WithBaseURL configures the base URL
func WithBaseURL(url string) Option {
	return newFuncOption(func(o *options) {
		o.baseURL = url
	})
}

// WithCredentials configures the key and secret to use
func WithCredentials(key, secret string) Option {
	return newFuncOption(func(o *options) {
		if key != "" {
			o.key = key
		}
		if secret != "" {
			o.secret = secret
		}
	})
}

// WithReconnectSettings configures how many consecutive connection
// errors should be accepted and the delay (that is multipled by the number of consecutive errors)
// between retries. limit = 0 means the client will try restarting indefinitely unless it runs into
// an irrecoverable error (such as invalid credentials).
func WithReconnectSettings(limit int, delay time.Duration) Option {
	return newFuncOption(func(o *options) {
		o.reconnectLimit = limit
		o.reconnectDelay = delay
	})
}

// WithProcessors configures how many goroutines should process incoming
// messages. Increasing this past 1 means that the order of processing is not
// necessarily the same as the order of arrival the from server.
func WithProcessors(count int) Option {
	return newFuncOption(func(o *options) {
		o.processorCount = count
	})
}

// WithBufferSize sets the size for the buffer that is used for messages received
// from the server
func WithBufferSize(size int) Option {
	return newFuncOption(func(o *options) {
		o.bufferSize = size
	})
}

func withConnCreator(connCreator func(ctx context.Context, u url.URL) (conn, error)) Option {
	return newFuncOption(func(o *options) {
		o.connCreator = connCreator
	})
}

type stockOptions struct {
	options
	quoteHandler         func(Quote)
}

// defaultStockOptions are the default options for a client.
// Don't change this in a backward incompatible way!
func defaultStockOptions() *stockOptions {
	baseURL := "https://stream.data.alpaca.markets/v2"
	if s := os.Getenv("DATA_PROXY_WS"); s != "" {
		baseURL = s
	}

	return &stockOptions{
		options: options{
			baseURL:        baseURL,
			key:            os.Getenv("APCA_API_KEY_ID"),
			secret:         os.Getenv("APCA_API_SECRET_KEY"),
			reconnectLimit: 20,
			reconnectDelay: 150 * time.Millisecond,
			processorCount: 1,
			bufferSize:     100000,
			sub: subscriptions{
				trades:       []string{},
				quotes:       []string{},
				bars:         []string{},
				updatedBars:  []string{},
				dailyBars:    []string{},
				statuses:     []string{},
				lulds:        []string{},
				cancelErrors: []string{},
				corrections:  []string{},
			},
			connCreator: func(ctx context.Context, u url.URL) (conn, error) {
				return newNhooyrWebsocketConn(ctx, u)
			},
		},
		quoteHandler:         func(q Quote) {},
	}
}

func (o *stockOptions) applyStock(opts ...StockOption) {
	for _, opt := range opts {
		opt.applyStock(o)
	}
}

type funcStockOption struct {
	f func(*stockOptions)
}

func (fo *funcStockOption) applyStock(o *stockOptions) {
	fo.f(o)
}

func newFuncStockOption(f func(*stockOptions)) StockOption {
	return &funcStockOption{
		f: f,
	}
}


// WithQuotes configures initial quote symbols to subscribe to and the handler
func WithQuotes(handler func(Quote), symbols ...string) StockOption {
	return newFuncStockOption(func(o *stockOptions) {
		o.sub.quotes = symbols
		o.quoteHandler = handler
	})
}

