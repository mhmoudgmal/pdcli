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

var _ = Describe("PD Backend API", func() {
	Describe("GetIncidents", func() {
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
			incidentsString := `{
				"incidents":
					[
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
					]
				}`

			JustBeforeEach(func() {
				gock.New("https://api.pagerduty.com/incidents").
					HeaderPresent("Authorization").
					HeaderPresent("Accept").
					Get("/").
					Reply(200).
					BodyString(incidentsString)
			})

			It("returns the incidents", func() {
				result := struct{ Incidents []PDIncident }{}

				incidents := ctx.Backend.GetIncidents(&ctx, map[string]string{})
				json.Unmarshal([]byte(incidentsString), &result)

				Expect(incidents[0].(PDIncident)).To(Equal(result.Incidents[0]))
				Expect(gock.IsDone()).To(Equal(true))
			})

			It("does not send any messages to the failure chan", func() {
				ctx.Backend.GetIncidents(&ctx, map[string]string{})

				Expect(*ctx.FailuresChannel).NotTo(Receive())
				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		Context("when bad request", func() {
			It("sends error message through the failures channel", func() {
				gock.New("https://api.pagerduty.com/incidents").
					Get("/").
					Reply(400)

				go ctx.Backend.GetIncidents(&ctx, map[string]string{})

				Eventually(*ctx.FailuresChannel).Should(Receive(Equal("unexpected end of JSON input")))
				Expect(gock.IsDone()).To(Equal(true))
			})
		})
	})
})
