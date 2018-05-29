package cui

import (
	ui "github.com/pdevine/termui"

	"pdcli/config"
	. "pdcli/cui/helpers"
)

// InitUI Initializes the app CUI
func InitUI(ctx *config.AppContext) {
	defer ui.Close()

	if err := ui.Init(); err != nil {
		panic(err)
	}

	// create widgets
	widgets := Widgets{
		helpWidget(),
		modeWidget(ctx),
		incidentsWidget(),
		onCallStatusWidget(ctx),
		incidentDetailsWidget(),
	}

	// build cui
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 0, widgets.HelpWidget),
			ui.NewCol(6, 0, widgets.OnCallStatusWidget, widgets.ModeWidget),
		),
		ui.NewRow(
			ui.NewCol(6, 0, widgets.IncidentsWidget),
			ui.NewCol(6, 0, widgets.IncidentDetailsWidget),
		),
	)

	// calculate layout and render
	ui.Body.Align()
	ui.Render(ui.Body)

	go HandleEvents(ctx, widgets)
	go UpdateIncidentsWidget(ctx, widgets.IncidentsWidget)
	go UpdateIncidentDetailsWidget(ctx, widgets.IncidentDetailsWidget)

	ui.Loop()
}
