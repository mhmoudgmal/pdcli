package pd

import (
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.pagerduty.com"

// HTTPRequest is a generic http request to pd APIs
func HTTPRequest(backend Backend, method string, url string, data io.Reader) (http.Client, *http.Request) {
	client := http.Client{Timeout: time.Second * 2}

	req, reqErr := http.NewRequest(method, url, data)
	if reqErr != nil {
		panic(reqErr)
	}

	token := backend.Token

	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Authorization", "Token token="+token)

	return client, req
}
