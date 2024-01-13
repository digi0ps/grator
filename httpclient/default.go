package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DefaultClient struct {
	client *http.Client
}

func (d DefaultClient) MakeRequest(method, url, body string, headers map[string]string) (string, int, error) {
	fmt.Printf("Doing %s %s Body = %s\n", method, url, body)
	// Check if the client is nil and return an error if it is
	if d.client == nil {
		return "", 0, fmt.Errorf("HTTP client is not initialized")
	}

	var req *http.Request
	var err error

	if body != "" {
		bodyReader := strings.NewReader(body)
		req, err = http.NewRequest(method, url, bodyReader)
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if err != nil {
		return "", 0, err
	}

	res, err := d.client.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", res.StatusCode, err
	}

	return string(responseBody), res.StatusCode, nil
}

func NewDefaultClient(client *http.Client) *DefaultClient {
	return &DefaultClient{
		client: client,
	}
}
