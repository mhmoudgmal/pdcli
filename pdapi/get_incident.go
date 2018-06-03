package pdapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"pdcli/config"
	"pdcli/models"
)

type iResult struct {
	Incident models.Incident
}

// GetIncident - requests incident information from PD service
func GetIncident(ctx *config.AppContext, id string) models.Incident {
	apiURL := baseURL + "/incidents/" + id

	client, request := APIRequest(
		ctx,
		http.MethodGet,
		apiURL,
		nil,
	)

	res, getErr := client.Do(request)
	if getErr != nil {
		*ctx.FailuresChannel <- getErr.Error()
		return models.Incident{}
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return models.Incident{}
	}

	result := iResult{}
	if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return models.Incident{}
	}

	return result.Incident
}
