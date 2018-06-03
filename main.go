package main

import (
	"flag"
	"time"

	ui "github.com/pdevine/termui"

	"pdcli/config"
	"pdcli/cui"
	"pdcli/models"
)

var (
	term                    = make(chan bool)                      // application terminate.
	stopFrequesting         = make(chan bool)                      // stop requesting pd api.
	failuresChannel         = make(chan string)                    // send failure message when something goes wrong.
	incidentsChannel        = make(chan []models.Incident)         // pass incidents between different application parts whenever you have one.
	updateStatusChannel     = make(chan models.Incident)           // pass incident when it is updated, to update other parts in the application.
	incidentDetailsChannel  = make(chan models.Incident)           // pass incident details between different application parts whenever you have one.
	getIncidentChannel      = make(chan string)                    // pass incident id to get the corresponding incident details from PD.
	incidentUpdatingChannel = make(chan models.IncidentUpdateInfo) // send incident update message for updating the incident status.
)

func main() {
	// TODO: refactor and extract flags/options to different package, and also support to read/serialize the options vlaues from/to file.
	email := flag.String("email", "", "Your pagerduty account email")
	token := flag.String("token", "", "Your pagerduty api access token")

	flag.Parse()

	ctx := config.AppContext{
		Mode:                   &config.ModeM,
		PDConfig:               &config.PDConfig{Token: *token, Email: *email},
		TermChannel:            &term,
		FailuresChannel:        &failuresChannel,
		FrequestDuration:       2 * time.Second,
		IncidentsChannel:       &incidentsChannel,
		PDUpdatingChannel:      &incidentUpdatingChannel,
		UpdateStatusChannel:    &updateStatusChannel,
		IncidentDetailsChannel: &incidentDetailsChannel,
		PDGetIncidentChannel:   &getIncidentChannel,
		StopFrequestingChannel: &stopFrequesting,
	}

	go cui.InitUI(&ctx)
	go PDWorker(&ctx)

	<-term
	ui.Clear()
}
