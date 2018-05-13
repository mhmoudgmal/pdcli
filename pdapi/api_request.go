package pdapi

import (
	"io"
	"net/http"
	"time"

	"pdcli/config"
)

const baseURL = "https://api.pagerduty.com"

// APIRequest - PD Api request
func APIRequest(
	ctx *config.AppContext,
	method string,
	url string,
	data io.Reader,
) (http.Client, *http.Request) {

	client := http.Client{Timeout: time.Second * 2}

	req, reqErr := http.NewRequest(method, url, data)
	if reqErr != nil {
		panic(reqErr)
	}

	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Authorization", "Token token="+ctx.PDConfig.Token)

	return client, req
}
