package cui

import (
	. "github.com/mhmoudgmal/pdcli/config"
	. "github.com/mhmoudgmal/pdcli/models"

	ui "github.com/pdevine/termui"
)

func handleEvents(ctx *AppContext, wdgts Widgets) {

	// send an "acknowledge" message to the PD updating channel
	ui.Handle("/sys/kbd/C-k", func(ui.Event) {
		*ctx.PDUpdatingChannel <- IncidentUpdateInfo{
			ID:     wdgts.incidentsWidget.Current().ItemVal,
			From:   ctx.PDConfig.Email,
			Status: ACKNOWLEDGED,
		}
	})

	// send a "resolve" message to the PD updating channel
	ui.Handle("/sys/kbd/C-r", func(ui.Event) {
		*ctx.PDUpdatingChannel <- IncidentUpdateInfo{
			ID:     wdgts.incidentsWidget.Current().ItemVal,
			From:   ctx.PDConfig.Email,
			Status: RESOLVED,
		}
	})

	// toggle Auto-Ack mode
	ui.Handle("/sys/kbd/C-t", func(ui.Event) {
		if ctx.Mode == &ModeM {
			ctx.Mode = &ModeA
		} else {
			ctx.Mode = &ModeM
		}

		wdgts.modeWidget.Label = ctx.Mode.Code
		wdgts.modeWidget.BarColor = ctx.Mode.Color
		wdgts.modeTextNoteWidget.Text = ctx.Mode.Note

		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		wdgts.incidentsWidget.Up()
		ui.Render(wdgts.incidentsWidget)
	})

	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		wdgts.incidentsWidget.Down()
		ui.Render(wdgts.incidentsWidget)
	})

	// Close app
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		*ctx.TermChannel <- true
		*ctx.StopFrequestingChannel <- true
		ui.StopLoop()
	})

}
