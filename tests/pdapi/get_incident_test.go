package pdapi_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"

	"pdcli/config"
	"pdcli/models"
	"pdcli/pdapi"
)

var _ = Describe("pdapi", func() {
	Describe("GetIncident", func() {
		var failuresChannel chan string
		var ctx config.AppContext

		BeforeEach(func() {
			failuresChannel = make(chan string)

			ctx = config.AppContext{
				FailuresChannel: &failuresChannel,
				PDConfig: &config.PDConfig{
					Email: "foo@bar.baz",
					Token: "pd_token",
				},
			}
		})

		Context("when request succeeds", func() {
			incidentString := `{
				"incident":

					{
						"id": "PT4KHLK",
						"type": "incident",
						"html_url": "https://subdomain.pagerduty.com/incidents/PT4KHLK",
						"created_at": "2015-10-06T21:30:42Z",
						"status": "triggered",
						"service": {
							"id": "PIJ90N7",
							"summary": "My Mail Service"
						}
					}

				}`

			JustBeforeEach(func() {
				gock.New("https://api.pagerduty.com/incident").
					HeaderPresent("Authorization").
					HeaderPresent("Accept").
					Get("/PT4KHLK").
					Reply(200).
					BodyString(incidentString)
			})

			It("returns the incident", func() {
				result := struct{ Incident models.Incident }{models.Incident{}}

				incident := pdapi.GetIncident(&ctx, "PT4KHLK")
				json.Unmarshal([]byte(incidentString), &result)

				Expect(incident).To(Equal(result.Incident))
				Expect(gock.IsDone()).To(Equal(true))
			})

			It("does not send any messages to the failure chan", func() {
				pdapi.GetIncident(&ctx, "PT4KHLK")

				Expect(*ctx.FailuresChannel).NotTo(Receive())
				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		It("sends unexpected JSON error message through the failures channel", func() {
			gock.New("https://api.pagerduty.com/incident").
				Get("/").
				Reply(400)

			go pdapi.GetIncident(&ctx, "")

			Eventually(*ctx.FailuresChannel).Should(Receive(Equal("unexpected end of JSON input")))
			Expect(gock.IsDone()).To(Equal(true))
		})
	})
})
