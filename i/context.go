package i

import "time"

// AppContext _
type AppContext struct {
	Mode                                            // Application mode (e.g Normal, Auto-Ack)
	FrequestDuration       time.Duration            // Frequent requests interval (e.g 2 seconds)
	Backend                IncidentBackend          // Incident backend (e.g PagerDuty, VectorOps ..etc)
	Notifiable             IncidentNotifiable       // Incident Notifiable is the downstream client that gets notified with incidents (e.g CUI, Siren)
	TerminateChannel       *chan bool               // Receives signal to terminate application
	StopFrequestingChannel *chan bool               // Receives signal to stop asking for incidents
	FailuresChannel        *chan string             // Receives failure messages
	GetIncidentChannel     *chan string             // Receives an id to get the corresponding incident details
	UpdateStatusChannel    *chan IIncident          // Pass an updated incident between app components
	IncidentDetailsChannel *chan IIncident          // Pass incident details between app components
	IncidentsChannel       *chan IIncidents         // Pass the received incidents between app components
	UpdateBackendChannel   *chan UpdateIncidentInfo // Receives a message to update incident in the backend
}
