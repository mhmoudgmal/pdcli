package pdapi_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"

	"github.com/mhmoudgmal/pdcli/config"
	"github.com/mhmoudgmal/pdcli/models"
	"github.com/mhmoudgmal/pdcli/pdapi"
)

var _ = Describe("pdapi", func() {
	Describe("UpdateIncident", func() {
		var failuresChannel chan string
		var ctx config.AppContext
		var incident struct{ Incident models.Incident }
		var incidentString string

		BeforeEach(func() {
			failuresChannel = make(chan string)
			ctx = config.AppContext{FailuresChannel: &failuresChannel, PDConfig: &config.PDConfig{Email: "foo@bar.baz", Token: "pd_token"}}
			incident = struct{ Incident models.Incident }{models.Incident{}}

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

			json.Unmarshal([]byte(fmt.Sprintf(incidentString, "triggered")), &incident)
		})

		Context("when 200", func() {
			It("returns the updated incident", func() {
				gock.New("https://api.pagerduty.com/incidents/"+incident.Incident.ID).
					HeaderPresent("Authorization").
					MatchHeader("Content-Type", "application/json").
					MatchHeader("Accept", `application/vnd.pagerduty\+json;version=2`).
					Put("/").
					Reply(200).
					BodyString(fmt.Sprintf(incidentString, "acknowledged"))

				resultIncident := pdapi.UpdateIncident(
					&ctx, models.UpdateIncidentInfo{
						ID:     incident.Incident.ID,
						From:   ctx.PDConfig.Email,
						Status: "acknowledged",
					},
				)

				Expect(*ctx.FailuresChannel).NotTo(Receive())
				Expect(resultIncident.Status).To(Equal("acknowledged"))

				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		Context("when 400", func() {
			It("sends error message through the failures channel", func() {
				gock.New("https://api.pagerduty.com/incidents/" + incident.Incident.ID).
					Put("/").
					Reply(400)

				go pdapi.UpdateIncident(
					&ctx, models.UpdateIncidentInfo{
						ID:     incident.Incident.ID,
						From:   ctx.PDConfig.Email,
						Status: "acknowledged",
					},
				)

				Eventually(*ctx.FailuresChannel).Should(Receive(Equal("unexpected end of JSON input")))
				Expect(gock.IsDone()).To(Equal(true))
			})
		})
	})
})
