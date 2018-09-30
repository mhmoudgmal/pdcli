package models

import (
	. "pdcli/i"
	"strings"
)

// PDIncident minimal representation
type PDIncident struct {
	IncidentNumber int          `json:"incident_number"`
	ID             string       `json:"id"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	CreatedAt      string       `json:"created_at"`
	Urgency        string       `json:"urgency"`
	HTMLURL        string       `json:"html_url"`
	Status         string       `json:"status"`
	Assignments    []Assignment `json:"assignments"`
	PDService      `json:"service"`
}

// GetIncidentNumber ..
func (pdi PDIncident) GetIncidentNumber() int {
	return pdi.IncidentNumber
}

// GetID ..
func (pdi PDIncident) GetID() string {
	return pdi.ID
}

// GetTitle ..
func (pdi PDIncident) GetTitle() string {
	return pdi.Title
}

// GetDescription ..
func (pdi PDIncident) GetDescription() string {
	return pdi.Description
}

// GetCreatedAt ..
func (pdi PDIncident) GetCreatedAt() string {
	return pdi.CreatedAt
}

// GetUrgency ..
func (pdi PDIncident) GetUrgency() string {
	return pdi.Urgency
}

// GetURL ..
func (pdi PDIncident) GetURL() string {
	return pdi.HTMLURL
}

// GetAssignedTo ..
func (pdi PDIncident) GetAssignedTo() string {
	assingees := []string{}
	for _, ass := range pdi.Assignments {
		assingees = append(assingees, ass.AssignedTo())
	}
	return strings.Join(assingees, ",")
}

// GetStatus ..
func (pdi PDIncident) GetStatus() string {
	return pdi.Status
}

// GetService ..
func (pdi PDIncident) GetService() IService {
	return pdi.PDService
}
