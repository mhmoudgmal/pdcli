package main

import (
	"time"

	pd "pdcli/backend/pd"
	. "pdcli/i"
)

// PDWorker ..starts a ticker asking for incidents every 2/? seconds,
// Sends the result through the IncidentsChannel.
// Auto Acknowledge the incidents if the auto-ack mode is enabled.
func PDWorker(ctx *AppContext) {
	incidentsBackend := ctx.Backend

	go func() {
		params := map[string]string{"since": ""}

		for _ = range time.Tick(ctx.FrequestDuration) {
			incidents := incidentsBackend.GetIncidents(ctx, params)

			t, _ := time.Now().MarshalText()
			params["since"] = string(t)

			if len(incidents) == 0 {
				continue
			}

			ctx.Notifiable.Notify("new-incidents", incidents)

			if ctx.Mode.Code == ModeA.Code {
				for _, incident := range incidents {
					Ack(
						incident.GetID(),
						ctx.UpdateBackendChannel,
						ctx.Backend,
					)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case updateIncidentInfo := <-*ctx.UpdateBackendChannel:
				if incident := incidentsBackend.UpdateIncident(ctx, updateIncidentInfo); incident.GetStatus() != "" {
					ctx.Notifiable.Notify("updated-incident", incident)
					ctx.Notifiable.Notify("detailed-incident", incident)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case incidentID := <-*ctx.GetIncidentChannel:
				incident := incidentsBackend.GetIncident(ctx, incidentID)
				ctx.Notifiable.Notify("detailed-incident", incident)
			}
		}
	}()

	go func() {
		var teams []string

		select {
		case teams = <-*ctx.TeamsChannel:
		case <-time.After(time.Duration(5 * time.Second)):
			teams = []string{}
		}

		services := incidentsBackend.GetServices(ctx, teams)
		ctx.Notifiable.Notify("list-services", services)
	}()

	go func() {
		users := incidentsBackend.GetUsers(ctx)
		for _, user := range users {
			if user.GetEmail() == ctx.Backend.GetConfig().(pd.Config).Email {
				*ctx.TeamsChannel <- user.GetTeams()
			}
		}

	}()
}
