package main

import (
	"flag"
	"time"

	"github.com/mhmoudgmal/pdcli/config"
	"github.com/mhmoudgmal/pdcli/cui"
	"github.com/mhmoudgmal/pdcli/models"
)

var term = make(chan bool)                                         // application terminate.
var stopFrequesting = make(chan bool)                              // stop requesting pd api.
var failuresChannel = make(chan string)                            // send failure message when something goes wrong.
var incidentsChannel = make(chan []models.Incident)                // pass incidents between different application parts whenever you have one.
var updateStatusChannel = make(chan models.Incident)               // pass incident when it is updated, to update other parts in the application.
var incidentUpdatingChannel = make(chan models.UpdateIncidentInfo) // send incident update message for updating the incident status.

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
		StopFrequestingChannel: &stopFrequesting,
	}

	go cui.InitUI(&ctx)
	go PDWorker(&ctx)

	<-term
}
