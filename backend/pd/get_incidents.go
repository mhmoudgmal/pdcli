package pd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetIncidents requests the incidents from PD service
func (backend Backend) GetIncidents(params map[string]string) ([]Incident, error) {
	resourceURL := baseURL + "/incidents"

	client, request := HTTPRequest(backend, http.MethodGet, buildURL(resourceURL, params), nil)

	res, getErr := client.Do(request)
	if getErr != nil {
		return []Incident{}, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return []Incident{}, readErr
	}

	result := struct{ Incidents []Incident }{}

	if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
		return []Incident{}, jsonErr
	}

	return result.Incidents, nil
}

func buildURL(resourceURL string, params map[string]string) string {
	urlParams := url.Values{}
	for pName, pValue := range params {
		urlParams.Add(pName, pValue)
	}

	return fmt.Sprintf("%s?%s", resourceURL, urlParams.Encode())
}
