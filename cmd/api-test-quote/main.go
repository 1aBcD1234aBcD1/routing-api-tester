package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/1aBcD1234aBcD1/routing-api-tester/internal/client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

func main() {
	amountInStr := flag.String("amountIn", "", "Amount in (wei)")
	tokenInStr := flag.String("tokenIn", "", "Token in address")
	tokenOutStr := flag.String("tokenOut", "", "Token out address")
	maxHops := flag.Int("maxHops", 3, "Max hops")
	maxPaths := flag.Int("maxPaths", 5, "Max paths")
	endpoint := flag.String("endpoint", "localhost", "API Endpoint")
	withCerts := flag.Bool("withCerts", true, "Certs usage")
	flag.Parse()

	amountIn := new(big.Int)
	amountIn.SetString(*amountInStr, 10)

	req := client.QuoteRequest{
		AmountIn: amountIn,
		TokenIn:  common.HexToAddress(*tokenInStr),
		TokenOut: common.HexToAddress(*tokenOutStr),
		MaxHops:  *maxHops,
		MaxPaths: *maxPaths,
	}

	c := client.NewAPIClient(*endpoint, *withCerts)
	resp, err := c.GetSimpleQuote(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	jsonOut, _ := json.MarshalIndent(resp, "", "  ")

	fmt.Println(string(jsonOut))
}
