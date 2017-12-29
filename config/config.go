package config

import (
	"time"

	"github.com/mhmoudgmal/pdcli/models"
	ui "github.com/pdevine/termui"
)

var (
	// ModeA - Auto-Ack mode
	ModeA = Mode{
		"Auto-Ack",
		"\n All incidents are going to be acknowledged automatically.",
		ui.ColorYellow,
	}

	// ModeM - Manual mode
	ModeM = Mode{
		"Manual",
		"\n You are responsible to acknowledge all incoming incidents.",
		ui.ColorUndef,
	}
)

// Mode - Application modes (auto/manual)
type Mode struct {
	Code  string
	Note  string
	Color ui.Attribute
}

// PDConfig - PagerDuty configuration
type PDConfig struct {
	Token string
	Email string
}

// AppContext - Application configuration
type AppContext struct {
	PDConfig               *PDConfig
	Mode                   *Mode
	TermChannel            *chan bool
	FailuresChannel        *chan string
	IncidentsChannel       *chan []models.Incident
	FrequestDuration       time.Duration
	PDUpdatingChannel      *chan models.UpdateIncidentInfo
	UpdateStatusChannel    *chan models.Incident
	StopFrequestingChannel *chan bool
}
