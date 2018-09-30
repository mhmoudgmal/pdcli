package helpers

import (
	ui "github.com/mhmoudgmal/termui"

	. "pdcli/i"
)

// UpdateIncidentsWidget ...
func UpdateIncidentsWidget(ctx *AppContext, lb *ui.ListBox) {
	incidentsItems := []ui.Item{}

	for {
		select {
		case incidents := <-*ctx.IncidentsChannel:
			incidentsItems = append(incidentsItems, mapIncidentsToUIItems(incidents)...)
			updateIncidentListBox(lb, incidentsItems)

		case incident := <-*ctx.UpdateStatusChannel:
			updateIncidentStatus(lb, incidentsItems, incident)
		}
	}
}

func mapIncidentsToUIItems(incidents IIncidents) []ui.Item {
	incidentsItems := []ui.Item{}

	for _, incident := range incidents {
		incidentsItems = append(
			incidentsItems, ui.Item{
				ItemVal: incident.GetID(),
				Text:    Inspect(incident),
				Data:    map[string]string{"url": incident.GetURL()},
			},
		)
	}

	return incidentsItems
}

func updateIncidentListBox(lb *ui.ListBox, incidentsItems []ui.Item) {
	lb.Items = incidentsItems
	lb.Height = lb.Height + len(incidentsItems)
	ui.Body.Align()
	ui.Render(lb)
	ui.Render(ui.Body)
}

func updateIncidentStatus(lb *ui.ListBox, incidentsItems []ui.Item, incident IIncident) {
	itemIndex := getItemIndex(incidentsItems, incident)

	if itemIndex >= 0 {
		lb.Items[itemIndex] = ui.Item{
			ItemVal: incident.GetID(),
			Text:    Inspect(incident),
		}
		ui.Render(lb)
		ui.Render(ui.Body)
	}
}

func getItemIndex(incidentsItems []ui.Item, incident IIncident) int {
	itemIndex := -1

	for idx, incidentItem := range incidentsItems {
		if incidentItem.ItemVal == incident.GetID() {
			itemIndex = idx
		}
	}

	return itemIndex
}
