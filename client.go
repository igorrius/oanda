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
	Account Account
}

func NewPracticeClient(token AuthToken, id ClientId) (*Client, error) {
	return newClient(token, id, ApiEnvPractice)
}

func NewTradeClient(token AuthToken, id ClientId) (*Client, error) {
	return newClient(token, id, ApiEnvTrade)
}

func newClient(token AuthToken, id ClientId, clientMode int) (*Client, error) {
	apiMode, err := newApiMode(clientMode)
	opt = newOptions(token, id, apiMode)
	httpClient = http.DefaultClient

	// TODO: create methods to create Client with non-default http.Client
	client := &Client{}
	return client, err
}
