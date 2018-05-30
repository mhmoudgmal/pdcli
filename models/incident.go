package models

// ACKNOWLEDGED - incident status
const ACKNOWLEDGED = "acknowledged"

// RESOLVED - incident status
const RESOLVED = "resolved"

// TRIGGERED - incident status
const TRIGGERED = "triggered"

var incidentStatusColorMapper = map[string]string{
	TRIGGERED:    "(fg-red)",
	ACKNOWLEDGED: "(fg-yellow)",
	RESOLVED:     "(fg-green)",
}

// Incident minimal representation
type Incident struct {
	IncidentNumber int    `json:"incident_number"`
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
	Urgency        string `json:"urgency"`
	HTMLURL        string `json:"html_url"`
	Status         string `json:"status"`
	Service        `json:"service"`
}

// Service minimal representation
type Service struct {
	ID      string `json:"id"`
	Summary string `json:"summary"`
}

// IncidentUpdateInfo - represents the message to be send when attempting to update the incident
type IncidentUpdateInfo struct {
	ID     string
	From   string
	Status string
}

// Inspect - format and colorize the incident according to its status
func (incident *Incident) Inspect() string {
	return "[⬤]" + incidentStatusColorMapper[incident.Status] + " " + incident.ID + " @ " + incident.HTMLURL
}

// TODO: extract incident behaviours to different module/package

// Ack - sends a message to the update incident channel
func (incident *Incident) Ack(updateChan *chan IncidentUpdateInfo, from string) {
	if incident.Status == TRIGGERED {
		*updateChan <- IncidentUpdateInfo{
			incident.ID,
			from,
			ACKNOWLEDGED,
		}
	}
}

// Resolve - sends a message to the update incident channel
func (incident *Incident) Resolve(updateChan *chan IncidentUpdateInfo, from string) {
	if incident.Status != RESOLVED {
		*updateChan <- IncidentUpdateInfo{
			incident.ID,
			from,
			RESOLVED,
		}
	}
}
