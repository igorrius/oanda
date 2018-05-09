package oanda

import (
	"fmt"
	"oanda/v20"
)

const (
	accountSummaryUrlPattern     = "/accounts/%s/summary"
	accountInstrumentsUrlPattern = "/accounts/%s/instruments"
)

type Account struct {
	Api
}

/*
Get summary for AccountResponse
/accounts/{accountID}/summary
*/
func (a Account) Summary() (*v20.AccountResponse, error) {
	url := fmt.Sprintf(accountSummaryUrlPattern, opt.auth.ClientId)
	request, err := a.newRequest("GET", url)
	if err != nil {
		return nil, err
	}

	account := &v20.AccountResponse{}
	if err := a.unmarshalFromRequest(request, account); err != nil {
		return nil, err
	}

	return account, err
}

func (a Account) Instruments() (*v20.InstrumentsResponse, error) {
	url := fmt.Sprintf(accountInstrumentsUrlPattern, opt.auth.ClientId)
	request, err := a.newRequest("GET", url)
	if err != nil {
		return nil, err
	}

	instruments := &v20.InstrumentsResponse{}
	if err := a.unmarshalFromRequest(request, instruments); err != nil {
		return nil, err
	}

	return instruments, err
}
