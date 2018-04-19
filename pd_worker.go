package main

import (
	"time"

	. "pdcli/config"
	"pdcli/models"
	"pdcli/pdapi"
)

// PDWorker - starts a ticker every 2/? seconds gets incidents if any and sends the result through IncidentsChannel,
// Auto Acknowledge the incidents if the auto-ack mode is enabled.
func PDWorker(ctx *AppContext) {
	go func() {
		params := map[string]string{"since": ""}

		for _ = range time.Tick(ctx.FrequestDuration) {
			incidents := pdapi.GetIncidents(ctx, params)

			if len(incidents) > 0 {
				t, _ := time.Now().MarshalText()
				params["since"] = string(t)

				*ctx.IncidentsChannel <- incidents

				if ctx.Mode.Code == ModeA.Code {
					autoAck(ctx, incidents)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case updateIncidentInfo := <-*ctx.PDUpdatingChannel:
				*ctx.UpdateStatusChannel <- pdapi.UpdateIncident(ctx, updateIncidentInfo)
			}
		}
	}()

	<-*ctx.StopFrequestingChannel
}

func autoAck(ctx *AppContext, incidents []models.Incident) {
	for _, incident := range incidents {
		incident.Ack(ctx.PDUpdatingChannel, ctx.PDConfig.Email)
	}
}
