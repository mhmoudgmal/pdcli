package pd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

// GetServices ...
func (Backend) GetServices(ctx *AppContext, teams []string) IServices {
	resourceURL := baseURL + "/services?" + constructTeamParams(teams)

	client, request := HTTPRequest(ctx, http.MethodGet, resourceURL, nil)

	res, getErr := client.Do(request)
	if getErr != nil {
		*ctx.FailuresChannel <- getErr.Error()
		return IServices{}
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return IServices{}
	}

	result := struct{ Services []PDService }{}

	if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return IServices{}
	}

	var pdServices IServices
	for _, service := range result.Services {
		pdServices = append(pdServices, service)
	}
	return pdServices
}

func constructTeamParams(teams []string) string {
	teamsIDs := ""
	for _, teamID := range teams {
		teamsIDs += "team_ids[]=" + teamID + "&"
	}
	return strings.TrimSuffix(teamsIDs, "&")
}
