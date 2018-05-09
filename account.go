package oanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oanda/v20"
)

const (
	accountSummaryUrlPattern     = "/accounts/%s/summary"
	accountInstrumentsUrlPattern = "/accounts/%s/instruments"
)

type Account struct {
}

/*
Get summary for AccountResponse
/accounts/{accountID}/summary
*/
func (Account) Summary() (*v20.AccountResponse, error) {
	url := fmt.Sprintf(accountSummaryUrlPattern, opt.auth.ClientId)
	request, err := newRequest("GET", url)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	// TODO: extract method to {Do Request} and {Get Body From Response}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	account := &v20.AccountResponse{}
	if err := json.Unmarshal(body, account); err != nil {
		return nil, err
	}

	return account, err
}

func (Account) Instruments() (*v20.InstrumentsResponse, error) {
	url := fmt.Sprintf(accountInstrumentsUrlPattern, opt.auth.ClientId)
	request, err := newRequest("GET", url)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	// TODO: extract method to {Do Request} and {Get Body From Response}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	instruments := &v20.InstrumentsResponse{}
	if err := json.Unmarshal(body, instruments); err != nil {
		return nil, err
	}

	return instruments, err
}
