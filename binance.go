package cxtgo

import (
	"context"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo/base"
)

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	base   *base.Exchange
	client *binance.Client
}

// NewBinance returns an instance of the binance exchange
func NewBinance(opts ...base.ExchangeOpt) *Binance {
	binanceOpts := []base.ExchangeOpt{
		base.WithName("Binance"),
		base.WithUserAgent("cxt/0.1"),
		base.WithVersion("v3"),
	}
	binanceOpts = append(binanceOpts, opts...)

	ex := base.NewExchange(binanceOpts...)
	b := &Binance{
		base:   ex,
		client: binance.NewClient(ex.APIKEY, ex.APISecret),
	}
	return b
}

func (b *Binance) Info() base.Exchange {
	return *b.base
}

func (b *Binance) LoadMarkets(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchMarkets(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchTicker(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchTickers(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOrderBook(ctx context.Context) (OrderBookResponse, error) {
	panic("not implemented")
}

func (b *Binance) FetchOHLCV(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchTrades(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchBalance(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) CreateOrder(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) CancelOrder(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOrder(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOrders(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOpenOrders(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchClosedOrders(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchMyTrades(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) Deposit(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) Withdraw(ctx context.Context) (Response, error) {
	panic("not implemented")
}
