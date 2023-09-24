package pd_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mhmoudgmal/pdcli/backend/pd"
)

var _ = Describe("Incident", func() {
	var inspect func(string, string, string) string

	BeforeEach(func() {
		inspect = func(color string, id string, url string) string {
			return fmt.Sprintf("[⬤](fg-%s) %s @ %s", color, id, url)
		}
	})

	Describe("Inspect()", func() {
		Context("details", func() {
			incident := Incident{
				ID:          "I1",
				Description: "I1 Desc",
				URL:         "http://incident.backend.test/incidents/ti1",
				Status:      TRIGGERED,
			}

			It("returns incident data in [][] format", func() {
				Expect(incident.Inspect("details")).To(Equal(
					[][]string{
						[]string{"Status", incident.Status},
						[]string{"Severity", incident.Urgency},
						[]string{"Summary", incident.Description},
						[]string{"Created", incident.CreatedAt},
						[]string{"Service", incident.Service.Summary},
					},
				))
			})
		})

		Context("status-line", func() {
			It("returns inspected incident with RED indicator when status is TRIGGERED", func() {
				triggeredIncident := Incident{
					ID:          "I1",
					Description: "I1 Desc",
					URL:         "http://incident.backend.test/incidents/ti1",
					Status:      TRIGGERED,
				}

				Expect(triggeredIncident.Inspect("status-line")).To(Equal(
					inspect(
						"red",
						triggeredIncident.ID,
						triggeredIncident.URL,
					),
				))
			})

			It("returns inspected incident with YELLOW indicator when status is ACKNOWLEDED", func() {
				acknowlededIncident := Incident{
					ID:          "I2",
					Description: "I2 Desc",
					URL:         "http://incident.backend.test/incidents/ai1",
					Status:      ACKNOWLEDGED,
				}

				Expect(acknowlededIncident.Inspect("status-line")).To(Equal(
					inspect(
						"yellow",
						acknowlededIncident.ID,
						acknowlededIncident.URL,
					),
				))
			})

			It("returns inspected incident with GREEN indicator when status is RESOLVED", func() {
				resolvedIncident := Incident{
					ID:          "I3",
					Description: "I3 Desc",
					URL:         "http://incident.backend.test/incidents/ri1",
					Status:      RESOLVED,
				}

				Expect(resolvedIncident.Inspect("status-line")).To(Equal(
					inspect(
						"green",
						resolvedIncident.ID,
						resolvedIncident.URL,
					),
				))
			})
		})
	})
})
