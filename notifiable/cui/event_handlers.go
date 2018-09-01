package cui

import (
	ui "github.com/pdevine/termui"

	. "pdcli/i"
)

// Render ...
// As long as ui.Render is blocking, this is one way -
// to make it easy to test different evetns.
// @see.. tests/cui/event_handlers_test.go
var Render = ui.Render

// HandleEvents handls the kyboard commands.
//
// j   --- navigate down
// k   --- navigate up
// C-k --- acknowledge incident
// C-r --- resolve incident
// C-t --- toggle auto-ack mode
// C-c --- close app
func HandleEvents(ctx *AppContext, wdgts Widgets) {
	ui.Handle("/sys/kbd/C-k", func(ui.Event) {
		*ctx.UpdateBackendChannel <- UpdateIncidentInfo{
			ID:     wdgts.IncidentsWidget.Current().ItemVal,
			Status: ACKNOWLEDGED,
			Config: ctx.Backend,
		}
	})

	ui.Handle("/sys/kbd/C-r", func(ui.Event) {
		*ctx.UpdateBackendChannel <- UpdateIncidentInfo{
			ID:     wdgts.IncidentsWidget.Current().ItemVal,
			Status: RESOLVED,
			Config: ctx.Backend,
		}
	})

	ui.Handle("/sys/kbd/C-t", func(ui.Event) {
		if ctx.Mode == ModeN {
			ctx.Mode = ModeA
		} else {
			ctx.Mode = ModeN
		}

		wdgts.ModeWidget.Label = ctx.Mode.Code
		wdgts.ModeWidget.BarColor = ctx.Mode.Color

		Render(ui.Body)
	})

	ui.Handle("/sys/kbd/C-v", func(ui.Event) {
		*ctx.GetIncidentChannel <- wdgts.IncidentsWidget.Current().ItemVal
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		Render(ui.Body)
	})

	ui.Handle("/sys/kbd/k", func(ui.Event) {
		wdgts.IncidentsWidget.Up()
		Render(wdgts.IncidentsWidget)
	})

	ui.Handle("/sys/kbd/j", func(ui.Event) {
		wdgts.IncidentsWidget.Down()
		Render(wdgts.IncidentsWidget)
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		*ctx.TerminateChannel <- true
		*ctx.StopFrequestingChannel <- true
		ui.StopLoop()
	})
}
