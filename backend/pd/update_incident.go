package pd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UpdateIncident - update the incident on PD service
func (backend Backend) UpdateIncident(info struct {
	ID     string
	Status string
}) (Incident, error) {
	apiURL := baseURL + "/incidents/" + info.ID

	client, request := HTTPRequest(
		backend,
		http.MethodPut,
		apiURL,
		bytes.NewReader(jsonBody(info.Status)),
	)

	from := backend.Email

	request.Header.Set("From", from)
	request.Header.Set("Content-Type", "application/json")

	res, putErr := client.Do(request)

	if putErr != nil {
		return Incident{}, putErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return Incident{}, readErr
	}

	result := struct{ Incident Incident }{}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return Incident{}, jsonErr
	}

	return result.Incident, nil
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
