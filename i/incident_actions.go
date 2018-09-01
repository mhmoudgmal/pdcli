package i

type configurable interface {
	GetConfig() interface{}
}

// UpdateIncidentInfo - represents the message to be send when attempting to update the incident
type UpdateIncidentInfo struct {
	ID     string
	Status string
	Config configurable
}

// Ack - sends an ack message to the update incident channel
func Ack(incident IIncident, updateChan *chan UpdateIncidentInfo, config configurable) {
	if incident.GetStatus() == TRIGGERED {
		*updateChan <- UpdateIncidentInfo{
			ID:     incident.GetID(),
			Status: ACKNOWLEDGED,
			Config: config,
		}
	}
}

// Resolve - sends a resolve message to the update incident channel
func Resolve(incident IIncident, updateChan *chan UpdateIncidentInfo, config configurable) {
	if incident.GetStatus() != RESOLVED {
		*updateChan <- UpdateIncidentInfo{
			ID:     incident.GetID(),
			Status: RESOLVED,
			Config: config,
		}
	}
}
