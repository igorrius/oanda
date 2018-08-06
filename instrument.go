package oanda

import "oanda/v20"

const (
	instrumentsBaseUrl = "/instruments/"
)

type Instrument struct {
	Request
	RequestTools
}

/*
Get candles for the instrument
v3/instruments/EUR_USD/candles
*/
func (i *Instrument) Candles(candleRequest v20.CandlesRequest) (*v20.CandlesResponse, error) {
	// TODO: create request from v20.CandlesRequest struct
	url := instrumentsBaseUrl + candleRequest.Instrument + "/candles"
	request, err := i.newRequest(RequestMethodGet, url)
	if err != nil {
		return nil, err
	}

	account := &v20.CandlesResponse{}
	if err := i.unmarshalFromResponse(request, account); err != nil {
		return nil, err
	}

	return account, err
}
