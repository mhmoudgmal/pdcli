package main

import (
	"flag"
	"time"

	pd "pdcli/backend/pd"

	. "pdcli/i"
	. "pdcli/notifiable/Cui"
)

var (
	reqInterval            = 5 * time.Second
	terminateChannel       = make(chan bool)
	failuresChannel        = make(chan string)
	getIncidentChannel     = make(chan string)
	updateStatusChannel    = make(chan IIncident)
	incidentDetailsChannel = make(chan IIncident)
	incidentsChannel       = make(chan IIncidents)
	servicesChannel        = make(chan IServices)
	teamsChannel           = make(chan []string)
	updateIncidentChannel  = make(chan UpdateIncidentInfo)
)

func main() {
	backend := flag.String("backend", "pagerduty", "Incident managment backend")
	notifiable := flag.String("notifiable", "cui", "Incident Notifiable")

	// required by pd backend.
	email := flag.String("email", "", "Your pagerduty account email")
	token := flag.String("token", "", "Your pagerduty api access token")

	flag.Parse()

	var (
		ctx                AppContext
		incidentBackend    IncidentBackend
		incidentNotifiable IncidentNotifiable
	)

	incidentBackend = backendFor(
		*backend,
		map[string]string{
			"token": *token,
			"email": *email,
		},
	)
	incidentNotifiable = notifiableFor(*notifiable, &ctx)

	ctx = AppContext{
		Mode:             ModeN,
		Backend:          incidentBackend,
		Notifiable:       incidentNotifiable,
		FrequestDuration: reqInterval,
		TerminateChannel: &terminateChannel,

		BackendChannels: BackendChannels{
			TeamsChannel:           &teamsChannel,
			ServicesChannel:        &servicesChannel,
			FailuresChannel:        &failuresChannel,
			IncidentsChannel:       &incidentsChannel,
			GetIncidentChannel:     &getIncidentChannel,
			UpdateStatusChannel:    &updateStatusChannel,
			UpdateBackendChannel:   &updateIncidentChannel,
			IncidentDetailsChannel: &incidentDetailsChannel,
		},
	}

	go incidentNotifiable.Init(&ctx)
	go PDWorker(&ctx)

	<-*ctx.TerminateChannel
	incidentNotifiable.Clean()
}

func backendFor(b string, opts map[string]string) IncidentBackend {
	switch b {
	case "pagerduty":
		return pd.Backend{
			pd.Config{
				Token: opts["token"],
				Email: opts["email"],
			},
		}
	default:
		panic("backend is not supported")
	}
}

func notifiableFor(n string, ctx *AppContext) IncidentNotifiable {
	switch n {
	case "cui":
		return Cui{
			AppContext: ctx,
		}
	default:
		panic("notifiable is not supported")
	}
}
