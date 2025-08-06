package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PriceRequest struct {
	Token common.Address
}

type PriceResponse struct {
	PriceInBase   *big.Float
	PriceInStable *big.Float
	Liquidity     float64

	Error string
}

type QuoteRequest struct {
	AmountIn *big.Int
	TokenIn  common.Address
	TokenOut common.Address
	MaxHops  int
	MaxPaths int
}

type QuoteResponse struct {
	AmountIn          *big.Int
	AmountOut         *big.Int
	TokenIn           common.Address
	TokenOut          common.Address
	Liquidity         float64
	PriceData         PriceData
	QuoteResponseInfo []RouteResponseData

	Error string
}

type PriceData struct {
	TokenInBasePrice    *big.Float
	TokenInStablePrice  *big.Float
	TokenOutBasePrice   *big.Float
	TokenOutStablePrice *big.Float

	AmountInValueBase    *big.Float
	AmountInValueStable  *big.Float
	AmountOutValueBase   *big.Float
	AmountOutValueStable *big.Float

	PriceImpact float64
}

type RouteResponseData struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Path      []common.Address
	PoolsInfo []PoolInfo
}

type PoolInfo struct {
	Kind    string
	Address common.Address

	Fee         uint64
	TickSpacing int
	Hook        common.Address
	V4ID        common.Hash
}
