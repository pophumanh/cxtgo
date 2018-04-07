package cxtgo

import (
	"sync"
)

// Offer defines an offer from the order book
type Offer struct {
	Price    float64
	Quantity float64
}

// Bid represents a bid offer from the order book
type Bid Offer

// Ask represents an ask offer from the order book
type Ask Offer

// Orderbook is a definition for an orderbook
type Orderbook interface {
	Symbol() Symbol
	Head(n int) Summary
	Spread() float64
	Depth() Summary
}

// Summary defines a view of the order book.
// The bids are sorted descending and the ask ascending.
type Summary struct {
	Bids []Bid
	Asks []Ask
}

type ConcurrentOrderbook struct {
	sync.RWMutex

	symbol Symbol
}

func (co *ConcurrentOrderbook) Symbol() Symbol {
	return co.symbol
}
func (co *ConcurrentOrderbook) Head(n int) Summary {
	return Summary{}
}
func (co *ConcurrentOrderbook) Spread() float64 {
	return -1.0
}
func (co *ConcurrentOrderbook) Depth() Summary {
	return Summary{}
}
