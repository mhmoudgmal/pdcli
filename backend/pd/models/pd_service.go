package models

// PDService minimal representation
type PDService struct {
	ID      string `json:"id"`
	Summary string `json:"summary"`
}

// GetID ..
func (pds PDService) GetID() string {
	return pds.ID
}

// GetSummary ..
func (pds PDService) GetSummary() string {
	return pds.Summary
}
