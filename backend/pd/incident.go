package pd

import "fmt"

// ACKNOWLEDGED an incident status
const ACKNOWLEDGED = "acknowledged"

// RESOLVED an incident status
const RESOLVED = "resolved"

// TRIGGERED an incident status
const TRIGGERED = "triggered"

var incidentStatusColorMapper = map[string]string{
	TRIGGERED:    "(fg-red)",
	ACKNOWLEDGED: "(fg-yellow)",
	RESOLVED:     "(fg-green)",
}

// Service minimal representation for PD
type Service struct {
	Id      string `json:"id"`
	Summary string `json:"summary"`
}

// Incident minimal representation for PD
type Incident struct {
	Service `json:"service"`

	IncidentNumber int    `json:"incident_number"`
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
	Urgency        string `json:"urgency"`
	URL            string `json:"html_url"`
	Status         string `json:"status"`
}

// Inspect formats and colorize the incident according to its status
func (i Incident) Inspect(mode string) interface{} {
	switch mode {
	case "status-line":
		status := fmt.Sprintf("[⬤]%s", incidentStatusColorMapper[i.Status])
		return fmt.Sprintf("%s %s @ %s", status, i.ID, i.URL)

	case "details":
		return [][]string{
			[]string{"Status", i.Status},
			[]string{"Severity", i.Urgency},
			[]string{"Summary", i.Description},
			[]string{"Created", i.CreatedAt},
			[]string{"Service", i.Service.Summary},
		}

	default:
		return i.URL
	}
}
