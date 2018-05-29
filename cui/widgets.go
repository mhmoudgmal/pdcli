package cui

import (
	ui "github.com/pdevine/termui"

	"pdcli/config"
)

func helpWidget() *ui.Par {
	helpWidgetText := `
C-c     Quit
C-t     Toggle Auto-Ack mode
C-k     Acknowledge selected incident
C-r     Resolve selected incident
C-v     Show the details of the selected incident

[⬤](fg-green)  Resolved [⬤](fg-yellow)  Acknowledged [⬤](fg-red)  Triggered
`
	helpWidget := ui.NewPar(helpWidgetText)
	helpWidget.Height = 10
	helpWidget.Width = 50
	helpWidget.TextFgColor = ui.ColorWhite
	helpWidget.BorderLabel = "Help center"
	helpWidget.BorderFg = ui.ColorCyan

	return helpWidget
}

func modeWidget(cfg *config.AppContext) *ui.Gauge {
	modeWidget := ui.NewGauge()
	modeWidget.Percent = 100
	modeWidget.Label = cfg.Mode.Code
	modeWidget.BarColor = cfg.Mode.Color
	modeWidget.PercentColorHighlighted = ui.ColorDefault

	return modeWidget
}

func onCallStatusWidget(cfg *config.AppContext) *ui.Gauge {
	onCallStatusWidget := ui.NewGauge()
	onCallStatusWidget.Percent = 100
	onCallStatusWidget.Label = "..."
	onCallStatusWidget.BarColor = ui.ColorDefault
	onCallStatusWidget.PercentColorHighlighted = ui.ColorDefault

	return onCallStatusWidget
}

func incidentsWidget() *ui.ListBox {
	incidentItems := []ui.Item{}
	incidentsListBox := ui.NewListBox()
	incidentsListBox.Items = incidentItems
	incidentsListBox.ItemFgColor = ui.ColorYellow
	incidentsListBox.BorderLabel = "Incidents"
	incidentsListBox.Height = 2
	incidentsListBox.Width = 20
	incidentsListBox.Y = 0

	return incidentsListBox
}

func incidentDetailsWidget() *ui.Table {
	incidentDetails := ui.NewTable()
	incidentDetails.Border = true
	incidentDetails.TextAlign = ui.AlignCenter
	return incidentDetails
}

// Widgets ...
type Widgets struct {
	HelpWidget            *ui.Par
	ModeWidget            *ui.Gauge
	IncidentsWidget       *ui.ListBox
	OnCallStatusWidget    *ui.Gauge
	IncidentDetailsWidget *ui.Table
}
