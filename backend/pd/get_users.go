package pd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

// GetIncident - requests incident information from PD service
func (Backend) GetUsers(ctx *AppContext) IUsers {
	resourceURL := baseURL + "/users"

	client, request := HTTPRequest(ctx, http.MethodGet, resourceURL, nil)

	res, getErr := client.Do(request)
	if getErr != nil {
		*ctx.FailuresChannel <- getErr.Error()
		return IUsers{}
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return IUsers{}
	}

	result := struct{ Users []PDUser }{}

	if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return IUsers{}
	}

	var pdUsers IUsers
	for _, pdUser := range result.Users {
		pdUsers = append(pdUsers, pdUser)
	}

	return pdUsers
}
