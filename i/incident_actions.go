package i

// UpdateIncidentInfo is the message to be sent on updating the incident
type UpdateIncidentInfo struct {
	ID     string
	Status string
	Config Configurable
}

// Ack sends an ack message to UpdateIncidentInfo channel
func Ack(id string, updateChan *chan UpdateIncidentInfo, config Configurable) {
	*updateChan <- UpdateIncidentInfo{
		ID:     id,
		Status: ACKNOWLEDGED,
		Config: config,
	}
}

// Resolve sends a resolve message to UpdateIncidentInfo channel
func Resolve(id string, updateChan *chan UpdateIncidentInfo, config Configurable) {
	*updateChan <- UpdateIncidentInfo{
		ID:     id,
		Status: RESOLVED,
		Config: config,
	}
}
