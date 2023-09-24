package pd_test

import (
	"encoding/json"

	"gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mhmoudgmal/pdcli/backend/pd"
)

var _ = Describe("PD backend API", func() {
	Describe("GetIncident", func() {
		backend := Backend{
			Config{
				Token: "pd_token",
				Email: "foo@bar.baz",
			},
		}

		Context("when request succeeds", func() {
			response := `{
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
					BodyString(response)
			})

			It("returns the incident", func() {
				result := struct{ Incident Incident }{}

				incident, err := backend.GetIncident("PT4KHLK")
				json.Unmarshal([]byte(response), &result)

				Expect(err).To(BeNil())
				Expect(incident).To(Equal(result.Incident))

				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		It("returns error when bad request", func() {
			gock.New("https://api.pagerduty.com/incident").
				Get("/").
				Reply(400)

			incident, err := backend.GetIncident("")

			Expect(incident).To(Equal(Incident{}))
			Expect(err.Error()).To(Equal("unexpected end of JSON input"))

			Expect(gock.IsDone()).To(Equal(true))
		})
	})
})
