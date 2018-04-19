package pdapi

import (
	. "pdcli/config"
	"pdcli/models"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UpdateIncident - update the incident on PD service
func UpdateIncident(ctx *AppContext, info models.IncidentUpdateInfo) models.Incident {
	client, request := APIRequest(
		ctx,
		http.MethodPut,
		baseURL+info.ID,
		bytes.NewReader(jsonBody(info.Status)),
	)

	request.Header.Set("From", info.From)
	request.Header.Set("Content-Type", "application/json")

	res, putErr := client.Do(request)
	result := struct{ Incident models.Incident }{models.Incident{}}

	if putErr != nil {
		*ctx.FailuresChannel <- putErr.Error()
		return result.Incident
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return result.Incident
	}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return result.Incident
	}

	return result.Incident
}

func jsonBody(status string) []byte {
	return []byte(fmt.Sprintf(`
		{
			"incident": {
				"type": "incident_reference",
				"status": "%s"
			}
		}
	`, status))
}
