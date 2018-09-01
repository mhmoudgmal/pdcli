package i_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pdcli/backend/pd/models"
	. "pdcli/i"
)

var _ = Describe("Incident", func() {
	var inspect func(string, string, string) string

	BeforeEach(func() {
		inspect = func(color string, id string, url string) string {
			return fmt.Sprintf("[⬤](fg-%s) %s @ %s", color, id, url)
		}
	})

	Describe("Inspect()", func() {
		It("returns inspected incident with RED indicator when status is TRIGGERED", func() {
			triggeredIncident := PDIncident{
				ID:          "I1",
				Description: "I1 Desc",
				HTMLURL:     "http://incident.backend.test/incidents/ti1",
				Status:      TRIGGERED,
			}

			Expect(Inspect(triggeredIncident)).To(Equal(
				inspect(
					"red",
					triggeredIncident.GetID(),
					triggeredIncident.GetURL(),
				),
			))
		})

		It("returns inspected incident with YELLOW indicator when status is ACKNOWLEDED", func() {
			acknowlededIncident := PDIncident{
				ID:          "I2",
				Description: "I2 Desc",
				HTMLURL:     "http://incident.backend.test/incidents/ai1",
				Status:      ACKNOWLEDGED,
			}

			Expect(Inspect(acknowlededIncident)).To(Equal(
				inspect(
					"yellow",
					acknowlededIncident.GetID(),
					acknowlededIncident.GetURL(),
				),
			))
		})

		It("returns inspected incident with GREEN indicator when status is RESOLVED", func() {
			resolvedIncident := PDIncident{
				ID:          "I3",
				Description: "I3 Desc",
				HTMLURL:     "http://incident.backend.test/incidents/ri1",
				Status:      RESOLVED,
			}

			Expect(Inspect(resolvedIncident)).To(Equal(
				inspect(
					"green",
					resolvedIncident.GetID(),
					resolvedIncident.GetURL(),
				),
			))
		})
	})
})
