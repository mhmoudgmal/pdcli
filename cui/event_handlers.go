package cui

import (
	ui "github.com/pdevine/termui"

	. "pdcli/config"
	. "pdcli/models"
)

// HandleEvents handls the kyboard commands.
//
// C-k ---> Ack!
// C-r ---> Resolve!
// C-t ---> Toggle Auto-Ack mode
// C-c ---> Close app
func HandleEvents(ctx *AppContext, wdgts Widgets) {

	// send an "acknowledge" message to the PD updating channel
	ui.Handle("/sys/kbd/C-k", func(ui.Event) {
		*ctx.PDUpdatingChannel <- IncidentUpdateInfo{
			ID:     wdgts.IncidentsWidget.Current().ItemVal,
			From:   ctx.PDConfig.Email,
			Status: ACKNOWLEDGED,
		}
	})

	ui.Handle("/sys/kbd/C-r", func(ui.Event) {
		*ctx.PDUpdatingChannel <- IncidentUpdateInfo{
			ID:     wdgts.IncidentsWidget.Current().ItemVal,
			From:   ctx.PDConfig.Email,
			Status: RESOLVED,
		}
	})

	ui.Handle("/sys/kbd/C-t", func(ui.Event) {
		if ctx.Mode == &ModeM {
			ctx.Mode = &ModeA
		} else {
			ctx.Mode = &ModeM
		}

		wdgts.ModeWidget.Label = ctx.Mode.Code
		wdgts.ModeWidget.BarColor = ctx.Mode.Color
		wdgts.ModeTextNoteWidget.Text = ctx.Mode.Note

		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/kbd/k", func(ui.Event) {
		wdgts.IncidentsWidget.Up()
		ui.Render(wdgts.IncidentsWidget)
	})

	ui.Handle("/sys/kbd/j", func(ui.Event) {
		wdgts.IncidentsWidget.Down()
		ui.Render(wdgts.IncidentsWidget)
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		*ctx.TermChannel <- true
		*ctx.StopFrequestingChannel <- true
		ui.StopLoop()
	})

}
