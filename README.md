# Routing API Tester

Helper repo for testing routing API calls

## Setup
1. Place certs in `certs/` directory if needed.
2. Build CLI: `go build -o api-tester ./cmd/api-tester`.
3. Run: `./api-tester --amountIn=1000000000000000000 --tokenIn=0x... --tokenOut=0x... --slippage=0.5 --maxHops=3 --maxPaths=5 --endpoing=your.api.ip.endpoint --withCerts=false`.

## Usage
CLI calls GetSimpleQuote with flags. Outputs JSON response.

Requires Go 1.23.0+.