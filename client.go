package oanda

import (
	"net/http"
)

var (
	opt        *Options
	httpClient *http.Client
)

type Client struct {
	*http.Client
	Account    Account
	Instrument Instrument
}

func NewPracticeClient(token AuthToken, id ClientId) (*Client, error) {
	return newClient(token, id, ApiEnvPractice, nil)
}

func NewTradeClient(token AuthToken, id ClientId) (*Client, error) {
	return newClient(token, id, ApiEnvTrade, nil)
}

func newClient(
	token AuthToken,
	id ClientId,
	clientMode int,
	customHttpClient *http.Client,
) (*Client, error) {
	apiMode, err := newApiMode(clientMode)
	opt = newOptions(token, id, apiMode)

	if nil != customHttpClient {
		httpClient = customHttpClient
	} else {
		httpClient = http.DefaultClient
	}

	// TODO: create methods to create Client with non-default http.Client
	client := &Client{}
	return client, err
}
