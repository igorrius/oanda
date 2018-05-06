package oanda

import (
	"net/http"
)

type Client struct {
	opt *Options
	*http.Client
}

func NewPracticeClient(token AuthToken, id ClientId) (*Client, error) {
	return newClient(token, id, ApiEnvPractice)
}

func NewTradeClient(token AuthToken, id ClientId) (*Client, error) {
	return newClient(token, id, ApiEnvTrade)
}

func newClient(token AuthToken, id ClientId, clientMode int) (*Client, error) {
	apiMode, err := newApiMode(clientMode)
	opt := newOptions(token, id, apiMode)

	// TODO: create methods to create Client with non-default http.Client
	client := &Client{
		opt:    opt,
		Client: http.DefaultClient,
	}
	return client, err
}
