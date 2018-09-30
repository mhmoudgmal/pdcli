package cui

import (
	ui "github.com/mhmoudgmal/termui"

	. "pdcli/i"
	. "pdcli/notifiable/cui/helpers"
)

// Cui the command line interface notifiable or client.
type Cui struct {
	*AppContext
	Widgets
}

// Init Initializes the app CUI
func (Cui) Init(ctx *AppContext) {
	defer ui.Close()

	if err := ui.Init(); err != nil {
		panic(err)
	}

	// create widgets
	widgets := Widgets{
		helpWidget(),
		modeWidget(ctx),
		incidentsWidget(),
		servicesWidget(),
		onCallStatusWidget(),
		incidentDetailsWidget(),
		incidentDescriptionWidget(),
	}

	// build ui
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 0, widgets.HelpWidget),
			ui.NewCol(6, 0, widgets.OnCallStatusWidget, widgets.ModeWidget),
		),

		ui.NewRow(
			ui.NewCol(6, 0, widgets.IncidentDetailsWidget, widgets.IncidentDescriptionWidget),
			ui.NewCol(6, 0, widgets.IncidentsWidget, widgets.ServicesWidget),
		),
	)

	// calculate layout and render
	ui.Body.Align()
	ui.Render(ui.Body)

	// Handle events and update widgets
	go HandleEvents(ctx, widgets)
	go UpdateIncidentsWidget(ctx, widgets.IncidentsWidget)
	go UpdateServicesWidget(ctx, widgets.ServicesWidget)
	go UpdateIncidentDetailsWidget(ctx,
		widgets.IncidentDetailsWidget,
		widgets.IncidentDescriptionWidget,
	)

	ui.Loop()
}

// Notify Cui
func (c Cui) Notify(msg string, data interface{}) {
	switch msg {
	case "new-incidents":
		*c.AppContext.IncidentsChannel <- data.(IIncidents)
	case "updated-incident":
		*c.AppContext.UpdateStatusChannel <- data.(IIncident)
	case "detailed-incident":
		*c.AppContext.IncidentDetailsChannel <- data.(IIncident)
	case "list-services":
		*c.AppContext.ServicesChannel <- data.(IServices)
	}
}

// Clean after you
func (Cui) Clean() {
	ui.Clear()
}
