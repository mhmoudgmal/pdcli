package cui

import (
	"github.com/mhmoudgmal/pdcli/config"
	ui "github.com/pdevine/termui"
)

// InitUI ....
func InitUI(ctx *config.AppContext) {
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
				ui.NewCol(6, 0, widgets.helpWidget),
				ui.NewCol(6, 0, ui.NewRow(
					ui.NewRow(
						ui.NewCol(12, 0, widgets.modeWidget, widgets.modeTextNoteWidget)))))),
		ui.NewRow(
			ui.NewCol(6, 0, widgets.incidentsWidget)))

	// calculate layout and render
	ui.Body.Align()
	ui.Render(ui.Body)

	go handleEvents(ctx, widgets)
	go updateIncidentsWidget(ctx, widgets.incidentsWidget)

	ui.Loop()
}
