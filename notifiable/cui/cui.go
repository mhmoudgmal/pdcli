package cui

import . "pdcli/i"

// Cui the command line interface notifiable or client.
type Cui struct {
	*AppContext
	Widgets
}

// Notify Cui
func (c Cui) Notify(msg string, data interface{}) {
	switch msg {
	case "new-incidents":
		*c.AppContext.IncidentsChannel <- data.(IIncidents)
	case "updated-incident":
		*c.AppContext.UpdateStatusChannel <- data.(IIncident)
	case "detailed-incident":
		*c.AppContext.IncidentDetailsChannel <- data.(IIncident)
	}
}
