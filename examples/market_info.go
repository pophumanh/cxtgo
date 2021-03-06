package main

import (
	"context"
	"fmt"
	"os"

	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/exchanges/binance"
)

func main() {
	binance := binance.New(
		cxtgo.WithAPIKey("test"),
		cxtgo.WithAPISecret("test"),
		cxtgo.WithDebug(true),
		cxtgo.WithDebuglogger(os.Stdout),
	)

	market, err := binance.Markets(context.Background())

	info := market[cxtgo.NewSymbol("ETH", "BTC")]
	fmt.Printf("%v\n", info.Precision.Amount)

	fmt.Println(err)
}
