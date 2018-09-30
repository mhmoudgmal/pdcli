package models

// PDService minimal representation
type PDService struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Summary      string `json:"summary"`
	LastIncident string `json:"last_incident_timestamp"`
}

// GetID ..
func (pds PDService) GetID() string {
	return pds.ID
}

// GetName ..
func (pds PDService) GetName() string {
	return " [⏽](fg-red) [⭘](fg-white) " + pds.Name + " | " + pds.LastIncident
}

// GetSummary ..
func (pds PDService) GetSummary() string {
	return pds.Summary
}

// GetSummary ..
func (pds PDService) GetLastIncident() string {
	return pds.LastIncident
}
