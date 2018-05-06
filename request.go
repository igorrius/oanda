package oanda

import (
	"fmt"
	"net/http"
)

func newRequest(method string, apiMethodUrl string) (*http.Request, error) {
	requestUrl := opt.mode.url + apiMethodUrl
	r, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", opt.auth.Token))
	return r, nil
}
