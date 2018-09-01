package i

// Configurable ..
type Configurable interface {
	GetConfig() interface{}
}

// IncidentBackend represents the ServiceProvider's capabilities for an incident management.
// PagerDuty is an example of incident-management service provider.
type IncidentBackend interface {
	Configurable
	// List incidents.
	GetIncidents(ctx *AppContext, options map[string]string) IIncidents
	// Get a specific incident details.
	GetIncident(ctx *AppContext, id string) IIncident
	// Update an incident info (e.g incident-status).
	UpdateIncident(ctx *AppContext, info UpdateIncidentInfo) IIncident
}
