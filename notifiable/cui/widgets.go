package cui

import (
	ui "github.com/mhmoudgmal/termui"

	. "pdcli/i"
)

func helpWidget() *ui.Par {
	helpWidgetText := `
C-c   Quit
C-t   Toggle Auto-Ack mode
C-k   Acknowledge selected incident
C-r   Resolve selected incident
C-v   Show the details of the selected incident

[⬤](fg-green)  Resolved [⬤](fg-yellow)  Acknowledged [⬤](fg-red)  Triggered
`
	helpWidget := ui.NewPar(helpWidgetText)
	helpWidget.Height = 10
	helpWidget.BorderFg = ui.ColorCyan
	helpWidget.TextFgColor = ui.ColorWhite
	helpWidget.BorderLabel = "Help center"

	return helpWidget
}

func modeWidget(ctx *AppContext) *ui.Gauge {
	modeWidget := ui.NewGauge()
	modeWidget.Percent = 100
	modeWidget.Label = ctx.Mode.Code
	modeWidget.BarColor = ctx.Mode.Color
	modeWidget.PercentColorHighlighted = ui.ColorDefault

	return modeWidget
}

func onCallStatusWidget() *ui.Gauge {
	onCallStatusWidget := ui.NewGauge()
	onCallStatusWidget.Percent = 100
	onCallStatusWidget.BarColor = ui.ColorDefault
	onCallStatusWidget.PercentColorHighlighted = ui.ColorDefault
	onCallStatusWidget.Label = "Currently off call"

	return onCallStatusWidget
}
func incidentsWidget() *ui.ListBox {
	incidentsListBox := ui.NewListBox()
	incidentsListBox.Height = 2
	incidentsListBox.ItemFgColor = ui.ColorWhite
	incidentsListBox.BorderLabel = "Incidents"

	return incidentsListBox
}

func incidentDetailsWidget() *ui.Table {
	incidentDetails := ui.NewTable()
	incidentDetails.Border = false
	return incidentDetails
}

func incidentDescriptionWidget() *ui.Par {
	incidentDescriptionWidget := ui.NewPar("")
	incidentDescriptionWidget.TextFgColor = ui.ColorWhite
	incidentDescriptionWidget.Border = false

	return incidentDescriptionWidget
}

func servicesWidget() *ui.ListBox {
	servicesListBox := ui.NewListBox()
	servicesListBox.Height = 2
	servicesListBox.ItemFgColor = ui.ColorYellow
	servicesListBox.BorderLabel = "Services"

	return servicesListBox
}

// Widgets ..
type Widgets struct {
	HelpWidget                *ui.Par
	ModeWidget                *ui.Gauge
	IncidentsWidget           *ui.ListBox
	ServicesWidget            *ui.ListBox
	OnCallStatusWidget        *ui.Gauge
	IncidentDetailsWidget     *ui.Table
	IncidentDescriptionWidget *ui.Par
}
