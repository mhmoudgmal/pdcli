package i

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

// IService interface of a backend representation to the service entity or concept
type IService interface {
	GetID() string
	GetSummary() string
}

// IIncident interface of a backend representation to the incident entity or concept
type IIncident interface {
	GetIncidentNumber() int
	GetID() string
	GetTitle() string
	GetDescription() string
	GetCreatedAt() string
	GetUrgency() string
	GetURL() string
	GetStatus() string
	GetService() IService
}

// IIncidents is the collection of incidents
type IIncidents []IIncident

// Inspect formats and colorize the incident according to its status
func Inspect(incident IIncident) string {
	return "[⬤]" + incidentStatusColorMapper[incident.GetStatus()] + " " + incident.GetID() + " @ " + incident.GetURL()
}
