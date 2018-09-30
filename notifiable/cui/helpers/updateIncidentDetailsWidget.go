package helpers

import (
	"fmt"

	ui "github.com/mhmoudgmal/termui"

	. "pdcli/i"
)

// UpdateIncidentDetailsWidget ...
func UpdateIncidentDetailsWidget(ctx *AppContext, iDWidget *ui.Table, x *ui.Par) {
	for {
		select {
		case incident := <-*ctx.IncidentDetailsChannel:
			mapIncidentDetailsToUI(incident, iDWidget, x)
		}
	}
}

func mapIncidentDetailsToUI(incident IIncident, table *ui.Table, par *ui.Par) {
	data := [][]string{
		[]string{"Status", incident.GetStatus()},
		[]string{"Severity", incident.GetUrgency()},
		[]string{"Service", incident.GetService().GetSummary()},
		[]string{"Assigned to", incident.GetAssignedTo()},
		[]string{"Created", incident.GetCreatedAt()},
	}

	details := `%s

[%s](fg-yellow)
`

	par.Text = fmt.Sprintf(details, incident.GetDescription(), incident.GetURL())
	par.Border = true
	par.Height = 20
	ui.Render(par)

	table.Rows = data
	table.Border = true
	bgColor := detailsColor(incident.GetStatus())
	bgColors := []ui.Attribute{}
	for range data {
		bgColors = append(bgColors, bgColor)
	}
	table.BgColors = bgColors
	table.FgColor = ui.ColorWhite
	table.TextAlign = ui.AlignLeft
	table.Analysis()
	table.SetSize()
	ui.Body.Align()
	ui.Render(table)
	ui.Render(ui.Body)
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
