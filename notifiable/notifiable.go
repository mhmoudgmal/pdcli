package notifiable

import . "github.com/mhmoudgmal/pdcli/backend/pd"

// Notifiable represents the client that gets notified with incidents.
// Supported notifiable:
//   - Cui: a command line interface
//
// Any other notifiable or client should implement this interface.
type Notifiable interface {
	Init(
		appMode *Mode,
		terminateChan *chan struct{},
		stopWorkerChan *chan struct{},
		inspectIncidentChan *chan string,

		incidentChan *chan Incident,
		incidentsChan *chan []Incident,

		updateIncidentStatus *chan struct {
			ID     string
			Status string
		},
	)

	Notify(func(...Incident), ...Incident)
	Clean()
}
