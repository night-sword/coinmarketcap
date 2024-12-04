package coinmarketcap

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var _apikey = os.Getenv("API_KEY")

func TestCoinMarketCap_PriceConversion(t *testing.T) {
	request := &PriceConversionRequest{
		Amount:  1,
		Symbol:  "TRX",
		Convert: "USDT",
	}
	response, err := getCoinMarketCap().PriceConversion(context.Background(), request)
	if err != nil {
		fmt.Println("fail", err)
	} else {
		fmt.Println("success", response.Quote.USDT)
	}

	request1 := &PriceConversionRequest{
		Amount:  -1,
		Symbol:  "TRX",
		Convert: "USDT",
	}
	response1, err := getCoinMarketCap().PriceConversion(context.Background(), request1)
	if err != nil {
		fmt.Println("fail", err)
	} else {
		fmt.Println("success", response1.Quote.USDT)
	}
}

func getCoinMarketCap() (cmc *CoinMarketCap) {
	cmc = NewCoinMarketCap("https://pro-api.coinmarketcap.com/", nil)
	cmc.SetKeys([]string{_apikey})

	return
}
