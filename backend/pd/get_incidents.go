package pd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

// GetIncidents requests the incidents from PD service
func (Backend) GetIncidents(ctx *AppContext, options map[string]string) IIncidents {
	resourceURL := baseURL + "/incidents"

	client, request := HTTPRequest(ctx, http.MethodGet, buildURL(resourceURL, options), nil)

	res, getErr := client.Do(request)
	if getErr != nil {
		*ctx.FailuresChannel <- getErr.Error()
		return IIncidents{}
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return IIncidents{}
	}

	result := struct{ Incidents []PDIncident }{}

	if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return IIncidents{}
	}

	var pdIncidents IIncidents
	for _, pdIncident := range result.Incidents {
		pdIncidents = append(pdIncidents, pdIncident)
	}
	return pdIncidents
}

func buildURL(resourceURL string, options map[string]string) string {
	params := url.Values{}
	for pName, pValue := range options {
		params.Add(pName, pValue)
	}

	return fmt.Sprintf("%s?%s", resourceURL, params.Encode())
}
