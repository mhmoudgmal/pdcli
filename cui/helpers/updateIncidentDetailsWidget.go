package helpers

import (
	ui "github.com/pdevine/termui"

	"pdcli/config"
	"pdcli/models"
)

// UpdateIncidentDetailsWidget ...
func UpdateIncidentDetailsWidget(ctx *config.AppContext, iDWidget *ui.Table) {
	for {
		select {
		case incident := <-*ctx.IncidentDetailsChannel:
			mapIncidentDetailsToUI(incident, iDWidget)
		}
	}
}

func mapIncidentDetailsToUI(incident models.Incident, table *ui.Table) {
	data := [][]string{
		[]string{"Status", incident.Status},
		[]string{"Severity", incident.Urgency},
		[]string{"Summary", incident.Description},
		[]string{"Created", incident.CreatedAt},
		[]string{"Service", incident.Service.Summary},
	}

	table.Rows = data
	table.FgColor = detailsColor(incident.Status)
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
