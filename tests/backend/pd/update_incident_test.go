package pd_test

import (
	"encoding/json"
	"fmt"
	"gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pdcli/backend/pd"
	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

var _ = Describe("PD Backend API", func() {
	Describe("UpdateIncident", func() {
		var failuresChannel chan string
		var ctx AppContext
		var incident struct{ Incident PDIncident }
		var incidentString string

		BeforeEach(func() {
			failuresChannel = make(chan string)

			ctx = AppContext{
				FailuresChannel: &failuresChannel,
				Backend: Backend{
					Config{
						Email: "foo@bar.baz",
						Token: "pd_token",
					},
				},
			}

			incident = struct{ Incident PDIncident }{}

			incidentString = `{
				"incident": {
							"id": "PT4KHLK",
							"type": "incident",
							"html_url": "https://subdomain.pagerduty.com/incidents/PT4KHLK",
							"created_at": "2015-10-06T21:30:42Z",
							"status": "%s",
							"service": {
								"id": "PIJ90N7",
								"summary": "My Mail Service"
							}
						}
					}`

			json.Unmarshal([]byte(fmt.Sprintf(incidentString, TRIGGERED)), &incident)
		})

		Context("when request succeeds", func() {
			var resultIncident IIncident

			JustBeforeEach(func() {
				gock.New("https://api.pagerduty.com/incidents/"+incident.Incident.GetID()).
					HeaderPresent("Authorization").
					MatchHeader("Content-Type", "application/json").
					MatchHeader("Accept", `application/vnd.pagerduty\+json;version=2`).
					Put("/").
					Reply(200).
					BodyString(fmt.Sprintf(incidentString, ACKNOWLEDGED))

				resultIncident = ctx.Backend.UpdateIncident(
					&ctx,
					UpdateIncidentInfo{
						ID:     incident.Incident.GetID(),
						Status: ACKNOWLEDGED,
						Config: ctx.Backend,
					},
				)
			})

			It("returns the updated incident", func() {
				Expect(resultIncident.GetStatus()).To(Equal(ACKNOWLEDGED))
				Expect(gock.IsDone()).To(Equal(true))
			})

			It("does not send any messages to the failure chan", func() {
				Expect(*ctx.FailuresChannel).NotTo(Receive())
				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		Context("when bad request", func() {
			It("sends error message through the failures channel", func() {
				gock.New("https://api.pagerduty.com/incidents/" + incident.Incident.GetID()).
					Put("/").
					Reply(400)

				go ctx.Backend.UpdateIncident(
					&ctx, UpdateIncidentInfo{
						ID:     incident.Incident.GetID(),
						Status: ACKNOWLEDGED,
						Config: ctx.Backend,
					},
				)

				Eventually(*ctx.FailuresChannel).Should(Receive(Equal("unexpected end of JSON input")))
				Expect(gock.IsDone()).To(Equal(true))
			})
		})
	})
})
