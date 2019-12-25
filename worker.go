package main

import (
	"fmt"
	. "pdcli/backend/pd"
	. "pdcli/notifiable"
)

func Worker(ctx *AppContext) {
	incidentsBackend := ctx.Backend

	go incidentsBackend.Worker(ctx.RequestInterval, func(incidents ...Incident) {
		*ctx.IncidentsChan <- incidents

		if ctx.Mode.Code == ModeA.Code {
			for _, incident := range incidents {
				*ctx.UpdateIncidentStatusChan <- struct {
					ID     string
					Status string
				}{
					ID:     incident.ID,
					Status: ACKNOWLEDGED,
				}
			}
		}

	})

	go func() {
		for {
			select {
			case updateIncidentInfo := <-*ctx.UpdateIncidentStatusChan:
				updatedIncident, err := incidentsBackend.UpdateIncident(updateIncidentInfo)
				if err != nil {
					fmt.Println("Error", err.Error())
					continue
				}

				ctx.Notifiable.Notify(
					func(data ...Incident) { *ctx.IncidentsChan <- data },
					updatedIncident,
				)
				ctx.Notifiable.Notify(
					func(data ...Incident) { *ctx.IncidentDetailsChan <- data[0] },
					updatedIncident,
				)
			}
		}
	}()

	go func() {
		for {
			select {
			case id := <-*ctx.InspectIncidentChan:
				incident, err := incidentsBackend.GetIncident(id)
				if err != nil {
					fmt.Println("Error", err.Error())
					continue
				}

				ctx.Notifiable.Notify(
					func(data ...Incident) { *ctx.IncidentDetailsChan <- data[0] },
					incident,
				)
			}
		}
	}()

	<-*ctx.StopWorkerChan
}
