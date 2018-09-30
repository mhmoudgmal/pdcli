package i

import "time"

// BackendChannels a set of channels related to the backend.
type BackendChannels struct {
	// Receives failure messages.
	FailuresChannel *chan string

	// Receives an id to get the corresponding incident details.
	GetIncidentChannel *chan string

	// Pass an updated incident between app components.
	UpdateStatusChannel *chan IIncident

	// Pass incident details between app components.
	IncidentDetailsChannel *chan IIncident

	// Pass the received incidents between app components.
	IncidentsChannel *chan IIncidents

	// Pass the received services between app components.
	ServicesChannel *chan IServices

	// Pass the received teams_ids.
	TeamsChannel *chan []string

	// Receives a message to update incident in the backend
	UpdateBackendChannel *chan UpdateIncidentInfo
}

// AppContext _
type AppContext struct {
	// Application mode
	//	- Normal
	//  - Auto-Ack
	Mode

	// Frequent requests interval (e.g 2 seconds)
	FrequestDuration time.Duration

	// Incident backend (e.g PagerDuty, VectorOps ..etc)
	Backend IncidentBackend

	// Incident Notifiable is the downstream client that gets notified with incidents (e.g CUI, Siren)
	Notifiable IncidentNotifiable

	// Receives signal to terminate application
	TerminateChannel *chan bool

	// Set of channels related to the backend.
	BackendChannels
}
