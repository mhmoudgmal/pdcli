package cui

import (
	. "pdcli/config"

	ui "github.com/pdevine/termui"
)

func updateIncidentsWidget(ctx *AppContext, incidentList *ui.ListBox) {
	incidentsItems := []ui.Item{}

	for {
		select {
		case incs := <-*ctx.IncidentsChannel:
			func() {
				for _, incident := range incs {
					incidentsItems = append(
						incidentsItems,
						ui.Item{incident.ID, incident.Inspect()})
				}

				incidentList.Items = incidentsItems
				incidentList.Height = len(incidentsItems) + 2

				ui.Render(incidentList)
			}()

		case incident := <-*ctx.UpdateStatusChannel:
			func() {
				itemIndex := func() int {
					for idx, incItem := range incidentsItems {
						if incItem.ItemVal == incident.ID {
							return idx
						}
					}
					return -1
				}()

				if itemIndex == -1 {
					return
				}

				incidentList.Items[itemIndex] = ui.Item{incident.ID, incident.Inspect()}
				ui.Render(incidentList)
			}()
		}
	}
}
