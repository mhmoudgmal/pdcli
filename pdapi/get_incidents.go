package pdapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"pdcli/config"
	"pdcli/models"
)

// GetIncidents - requests the incidents from PD service
func GetIncidents(ctx *config.AppContext, options map[string]string) []models.Incident {
	apiURL := baseURL + "/incidents"

	client, request := APIRequest(
		ctx,
		http.MethodGet,
		buildURL(apiURL, options),
		nil,
	)

	res, getErr := client.Do(request)
	result := struct{ Incidents []models.Incident }{[]models.Incident{}}

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

func buildURL(apiURL string, options map[string]string) string {
	params := url.Values{}
	for pName, pValue := range options {
		params.Add(pName, pValue)
	}

	return fmt.Sprintf("%s?%s", apiURL, params.Encode())
}
