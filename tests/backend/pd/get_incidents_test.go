package pd_test

import (
	"encoding/json"
	"gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pdcli/backend/pd"
)

var _ = Describe("PD Backend API", func() {
	Describe("GetIncidents", func() {
		backend := Backend{
			Config{
				Token: "pd_token",
				Email: "foo@bar.baz",
			},
		}

		Context("when request succeeds", func() {
			response := `{
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
					BodyString(response)
			})

			It("returns the incidents", func() {
				result := struct{ Incidents []Incident }{}

				incidents, err := backend.GetIncidents(map[string]string{})
				json.Unmarshal([]byte(response), &result)

				Expect(err).To(BeNil())
				Expect(incidents).To(Equal(result.Incidents))

				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		It("returns error when bad request", func() {
			gock.New("https://api.pagerduty.com/incidents").
				Get("/").
				Reply(400)

			incidents, err := backend.GetIncidents(map[string]string{})

			Expect(incidents).To(Equal([]Incident{}))
			Expect(err.Error()).To(Equal("unexpected end of JSON input"))

			Expect(gock.IsDone()).To(Equal(true))
		})
	})
})
