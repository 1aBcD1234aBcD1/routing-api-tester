package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type RequestConfig struct {
	AmountIn *big.Int
	TokenIn  common.Address
	TokenOut common.Address
	Slippage *big.Float
	MaxHops  int
	MaxPaths int
}

type QuoteResponse struct {
	AmountIn          *big.Int
	AmountOut         *big.Int
	AmountOutWithSlip *big.Int
	TokenIn           common.Address
	TokenOut          common.Address
	QuoteResponseInfo []RouteResponseData
	CallData          []byte
}

type RouteResponseData struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Path      []common.Address
	PoolsInfo []PoolInfo
}

type PoolInfo struct {
	Address     common.Address
	Kind        string
	PoolFee     *big.Float
	Liquidity   *big.Int
	PriceImpact *big.Float
}
