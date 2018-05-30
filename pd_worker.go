package main

import (
	"time"

	"pdcli/config"
	"pdcli/models"
	"pdcli/pdapi"
)

// PDWorker .. starts a ticker asking for incidents every 2/? seconds,
// Sends the result through IncidentsChannel.
// Auto Acknowledge the incidents if the auto-ack mode is enabled.
func PDWorker(ctx *config.AppContext) {
	go func() {
		params := map[string]string{"since": ""}

		for _ = range time.Tick(ctx.FrequestDuration) {
			incidents := pdapi.GetIncidents(ctx, params)

			if len(incidents) > 0 {
				t, _ := time.Now().MarshalText()
				params["since"] = string(t)

				*ctx.IncidentsChannel <- incidents

				if ctx.Mode.Code == config.ModeA.Code {
					autoAck(ctx, incidents)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case updateIncidentInfo := <-*ctx.PDUpdatingChannel:
				incident := pdapi.UpdateIncident(ctx, updateIncidentInfo)
				*ctx.UpdateStatusChannel <- incident
				*ctx.IncidentDetailsChannel <- incident
			}
		}
	}()

	go func() {
		for {
			select {
			case incidentID := <-*ctx.PDGetIncidentChannel:
				*ctx.IncidentDetailsChannel <- pdapi.GetIncident(ctx, incidentID)
			}
		}
	}()

	<-*ctx.StopFrequestingChannel
}

func autoAck(ctx *config.AppContext, incidents []models.Incident) {
	for _, incident := range incidents {
		incident.Ack(ctx.PDUpdatingChannel, ctx.PDConfig.Email)
	}
}
