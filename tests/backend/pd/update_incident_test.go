package pd_test

import (
	"fmt"

	. "github.com/mhmoudgmal/pdcli/backend/pd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
)

var _ = Describe("PD Backend API", func() {
	Describe("UpdateIncident", func() {
		backend := Backend{
			Config{
				Email: "foo@bar.baz",
				Token: "pd_token",
			},
		}

		response := `{
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

		Context("when request succeeds", func() {
			JustBeforeEach(func() {
				gock.New("https://api.pagerduty.com/incidents").
					HeaderPresent("Authorization").
					MatchHeader("Content-Type", "application/json").
					MatchHeader("Accept", `application/vnd.pagerduty\+json;version=2`).
					Put("/PT4KHLK").
					Reply(200).
					BodyString(fmt.Sprintf(response, ACKNOWLEDGED))
			})

			It("returns the updated incident", func() {
				incident, err := backend.UpdateIncident(
					struct {
						ID     string
						Status string
					}{
						ID:     "PT4KHLK",
						Status: ACKNOWLEDGED,
					},
				)

				Expect(err).To(BeNil())
				Expect(incident.Status).To(Equal(ACKNOWLEDGED))

				Expect(gock.IsDone()).To(Equal(true))
			})
		})

		It("returns error when bad request", func() {
			gock.New("https://api.pagerduty.com/incidents").
				Put("/PT4KHLK").
				Reply(400)

			incident, err := backend.UpdateIncident(
				struct {
					ID     string
					Status string
				}{
					ID:     "PT4KHLK",
					Status: ACKNOWLEDGED,
				},
			)

			Expect(incident).To(Equal(Incident{}))
			Expect(err.Error()).To(Equal("unexpected end of JSON input"))

			Expect(gock.IsDone()).To(Equal(true))
		})
	})
})
