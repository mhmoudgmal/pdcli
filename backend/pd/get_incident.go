package pd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetIncident - requests incident information from PD service
func (backend Backend) GetIncident(id string) (Incident, error) {
	resourceURL := baseURL + "/incidents/" + id

	client, request := HTTPRequest(backend, http.MethodGet, resourceURL, nil)

	res, getErr := client.Do(request)
	if getErr != nil {
		return Incident{}, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return Incident{}, readErr
	}

	result := struct{ Incident Incident }{}

	if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
		return Incident{}, jsonErr
	}

	return result.Incident, nil
}
