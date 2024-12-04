package coinmarketcap

import "time"

type onlyStatusResponse struct {
	Status *ResponseStatus `json:"status"`
}

type PriceConversionRequest struct {
	Amount  float64 `json:"amount" url:"amount"`   // The quantity to be converted
	Symbol  string  `json:"symbol" url:"symbol"`   // The cryptocurrency ID (e.g., TRX) or symbol.
	Convert string  `json:"convert" url:"convert"` // The target currency for conversion, e.g. USDT.
}

type _PriceConversionResponse struct {
	Status *ResponseStatus      `json:"status"`
	Data   *DataPriceConversion `json:"data"`
}

type DataPriceConversion struct {
	Id          int       `json:"id"`
	Symbol      string    `json:"symbol"`
	Name        string    `json:"name"`
	Amount      int       `json:"amount"`
	LastUpdated time.Time `json:"last_updated"`
	Quote       struct {
		USDT *Quote `json:"USDT"`
	} `json:"quote"`
}

type Quote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

type ResponseStatus struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage any       `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
	Notice       any       `json:"notice"`
}
