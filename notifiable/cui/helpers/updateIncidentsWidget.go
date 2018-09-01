package helpers

import (
	ui "github.com/pdevine/termui"

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
			},
		)
	}

	return incidentsItems
}

func updateIncidentListBox(lb *ui.ListBox, incidentsItems []ui.Item) {
	lb.Items = incidentsItems
	lb.Height = len(incidentsItems) + 2
	ui.Render(lb)
}

func updateIncidentStatus(lb *ui.ListBox, incidentsItems []ui.Item, incident IIncident) {
	itemIndex := getItemIndex(incidentsItems, incident)

	if itemIndex >= 0 {
		lb.Items[itemIndex] = ui.Item{
			ItemVal: incident.GetID(),
			Text:    Inspect(incident),
		}
		ui.Render(lb)
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
