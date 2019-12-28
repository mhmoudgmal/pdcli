package helpers

import (
	. "pdcli/backend/pd"

	ui "github.com/mhmoudgmal/termui"
)

// InspectIncidentWidget
func InspectIncidentWidget(incidentChan *chan Incident, iDWidget *ui.Table) {
	for {
		select {
		case incident := <-*incidentChan:
			mapIncidentDetailsToUI(incident, iDWidget)
		}
	}
}

func mapIncidentDetailsToUI(incident Incident, table *ui.Table) {
	details := incident.Inspect("details").([][]string)
	statusColor := detailsColor(incident.Status)

	var colors []ui.Attribute
	for range details {
		colors = append(colors, statusColor)
	}

	table.Rows = details
	table.BgColors = colors

	table.Analysis()
	table.SetSize()

	ui.Body.Align()
	ui.Render(table)
}

func detailsColor(status string) ui.Attribute {
	switch status {
	case "resolved":
		return ui.ColorGreen
	case "acknowledged":
		return ui.ColorYellow
	case "triggered":
		return ui.ColorRed
	}
	return ui.ColorDefault
}
