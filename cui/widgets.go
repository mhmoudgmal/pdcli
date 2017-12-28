package cui

import (
	"github.com/mhmoudgmal/pdcli/config"
	ui "github.com/pdevine/termui"
)

func helpWidget() *ui.Par {
	helpWidgetText := `
C-c     Quit
C-t     Toggle Auto-Ack mode
C-k     Acknowledge selected incident
C-r     Resolve selected incident

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
	modeWidget.Width = 50
	modeWidget.Y = 11
	modeWidget.BorderFg = ui.ColorRed
	modeWidget.BorderLabelFg = ui.ColorCyan
	modeWidget.BorderLabel = "Mode"

	return modeWidget
}

func modeTextNoteWidget(cfg *config.AppContext) *ui.Par {
	modeTextNoteWidget := ui.NewPar(cfg.Mode.Note)
	modeTextNoteWidget.Border = false
	modeTextNoteWidget.Height = 4
	modeTextNoteWidget.TextFgColor = ui.ColorWhite

	return modeTextNoteWidget
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

type Widgets struct {
	helpWidget         *ui.Par
	modeWidget         *ui.Gauge
	incidentsWidget    *ui.ListBox
	modeTextNoteWidget *ui.Par
}
