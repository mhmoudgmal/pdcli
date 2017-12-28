package pdapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mhmoudgmal/pdcli/config"
	. "github.com/mhmoudgmal/pdcli/models"
)

// GetIncidents - requests the incidents from PD service
func GetIncidents(ctx *config.AppContext, options map[string]string) []Incident {
	client, request := APIRequest(ctx, http.MethodGet, buildURL(options), nil)

	res, getErr := client.Do(request)
	result := struct{ Incidents []Incident }{[]Incident{}}

	if getErr != nil {
		*ctx.FailuresChannel <- getErr.Error()
		return result.Incidents
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return result.Incidents
	}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return result.Incidents
	}

	return result.Incidents
}

func buildURL(options map[string]string) string {
	params := url.Values{}
	for pName, pValue := range options {
		params.Add(pName, pValue)
	}

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}
