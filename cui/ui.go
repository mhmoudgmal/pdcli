package cui

import (
	. "pdcli/config"

	ui "github.com/pdevine/termui"
)

// InitUI ....
func InitUI(ctx *AppContext) {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

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
