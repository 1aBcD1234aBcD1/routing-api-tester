package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/1aBcD1234aBcD1/routing-api-tester/internal/client"
	"github.com/ethereum/go-ethereum/common"
	"os"
)

func main() {

	tokenStr := flag.String("token", "", "Token address")
	endpoint := flag.String("endpoint", "localhost", "API Endpoint")
	withCerts := flag.Bool("withCerts", true, "Certs usage")
	flag.Parse()

	req := client.PriceRequest{
		Token: common.HexToAddress(*tokenStr),
	}

	c := client.NewAPIClient(*endpoint, *withCerts)
	resp, err := c.GetTokenPrice(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	jsonOut, _ := json.MarshalIndent(resp, "", "  ")

	fmt.Println(string(jsonOut))
}
