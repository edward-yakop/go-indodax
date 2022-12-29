package indodax

import "encoding/json"

type Pair struct {
	Id                     string      `json:"id"`
	Symbol                 string      `json:"symbol"`
	BaseCurrency           string      `json:"base_currency"`
	TradedCurrency         string      `json:"traded_currency"`
	TradedCurrencyUnit     string      `json:"traded_currency_unit"`
	Description            string      `json:"description"`
	TickerId               string      `json:"ticker_id"`
	VolumePrecision        int         `json:"volume_precision"`
	PricePrecision         int         `json:"price_precision"`
	PriceRound             int         `json:"price_round"`
	Pricescale             int         `json:"pricescale"`
	TradeMinBaseCurrency   int         `json:"trade_min_base_currency"`
	TradeMinTradedCurrency float64     `json:"trade_min_traded_currency"`
	HasMemo                bool        `json:"has_memo"`
	MemoName               interface{} `json:"memo_name"` // It could be either string or bool
	TradeFeePercent        float64     `json:"trade_fee_percent"`
	UrlLogo                string      `json:"url_logo"`
	UrlLogoPng             string      `json:"url_logo_png"`
	IsMaintenance          int         `json:"is_maintenance"`
	IsMarketSuspended      int         `json:"is_market_suspended"`
	CoingeckoId            string      `json:"coingecko_id"`
	CmcId                  interface{} `json:"cmc_id"` // It could either be int or bool
}

type Pairs struct {
	Pairs map[string]Pair
}

func (ps *Pairs) UnmarshalJSON(b []byte) (err error) {
	var pairs []Pair
	err = json.Unmarshal(b, &pairs)
	if err != nil {
		return err
	}

	ps.Pairs = make(map[string]Pair)
	for _, pair := range pairs {
		ps.Pairs[pair.Symbol] = pair
	}

	return nil
}
