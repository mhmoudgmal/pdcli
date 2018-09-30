package i

type IUsers []IUser
type IUser interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetTeams() []string
}

// Configurable ..
type Configurable interface {
	GetConfig() interface{}
}

// IncidentBackend represents the ServiceProvider's capabilities for an incident management.
// PagerDuty is an example of incident-management service provider.
type IncidentBackend interface {
	Configurable
	// Get users
	GetUsers(*AppContext) IUsers
	// List services.
	GetServices(*AppContext, []string) IServices
	// List incidents.
	GetIncidents(*AppContext, map[string]string) IIncidents
	// Get a specific incident details.
	GetIncident(*AppContext, string) IIncident
	// Update an incident info (e.g incident-status).
	UpdateIncident(*AppContext, UpdateIncidentInfo) IIncident
}
