package oanda

import (
	"errors"
)

const (
	ApiEnvPractice = iota + 1
	ApiEnvTrade
)

const (
	apiPracticeUrl = "https://api-fxpractice.oanda.com/"
	apiTradeUrl    = "https://api-fxtrade.oanda.com/"
	apiVersion     = "v3"
)

type ApiMode struct {
	url string
}

func newApiMode(modeType int) (*ApiMode, error) {
	var apiUrl string

	switch modeType {
	case ApiEnvPractice:
		apiUrl = apiPracticeUrl
	case ApiEnvTrade:
		apiUrl = apiTradeUrl
	default:
		return nil, errors.New("invalid mode type")
	}

	apiUrl = apiUrl + apiVersion

	return &ApiMode{url: apiUrl}, nil
}
