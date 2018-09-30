package cui

import (
	"os/exec"
	"runtime"

	ui "github.com/mhmoudgmal/termui"

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
// C-o --- opens incident in the browser
// C-c --- close app
func HandleEvents(ctx *AppContext, wdgts Widgets) {
	ui.Handle("/sys/kbd/C-k", func(ui.Event) {
		incidentID := wdgts.IncidentsWidget.Current().ItemVal
		Ack(incidentID, ctx.UpdateBackendChannel, ctx.Backend)
	})

	ui.Handle("/sys/kbd/C-r", func(ui.Event) {
		incidentID := wdgts.IncidentsWidget.Current().ItemVal
		Resolve(incidentID, ctx.UpdateBackendChannel, ctx.Backend)
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
		incidentID := wdgts.IncidentsWidget.Current().ItemVal
		*ctx.GetIncidentChannel <- incidentID
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
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-o", func(ui.Event) {
		incidentURL := wdgts.IncidentsWidget.Current().Data["url"]
		go open(incidentURL)
	})
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
