package i_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pdcli/backend/pd/models"
	. "pdcli/i"

	"pdcli/backend/pd"
)

var _ = Describe("Incident Actions", func() {
	updateBackendChannel := make(chan UpdateIncidentInfo)

	var pdbe IncidentBackend

	Describe("Ack()", func() {
		BeforeEach(func() {
			pdbe = pd.Backend{
				pd.Config{
					Token: "token",
					Email: "foo@bar.baz",
				},
			}
		})

		It("Sends ACK message to the updateBackendChannel when status is TRIGGERED", func() {
			incident := PDIncident{ID: "I1", Status: TRIGGERED}

			go Ack(incident.GetID(), &updateBackendChannel, pdbe)
			Eventually(updateBackendChannel).Should(Receive(
				Equal(UpdateIncidentInfo{
					ID:     incident.GetID(),
					Status: ACKNOWLEDGED,
					Config: pdbe,
				}),
			))
		})

		It("Does not send any messages to the updateBackendChannel when status is not TRIGGERED", func() {
			statuses := []string{ACKNOWLEDGED, RESOLVED}

			for _, status := range statuses {
				incident := PDIncident{ID: "I1", Status: status}

				go Ack(incident.GetID(), &updateBackendChannel, pdbe)
				Eventually(updateBackendChannel).Should(Not(Receive()))
			}
		})
	})

	Describe("Resolve()", func() {
		It("sends RESOLVED message to the updateBackendChannel when status is not RESOLVED", func() {
			statuses := []string{TRIGGERED, ACKNOWLEDGED}

			for _, status := range statuses {
				incident := PDIncident{ID: "I1", Status: status}

				go Resolve(incident.GetID(), &updateBackendChannel, pdbe)
				Eventually(updateBackendChannel).Should(Receive(
					Equal(UpdateIncidentInfo{
						ID:     incident.ID,
						Status: RESOLVED,
						Config: pdbe,
					}),
				))
			}
		})

		It("does not send any messages to the updateBackendChannel when status is RESOLVED", func() {
			incident := PDIncident{ID: "I1", Status: RESOLVED}
			go Resolve(incident.GetID(), &updateBackendChannel, pdbe)

			Expect(updateBackendChannel).NotTo(Receive())
		})
	})
})
