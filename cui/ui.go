package cui

import (
	ui "github.com/pdevine/termui"

	"pdcli/config"
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
		modeTextNoteWidget(ctx),
	}

	// build cui
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewRow(
				ui.NewCol(6, 0, widgets.HelpWidget),
				ui.NewCol(6, 0, ui.NewRow(
					ui.NewRow(
						ui.NewCol(12, 0, widgets.ModeWidget, widgets.ModeTextNoteWidget)))))),
		ui.NewRow(
			ui.NewCol(6, 0, widgets.IncidentsWidget)))

	// calculate layout and render
	ui.Body.Align()
	ui.Render(ui.Body)

	go HandleEvents(ctx, widgets)
	go updateIncidentsWidget(ctx, widgets.IncidentsWidget)

	ui.Loop()
}
