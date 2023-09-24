package cui

import (
	. "github.com/mhmoudgmal/pdcli/backend/pd"
	. "github.com/mhmoudgmal/pdcli/notifiable/cui/helpers"

	ui "github.com/mhmoudgmal/termui"
)

// Render ...
// As long as ui.Render is blocking, this is one way -
// to make it easy to test different evetns.
// @see.. tests/cui/event_handlers_test.go
var Render = ui.Render

// HandleEvents handls the kyboard commands.
func HandleEvents(
	wdgts Widgets,
	terminateChan *chan struct{},
	stopWorkerChan *chan struct{},
	inspectIncidentChan *chan string,
	updateIncidentStatusChan *chan struct {
		ID     string
		Status string
	},
) {
	ui.Handle("/sys/kbd/C-k", func(ui.Event) {
		incidentID := wdgts.IncidentsWidget.Current().ItemVal
		*updateIncidentStatusChan <- struct {
			ID     string
			Status string
		}{
			ID:     incidentID,
			Status: ACKNOWLEDGED,
		}
	})

	ui.Handle("/sys/kbd/C-r", func(ui.Event) {
		incidentID := wdgts.IncidentsWidget.Current().ItemVal
		*updateIncidentStatusChan <- struct {
			ID     string
			Status string
		}{
			ID:     incidentID,
			Status: RESOLVED,
		}
	})

	// ui.Handle("/sys/kbd/C-t", func(ui.Event) {
	// 	if ctx.Mode == ModeN {
	// 		ctx.Mode = ModeA
	// 	} else {
	// 		ctx.Mode = ModeN
	// 	}

	// 	wdgts.ModeWidget.Label = ctx.Mode.Code
	// 	wdgts.ModeWidget.BarColor = ctx.Mode.Color

	// 	Render(ui.Body)
	// })

	ui.Handle("/sys/kbd/C-v", func(ui.Event) {
		incidentID := wdgts.IncidentsWidget.Current().ItemVal
		*inspectIncidentChan <- incidentID
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

	ui.Handle("/sys/kbd/C-o", func(ui.Event) {
		incidentURL := wdgts.IncidentsWidget.Current().Data["url"]
		go OpenIncidentURL(incidentURL)
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		close(*terminateChan)
		close(*stopWorkerChan)
		ui.StopLoop()
	})
}
