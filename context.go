package main

import (
	. "pdcli/backend/pd"
	. "pdcli/notifiable"

	"time"
)

// AppContext _
type AppContext struct {
	*Mode      // Application mode (e.g Normal, Auto-Ack)
	Backend    // PagerDuty
	Notifiable // Incident Notifiable is the downstream client that gets notified with incidents (e.g CUI, Siren)

	RequestInterval time.Duration // Frequent requests interval (e.g 2 seconds)

	FailuresChan   *chan string   // Receives failure messages
	TerminateChan  *chan struct{} // Receives signal to terminate application
	StopWorkerChan *chan struct{} // Receives signal to stop requesting new incidents

	IncidentsChan       *chan []Incident // Pass the received incidents between app components
	IncidentDetailsChan *chan Incident   // Pass incident details between app components
	InspectIncidentChan *chan string     // Receives an incident id to get the corresponding incident details

	// Receives a message to update an incident status in the backend
	UpdateIncidentStatusChan *chan struct {
		ID     string
		Status string
	}
}
