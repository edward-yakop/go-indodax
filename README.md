# GO-Indodax - A Library trading platform Indodax using Go Language
- [Description](#description)
- [Features](#features)
- [Trade API Documentation](https://indodax.com/downloads/BITCOINCOID-API-DOCUMENTATION.pdf)
- [godoc](https://godoc.org/github.com/edward-yakop/go-indodax)
- [Example](#example)

## Description
Welcome to Indodax golang library. These library outline exchange functionality, market details, and APIs.

APIs are separated into two categories: private and public. Private APIs require authentication and provide access to placing orders and other account information. Public APIs provide market data.

## Features

Public Endpoint
- Ticker
- Depth (Order Book)
- Trades (Trade History)

Private Endpoint
- Get Information User
- Withdraw/Deposit History
- Trading History
- Order History
- Open Orders
- Trade
- Withdraw (Coming Soon)
``
## Example
To start use the library you can get the repository first:

`go get github.com/edward-yakop/go-indodax`

Public Endpoint 
``` go
package test

import (
	"context"
    "fmt"
    "github.com/edward-yakop/go-indodax"
)

func main() {
    cl, err := indodax.NewClient(
		"",
		"",
	)
	ticker, err := cl.GetTicker(context.Background(), "ten_idr")
	if err != nil {
		fmt.Println(err)
	}
    fmt.Printf("Ticker %v\n",ticker)
}
```

Private Endpoint 
```go
package test

import (
    "context"
    "fmt"
    "github.com/edward-yakop/go-indodax"
)

func main() {
    cl, err := indodax.NewClient(
		"key", 
		"secret", 
	)
	tradeBuy, err := cl.TradeBuy(context.Background(), "usdt_idr", 12000, 50000)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Printf("TradeBuy %v\n",tradeBuy)
}
```
## Notes
- For bug report you can refer to [this](https://github.com/edward-yakop/go-indodax/blob/master/.github/ISSUE_TEMPLATE/bug_report.md)
- For feature request you can refer to [this](https://github.com/edward-yakop/go-indodax/blob/master/.github/ISSUE_TEMPLATE/feature_request.md)
