package models_test

import (
	"fmt"

	. "github.com/mhmoudgmal/pdcli/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Models/Incident", func() {
	var incidentAsString func(string, string, string) string
	var updateChannel chan IncidentUpdateInfo

	BeforeEach(func() {
		updateChannel = make(chan IncidentUpdateInfo)

		incidentAsString = func(color string, id string, htmlURL string) string {
			return fmt.Sprintf("[⬤](fg-%s) %s @ %s", color, id, htmlURL)
		}

	})

	Describe("Inspect()", func() {
		Context("With 'triggered' Status", func() {
			It("Should be the red dot + description + htmlURL", func() {
				triggeredIncident := Incident{
					ID:          "I1",
					Description: "I1 Desc",
					HTMLURL:     "http://pagerduty.test/incidents/ti1",
					Status:      "triggered",
				}
				Expect(triggeredIncident.Inspect()).To(Equal(incidentAsString("red", triggeredIncident.ID, triggeredIncident.HTMLURL)))
			})
		})

		Context("With 'acknowledged' Status", func() {
			It("Should be the yellow dot + description + htmlURL", func() {
				acknowlededIncident := Incident{
					ID:          "I2",
					Description: "I2 Desc",
					HTMLURL:     "http://pagerduty.test/incidents/ai1",
					Status:      "acknowledged",
				}
				Expect(acknowlededIncident.Inspect()).To(Equal(incidentAsString("yellow", acknowlededIncident.ID, acknowlededIncident.HTMLURL)))
			})
		})

		Context("With 'resolved' Status", func() {
			It("Should be the green dot + description + htmlURL", func() {
				resolvedIncident := Incident{
					ID:          "I3",
					Description: "I3 Desc",
					HTMLURL:     "http://pagerduty.test/incidents/ri1",
					Status:      "resolved",
				}
				Expect(resolvedIncident.Inspect()).To(Equal(incidentAsString("green", resolvedIncident.ID, resolvedIncident.HTMLURL)))
			})
		})
	})

	Describe("Ack()", func() {
		Context("When status is triggered", func() {
			It("Sends ack message to the IncidentUpdateInfo channel", func() {
				incident := Incident{ID: "I1", Status: "triggered"}
				go incident.Ack(&updateChannel, "foo@bar.baz")

				Eventually(updateChannel).Should(Receive(
					Equal(IncidentUpdateInfo{ID: incident.ID, Status: "acknowledged", From: "foo@bar.baz"}),
				))
				//// The DIY way!
				//
				// Expect(<-updateChannel).To(ContainElement())
				// select {
				// case updateIncidentInfo := <-updateChannel:
				// 	Expect(updateIncidentInfo.ID).To(Equal(incident.ID))
				// 	Expect(updateIncidentInfo.Status).To(Equal("acknowledged"))
				// }
				//
				// OR --
				//
				// time.Sleep(time.Millisecond * 10)
				// Expect(updateChannel).To(Receive(Equal(...))
			})
		})

		Context("When status is not triggered", func() {
			It("Does not send ack message to the UpdadateIncidentInfo channel", func() {
				incident := Incident{ID: "Incident"}
				go incident.Ack(&updateChannel, "foo@bar.baz")

				Expect(updateChannel).NotTo(Receive())
				//// The DIY way
				//
				// time.Sleep(10 * time.Millisecond)
				// Expect(len(updateChannel)).To(Equal(1))
			})
		})
	})

	Describe("Resolve()", func() {
		Context("When status is not 'resolved", func() {
			It("Should send 'resolved' message to the updateChannel", func() {
				incident := Incident{ID: "I1", Status: "acknowledged"}
				go incident.Resolve(&updateChannel, "foo@bar.baz")

				Eventually(updateChannel).Should(Receive(
					Equal(IncidentUpdateInfo{ID: incident.ID, Status: "resolved", From: "foo@bar.baz"}),
				))
			})
		})

		Context("When status is 'resolved'", func() {
			It("Should not send any messages to the updateChannel", func() {
				incident := Incident{ID: "I1", Status: "resolved"}
				go incident.Resolve(&updateChannel, "foo@bar.baz")

				Expect(updateChannel).NotTo(Receive())
			})
		})
	})
})
