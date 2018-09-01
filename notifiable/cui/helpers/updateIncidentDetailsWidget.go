package helpers

import (
	ui "github.com/pdevine/termui"

	. "pdcli/i"
)

// UpdateIncidentDetailsWidget ...
func UpdateIncidentDetailsWidget(ctx *AppContext, iDWidget *ui.Table) {
	for {
		select {
		case incident := <-*ctx.IncidentDetailsChannel:
			mapIncidentDetailsToUI(incident, iDWidget)
		}
	}
}

func mapIncidentDetailsToUI(incident IIncident, table *ui.Table) {
	data := [][]string{
		[]string{"Status", incident.GetStatus()},
		[]string{"Severity", incident.GetUrgency()},
		[]string{"Summary", incident.GetDescription()},
		[]string{"Created", incident.GetCreatedAt()},
		[]string{"Service", incident.GetService().GetSummary()},
	}

	table.Rows = data
	table.FgColor = detailsColor(incident.GetStatus())
	table.Analysis()
	table.SetSize()
	ui.Body.Align()
	ui.Render(table)
}

func detailsColor(status string) ui.Attribute {
	switch status {
	case "resolved":
		return ui.ColorCyan
	case "acknowledged":
		return ui.ColorYellow
	case "triggered":
		return ui.ColorRed
	}
	return ui.ColorDefault
}
