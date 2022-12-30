package indodax

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// GetInfo returns user balances, user wallet, user id, username, profile picture and server's timestamp.
func (cl *Client) GetInfo(ctx context.Context) (usrInfo *UserInfo, err error) {
	respBody, err := cl.curlPrivate(ctx, apiViewGetInfo, nil)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respGetInfo := &respGetInfo{}

	err = json.Unmarshal(respBody, respGetInfo)
	if err != nil {
		err = fmt.Errorf("GetInfo: " + err.Error())
		return nil, err
	}
	if respGetInfo.Success != 1 {
		return nil, fmt.Errorf("GetInfo: " + respGetInfo.Message)
	}

	cl.Info = respGetInfo.Return

	printDebug(cl.Info)

	return cl.Info, nil
}

// TransHistory returns list of deposits and withdrawals of all currencies
func (cl *Client) TransHistory(ctx context.Context) (transHistory *TransHistory, err error) {
	respBody, err := cl.curlPrivate(ctx, apiViewTransactionHistory, nil)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	history := &respTransHistory{}

	err = json.Unmarshal(respBody, history)
	if err != nil {
		err = fmt.Errorf("TransHistory: " + err.Error())
		return nil, err
	}
	if history.Success != 1 {
		return nil, fmt.Errorf("TransHistory: " + history.Message)
	}

	printDebug(history)

	return history.Return, nil
}

// OpenOrders returns current open orders (buy and sell) by pair.
func (cl *Client) OpenOrders(ctx context.Context, pairName string) (openOrders []OpenOrders, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	respBody, err := cl.curlPrivate(ctx, apiViewOpenOrders, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respOpenOrders := &responseOpenOrders{}

	err = json.Unmarshal(respBody, respOpenOrders)
	if err != nil {
		err = fmt.Errorf("OpenOrders: " + err.Error())
		return nil, err
	}
	if respOpenOrders.Success != 1 {
		return nil, fmt.Errorf("OpenOrders: " + respOpenOrders.Message)
	}

	printDebug(respOpenOrders)

	return respOpenOrders.Return.Orders, nil
}

// AllOpenOrders returns the list of current open orders (buy and sell) all pair.
func (cl *Client) AllOpenOrders(ctx context.Context) (allOpenOrders map[string][]OpenOrders, err error) {
	respBody, err := cl.curlPrivate(ctx, apiViewOpenOrders, nil)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respOpenOrders := &responseAllOpenOrders{}

	err = json.Unmarshal(respBody, respOpenOrders)
	if err != nil {
		err = fmt.Errorf("AllOpenOrders: " + err.Error())
		return nil, err
	}
	if respOpenOrders.Success != 1 {
		return nil, fmt.Errorf("AllOpenOrders: " + respOpenOrders.Message)
	}

	printDebug(respOpenOrders)

	return respOpenOrders.Return.Orders, nil
}

// TradeHistory returns information about transaction in buying and selling history
func (cl *Client) TradeHistory(
	ctx context.Context,
	pairName string,
	count, startTradeID, endTradeID int64,
	sortOrder string,
	sinceTime *time.Time,
	endTime *time.Time,
) (openOrders []TradeHistory, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	if count > 0 {
		params.Set("count", strconv.FormatInt(count, 10))
	}
	if startTradeID > 0 {
		params.Set("from_id", strconv.FormatInt(startTradeID, 10))
	}
	if endTradeID > 0 {
		params.Set("end_id", strconv.FormatInt(endTradeID, 10))
	}

	sortOrder = strings.ToLower(sortOrder)
	switch sortOrder {
	case "asc":
		params.Set("order", "asc")
	case "desc":
		params.Set("order", "desc")
	}

	if sinceTime != nil {
		params.Set("since", strconv.FormatInt(sinceTime.Unix(), 10))
	}
	if endTime != nil {
		params.Set("end", strconv.FormatInt(endTime.Unix(), 10))
	}

	respBody, err := cl.curlPrivate(ctx, apiViewTradeHistory, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	tradeHistory := &respTradeHistory{}

	err = json.Unmarshal(respBody, tradeHistory)
	if err != nil {
		err = fmt.Errorf("TradeHistory: " + err.Error())
		return nil, err
	}
	if tradeHistory.Success != 1 {
		return nil, fmt.Errorf("TradeHistory: " + tradeHistory.Message)
	}

	printDebug(tradeHistory)

	return tradeHistory.Return.Trades, nil
}

// OrderHistory returns the list of order history (buy and sell).
func (cl *Client) OrderHistory(
	ctx context.Context,
	pairName string,
	count, from int64,
) (openOrders []OrderHistory, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	if count > 0 {
		params.Set("count", strconv.FormatInt(count, 10))
	}
	if from > 0 {
		params.Set("from", strconv.FormatInt(from, 10))
	}

	respBody, err := cl.curlPrivate(ctx, apiViewOrderHistory, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	orderHistory := &respOrderHistory{}

	err = json.Unmarshal(respBody, orderHistory)
	if err != nil {
		err = fmt.Errorf("OrderHistory: " + err.Error())
		return nil, err
	}
	if orderHistory.Success != 1 {
		return nil, fmt.Errorf("OrderHistory: " + orderHistory.Message)
	}

	printDebug(orderHistory)

	return orderHistory.Return.Orders, nil
}

// GetOrder returns specific order details by pairName and orderId
func (cl *Client) GetOrder(
	ctx context.Context,
	pairName string,
	orderId int64,
) (getOrder *GetOrder, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	if orderId > 0 {
		params.Set("order_id", strconv.FormatInt(orderId, 10))
	}

	respBody, err := cl.curlPrivate(ctx, apiViewGetOrder, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	getOrders := &respGetOrders{}

	err = json.Unmarshal(respBody, getOrders)
	if err != nil {
		err = fmt.Errorf("GetOrder: " + err.Error())
		return nil, err
	}
	if getOrders.Success != 1 {
		return nil, fmt.Errorf("GetOrder: " + getOrders.Message)
	}

	printDebug(getOrders)

	return getOrders.Return.Order, nil
}

// CancelOrderBuy cancels an existing open buy order.
func (cl *Client) CancelOrderBuy(
	ctx context.Context,
	pairName string,
	orderId int64,
) (cancelOrder *CancelOrder, err error) {
	cancelOrder, err = cl.cancelOrder(ctx, "buy", pairName, orderId)
	if err != nil {
		return nil, err
	}
	return cancelOrder, nil
}

// CancelOrderSell cancels an existing open sell order.
func (cl *Client) CancelOrderSell(
	ctx context.Context,
	pairName string,
	orderId int64,
) (cancelOrder *CancelOrder, err error) {
	cancelOrder, err = cl.cancelOrder(ctx, "sell", pairName, orderId)
	if err != nil {
		return nil, err
	}
	return cancelOrder, nil
}

// This method is for canceling an existing open order.
func (cl *Client) cancelOrder(
	ctx context.Context,
	method, pairName string,
	orderId int64,
) (cancelOrder *CancelOrder, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}
	if orderId == 0 {
		return nil, ErrInvalidOrderID
	}

	params := url.Values{}
	params.Set("pair", pairName)
	params.Set("type", method)
	params.Set("order_id", strconv.FormatInt(orderId, 10))

	respBody, err := cl.curlPrivate(ctx, apiTradeCancelOrder, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	order := &respCancelOrder{}

	err = json.Unmarshal(respBody, order)
	if err != nil {
		err = fmt.Errorf("CancelOrder Unmarshal: " + err.Error())
		return nil, err
	}
	if order.Success != 1 {
		return nil, fmt.Errorf("CancelOrder: " + order.Error)
	}

	printDebug(order)

	return order.Return, nil
}

// TradeBuy opens a new buy order
func (cl *Client) TradeBuy(
	ctx context.Context,
	pairName string,
	price, amount float64,
) (trade *Trade, err error) {

	keyName := strings.Split(pairName, "_")
	if len(keyName) < 2 {
		return nil, ErrInvalidPairName
	}
	assetName := keyName[1]
	trade, err = cl.trade(
		ctx,
		"buy", pairName, assetName,
		price, amount,
	)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

// TradeSell opens a new sell order
func (cl *Client) TradeSell(
	ctx context.Context,
	pairName string,
	price, amount float64,
) (trade *Trade, err error) {

	keyName := strings.Split(pairName, "_")
	if len(keyName) < 2 {
		return nil, ErrInvalidPairName
	}
	assetName := keyName[0]
	trade, err = cl.trade(
		ctx,
		"sell", pairName, assetName,
		price, amount,
	)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

// trade opens a new order
func (cl *Client) trade(
	ctx context.Context,
	method, pairName, assetName string,
	price, amount float64,
) (trade *Trade, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}
	if assetName == "" {
		return nil, ErrInvalidAssetName
	}
	if price == 0 {
		return nil, ErrInvalidPrice
	}
	if amount == 0 {
		return nil, ErrInvalidAmount
	}

	params := url.Values{}
	params.Set("pair", pairName)
	params.Set("type", method)
	params.Set("price", fmt.Sprintf("%.8f", price))
	params.Set(assetName, fmt.Sprintf("%.8f", amount))

	respBody, err := cl.curlPrivate(ctx, apiTrade, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	t := &respTrade{}

	err = json.Unmarshal(respBody, t)
	if err != nil {
		err = fmt.Errorf("trade: " + err.Error())
		return nil, err
	}
	if t.Success != 1 {
		return nil, fmt.Errorf("Trade: " + t.Error)
	}

	printDebug(t)

	return t.Return, nil
}
