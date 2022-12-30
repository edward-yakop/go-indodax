package indodax

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetTicker contains the price summary like volume, last price, open buy, open sell of an individual pair.
func (cl *Client) GetTicker(ctx context.Context, pairName string) (ticker *Ticker, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	urlPath := fmt.Sprintf(pathTicker, pairName)

	body, err := cl.curlPublic(ctx, urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetTicker: " + err.Error())
	}

	printDebug(string(body))

	tickerRes := tickerResponse{}
	err = json.Unmarshal(body, &tickerRes)
	if err != nil {
		return nil, fmt.Errorf("GetTicker: " + err.Error())
	}

	ticker, err = tickerRes.toTicker(pairName)
	if err != nil {
		return nil, fmt.Errorf("GetTicker: " + err.Error())
	}

	return ticker, nil
}

// GetOrderBook contains the order book buy and sell of an individual pair.
func (cl *Client) GetOrderBook(ctx context.Context, pairName string) (orderBook *OrderBook, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	urlPath := fmt.Sprintf(pathDepth, pairName)

	body, err := cl.curlPublic(ctx, urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetOrderBook: " + err.Error())
	}

	printDebug(string(body))

	orderBook = &OrderBook{}
	err = json.Unmarshal(body, &orderBook)
	if err != nil {
		return nil, fmt.Errorf("GetOrderBook: " + err.Error())
	}

	return orderBook, nil
}

// GetListTrades contains the historical trade of an individual pair.
func (cl *Client) GetListTrades(ctx context.Context, pairName string) (
	listTrade []*ListTrade, err error,
) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	urlPath := fmt.Sprintf(pathTrades, pairName)

	body, err := cl.curlPublic(ctx, urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetListTrades: " + err.Error())
	}

	printDebug(string(body))

	err = json.Unmarshal(body, &listTrade)
	if err != nil {
		return nil, fmt.Errorf("GetListTrades: " + err.Error())
	}

	return listTrade, nil
}

// GetSummaries contains the price summary like volume, last price, open buy, open sell of all pair.
func (cl *Client) GetSummaries(ctx context.Context) (summaries *Summary, err error) {

	urlPath := pathSummaries
	body, err := cl.curlPublic(ctx, urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetSummaries: " + err.Error())
	}

	printDebug(string(body))

	err = json.Unmarshal(body, &summaries)
	if err != nil {
		return nil, fmt.Errorf("GetSummaries: " + err.Error())
	}

	return summaries, nil
}

// GetPairs provides available pairs on exchange
func (cl *Client) GetPairs(ctx context.Context) (pairs *Pairs, err error) {
	urlPath := pathPairs
	body, err := cl.curlPublic(ctx, urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetPairs: " + err.Error())
	}

	printDebug(string(body))

	pairs = &Pairs{}
	err = pairs.UnmarshalJSON(body)
	if err != nil {
		return nil, fmt.Errorf("GetPairs: " + err.Error())
	}

	return pairs, nil
}

// GetPriceIncrements provide price increments of each pairs on exchange
func (cl *Client) GetPriceIncrements(ctx context.Context) (priceIncrements *PriceIncrements, err error) {
	urlPath := pathPriceIncrements
	body, err := cl.curlPublic(ctx, urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetPairs: " + err.Error())
	}

	printDebug(string(body))

	priceIncrements = &PriceIncrements{}
	err = priceIncrements.UnmarshalJSON(body)
	if err != nil {
		return nil, fmt.Errorf("GetPriceIncrements: " + err.Error())
	}

	return priceIncrements, nil
}
