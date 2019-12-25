package helpers

import (
	ui "github.com/pdevine/termui"

	. "pdcli/backend/pd"
)

// UpdateIncidentsWidget ...
func UpdateIncidentsWidget(lb *ui.ListBox, incidentsChan *chan []Incident) {
	incidentItemMap := map[string]ui.Item{}

	for {
		select {
		case incidents := <-*incidentsChan:
			incidentsItems := mapIncidentsToUIItems(incidentItemMap, incidents...)
			updateIncidentListBox(lb, incidentsItems)
		}
	}
}

func mapIncidentsToUIItems(incidentItemMap map[string]ui.Item, incidents ...Incident) []ui.Item {
	for _, incident := range incidents {
		incidentItemMap[incident.ID] = ui.Item{
			ItemVal: incident.ID,
			Text:    incident.Inspect("status-line").(string),
		}
	}

	var incidentsItems []ui.Item
	for _, item := range incidentItemMap {
		incidentsItems = append(incidentsItems, item)
	}

	return incidentsItems
}

func updateIncidentListBox(lb *ui.ListBox, incidentsItems []ui.Item) {
	lb.Items = incidentsItems
	lb.Height = len(incidentsItems) + 2 // margin the listbox height
	ui.Render(lb)
}
