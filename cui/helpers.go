package cui

import (
	ui "github.com/pdevine/termui"

	"pdcli/config"
	"pdcli/models"
)

func updateIncidentsWidget(ctx *config.AppContext, lb *ui.ListBox) {
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

func mapIncidentsToUIItems(incidents []models.Incident) []ui.Item {
	incidentsItems := []ui.Item{}

	for _, incident := range incidents {
		incidentsItems = append(
			incidentsItems, ui.Item{
				ItemVal: incident.ID,
				Text:    incident.Inspect(),
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

func updateIncidentStatus(lb *ui.ListBox, incidentsItems []ui.Item, incident models.Incident) {
	itemIndex := getItemIndex(incidentsItems, incident)

	if itemIndex >= 0 {
		lb.Items[itemIndex] = ui.Item{
			ItemVal: incident.ID,
			Text:    incident.Inspect(),
		}
		ui.Render(lb)
	}
}

func getItemIndex(incidentsItems []ui.Item, incident models.Incident) int {
	itemIndex := -1

	for idx, incidentItem := range incidentsItems {
		if incidentItem.ItemVal == incident.ID {
			itemIndex = idx
		}
	}

	return itemIndex
}
