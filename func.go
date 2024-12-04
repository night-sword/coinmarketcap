package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/night-sword/kratos-kit/errors"
)

func httpGetParams[Request, Response any](ctx context.Context, client *resty.Client, api string, req *Request) (rsp *Response, err error) {
	if req != nil {
		r, e := query.Values(req)
		if err = e; err != nil {
			err = errors.BadRequest(errors.RsnParams, "build query params error").Degrade().AddMetadata("err", err.Error())
			return
		}

		api = api + "?" + r.Encode()
	}

	r, err := client.R().
		SetContext(ctx).
		Get(api)
	if err != nil {
		return
	}

	return decodeJsonResponse[Response](r)
}

func decodeJsonResponse[T any](response *resty.Response) (rsp *T, err error) {
	if response.StatusCode() != 200 {
		var status *onlyStatusResponse
		err = json.Unmarshal(response.Body(), &status)
		if err != nil {
			err = errors.BadRequest(errors.RsnAccessRepoFail, "decode response status fail").AsWarn().WithCause(err)
			return
		}

		msg := fmt.Sprintf("%v", status.Status.ErrorMessage)
		err = errors.BadRequest(errors.RsnAccessRepoFail, msg).AsWarn().AddMetadata("error_code", status.Status.ErrorCode)
		return
	}

	err = json.Unmarshal(response.Body(), &rsp)
	if err != nil {
		err = errors.BadRequest(errors.RsnAccessRepoFail, "decode response fail").AsWarn().WithCause(err)
	}

	return
}
