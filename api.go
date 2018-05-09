package oanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
}

func (Api) newRequest(method string, apiMethodUrl string) (*http.Request, error) {
	requestUrl := opt.mode.url + apiMethodUrl
	r, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", opt.auth.Token))
	return r, nil
}

func (Api) unmarshalFromRequest(request *http.Request, v interface{}) error {
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}
