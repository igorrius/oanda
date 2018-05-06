package oanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oanda/v20"
)

const (
	accountSummaryUrlPattern = "/accounts/%s/summary"
)

type Account struct {
}

/*
Get summary for Account
/accounts/{accountID}/summary
*/
func (Account) Summary() (*v20.AccountSummary, error) {
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

	account := &v20.AccountSummary{}
	if err := json.Unmarshal(body, account); err != nil {
		return nil, err
	}

	return account, err
}
