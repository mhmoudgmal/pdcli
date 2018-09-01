package pd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

// GetIncident - requests incident information from PD service
func (Backend) GetIncident(ctx *AppContext, id string) IIncident {
	resourceURL := baseURL + "/incidents/" + id

	client, request := HTTPRequest(ctx, http.MethodGet, resourceURL, nil)

	res, getErr := client.Do(request)
	if getErr != nil {
		*ctx.FailuresChannel <- getErr.Error()
		return PDIncident{}
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return PDIncident{}
	}

	result := struct{ Incident PDIncident }{}

	if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return PDIncident{}
	}

	return result.Incident
}
