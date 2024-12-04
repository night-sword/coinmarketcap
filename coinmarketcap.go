package coinmarketcap

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/night-sword/kratos-kit/errors"
)

type CoinMarketCap struct {
	*apiKeyWRR
	client *resty.Client
}

func NewCoinMarketCap(endpoint string, client *resty.Client) *CoinMarketCap {
	if client == nil {
		client = NewHttpClient(endpoint)
	}

	return &CoinMarketCap{
		apiKeyWRR: newApiKeyWRR([]string{}),
		client:    client,
	}
}

func (inst *CoinMarketCap) PriceConversion(ctx context.Context, req *PriceConversionRequest) (response *DataPriceConversion, err error) {
	api := "/v1/tools/price-conversion"

	client := inst.clientWithKey()

	rsp, err := httpGetParams[PriceConversionRequest, _PriceConversionResponse](ctx, client, api, req)
	if err != nil {
		return
	}

	err = inst.checkStatus(rsp.Status)
	if err != nil {
		return
	}

	response = rsp.Data
	return
}

func (inst *CoinMarketCap) clientWithKey() *resty.Client {
	key := inst.getKey()
	return inst.client.SetHeader(API_KEY_HEADER_NAME, key)
}

func (inst *CoinMarketCap) checkStatus(status *ResponseStatus) (err error) {
	if status.ErrorCode != 0 {
		err = errors.BadRequest(errors.RsnAccessRepoFail, fmt.Sprintf("%v", status.ErrorMessage))
	}
	return
}
