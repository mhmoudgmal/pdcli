package cui

import (
	ui "github.com/mhmoudgmal/termui"

	. "github.com/mhmoudgmal/pdcli/backend/pd"
	. "github.com/mhmoudgmal/pdcli/notifiable"
	. "github.com/mhmoudgmal/pdcli/notifiable/cui/helpers"
)

// Cui the command line interface notifiable client.
type Cui struct {
	Widgets
}

// Init initializes command line interface client
func (Cui) Init(
	mode *Mode,
	terminateChan *chan struct{},
	stopWorkerChan *chan struct{},
	inspectIncidentChan *chan string,

	incidentChan *chan Incident,
	incidentsChan *chan []Incident,

	updateIncidentStatusChan *chan struct {
		ID     string
		Status string
	},
) {
	defer ui.Close()

	if err := ui.Init(); err != nil {
		panic(err)
	}

	// create widgets
	widgets := Widgets{
		helpWidget(),
		modeWidget(mode),
		incidentsWidget(),
		onCallStatusWidget(),
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

	go HandleEvents(widgets, terminateChan, stopWorkerChan, inspectIncidentChan, updateIncidentStatusChan)
	go UpdateIncidentsWidget(widgets.IncidentsWidget, incidentsChan)
	go InspectIncidentWidget(incidentChan, widgets.IncidentDetailsWidget)

	ui.Loop()
}

// Notify notifies Cui
func (c Cui) Notify(cb func(...Incident), data ...Incident) {
	cb(data...)
}

// Clean cleans Cui
func (Cui) Clean() {
	ui.Clear()
}
