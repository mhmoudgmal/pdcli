package pd_test

import (
	"encoding/json"
	"gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pdcli/backend/pd"
	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

var _ = Describe("PD backend API", func() {
	Describe("GetIncident", func() {
		var failuresChannel chan string
		var ctx AppContext

		BeforeEach(func() {
			failuresChannel = make(chan string)

			ctx = AppContext{
				FailuresChannel: &failuresChannel,
				Backend: Backend{
					Config{
						Token: "pd_token",
						Email: "foo@bar.baz",
					},
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
				result := struct{ Incident PDIncident }{}

				incident := ctx.Backend.GetIncident(&ctx, "PT4KHLK")
				json.Unmarshal([]byte(incidentString), &result)

				Expect(incident).To(Equal(result.Incident))
				Expect(gock.IsDone()).To(Equal(true))
			})

			It("does not send any messages to the failure chan", func() {
				ctx.Backend.GetIncident(&ctx, "PT4KHLK")

				Expect(*ctx.FailuresChannel).NotTo(Receive())
				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		It("sends unexpected JSON error message through the failures channel", func() {
			gock.New("https://api.pagerduty.com/incident").
				Get("/").
				Reply(400)

			go ctx.Backend.GetIncident(&ctx, "")

			Eventually(*ctx.FailuresChannel).Should(Receive(Equal("unexpected end of JSON input")))
			Expect(gock.IsDone()).To(Equal(true))
		})
	})
})
