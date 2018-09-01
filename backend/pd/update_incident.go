package pd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

// UpdateIncident - update the incident on PD service
func (Backend) UpdateIncident(ctx *AppContext, info UpdateIncidentInfo) IIncident {
	apiURL := baseURL + "/incidents/" + info.ID

	client, request := HTTPRequest(
		ctx,
		http.MethodPut,
		apiURL,
		bytes.NewReader(jsonBody(info.Status)),
	)

	request.Header.Set("From", info.Config.GetConfig().(Config).Email)
	request.Header.Set("Content-Type", "application/json")

	res, putErr := client.Do(request)

	if putErr != nil {
		*ctx.FailuresChannel <- putErr.Error()
		return PDIncident{}
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		*ctx.FailuresChannel <- readErr.Error()
		return PDIncident{}
	}

	result := struct{ Incident PDIncident }{}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		*ctx.FailuresChannel <- jsonErr.Error()
		return PDIncident{}
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
