package oanda

import (
	"fmt"
	"oanda/v20"
)

const (
	accountBaseUrl              = "/accounts/%s"
	accountSummaryUrlSuffix     = "/summary"
	accountInstrumentsUrlSuffix = "/instruments"
)

type Account struct {
	Request
	RequestTools
}

func (a Account) getUrl(suffix string) string {
	url := fmt.Sprintf(accountBaseUrl, opt.auth.ClientId)
	return url + suffix
}

/*
Get summary for AccountResponse
/accounts/{accountID}/summary
*/
func (a Account) Summary() (*v20.AccountResponse, error) {
	url := a.getUrl(accountSummaryUrlSuffix)
	request, err := a.newRequest(RequestMethodGet, url)
	if err != nil {
		return nil, err
	}

	account := &v20.AccountResponse{}
	if err := a.unmarshalFromResponse(request, account); err != nil {
		return nil, err
	}

	return account, err
}

/*
Get instruments for AccountResponse
/accounts/{accountID}/instruments
*/
func (a Account) Instruments() (*v20.InstrumentsResponse, error) {
	url := a.getUrl(accountInstrumentsUrlSuffix)
	request, err := a.newRequest(RequestMethodGet, url)
	if err != nil {
		return nil, err
	}

	instruments := &v20.InstrumentsResponse{}
	if err := a.unmarshalFromResponse(request, instruments); err != nil {
		return nil, err
	}

	return instruments, err
}
