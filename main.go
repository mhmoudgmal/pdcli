package main

import (
	"flag"
	"time"

	. "pdcli/backend/pd"
	. "pdcli/notifiable"
	. "pdcli/notifiable/Cui"
)

var (
	reqInterval         = 2 * time.Second
	stopWorkerChan      = make(chan struct{})
	terminateChan       = make(chan struct{})
	failuresChan        = make(chan string)
	inspectIncidentChan = make(chan string)
	incidentDetailsChan = make(chan Incident)
	incidentsChan       = make(chan []Incident)
	updateIncidentChan  = make(chan struct {
		ID     string
		Status string
	})
)

func main() {
	email := flag.String("email", "", "Your pagerduty account email")
	token := flag.String("token", "", "Your pagerduty api access token")
	notifiable := flag.String("notifiable", "cui", "Incident Notifiable")

	flag.Parse()

	var ctx AppContext

	incidentNotifiable := notifiableFor(*notifiable, &ctx)
	incidentBackend := Backend{
		Config{
			Token: *token,
			Email: *email,
		},
	}

	ctx = AppContext{
		Backend:         incidentBackend,
		Notifiable:      incidentNotifiable,
		RequestInterval: reqInterval,

		Mode:           &ModeN,
		FailuresChan:   &failuresChan,
		TerminateChan:  &terminateChan,
		StopWorkerChan: &stopWorkerChan,

		IncidentsChan:       &incidentsChan,
		IncidentDetailsChan: &incidentDetailsChan,
		InspectIncidentChan: &inspectIncidentChan,

		UpdateIncidentStatusChan: &updateIncidentChan,
	}

	go incidentNotifiable.Init(
		ctx.Mode,
		ctx.TerminateChan,
		ctx.StopWorkerChan,
		ctx.InspectIncidentChan,
		ctx.IncidentDetailsChan,
		ctx.IncidentsChan,
		ctx.UpdateIncidentStatusChan,
	)
	go Worker(&ctx)

	<-*ctx.TerminateChan
	incidentNotifiable.Clean()
}

func notifiableFor(n string, ctx *AppContext) Notifiable {
	switch n {
	case "cui":
		return Cui{}
	default:
		panic("notifiable is not supported")
	}
}
