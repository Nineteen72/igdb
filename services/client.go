package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PerformRequest performs a network request of type HTTPRequestType, using the passed in url, headers, and parameters.
func PerformRequest(apiKey, uri string, parameters url.Values, client *http.Client) ([]byte, error) {
	baseURL := fmt.Sprintf("%s?", uri)

	req, err := http.NewRequest(http.MethodGet, baseURL+parameters.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header, err = addHeaders(apiKey)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func addHeaders(key string) (http.Header, error) {
	if len(key) == 0 {
		return nil, errors.New("API key is missing")
	}

	h := http.Header{}
	h.Set("user-key", key)
	h.Set("Accept", "application/json")

	return h, nil
}
