package helpers

import (
	ui "github.com/mhmoudgmal/termui"

	. "pdcli/i"
)

// UpdateIncidentsWidget ...
func UpdateServicesWidget(ctx *AppContext, lb *ui.ListBox) {
	servicesItems := []ui.Item{}

	select {
	case services := <-*ctx.ServicesChannel:
		servicesItems = append(servicesItems, mapServicesToUIItems(services)...)
		updateServicesListBox(lb, servicesItems)
	}
}

func mapServicesToUIItems(services IServices) []ui.Item {
	servicesItems := []ui.Item{ui.Item{}}

	for _, service := range services {
		servicesItems = append(
			servicesItems, ui.Item{
				ItemVal: service.GetID(),
				Text:    service.GetName(),
			}, ui.Item{},
		)
	}

	return servicesItems
}

func updateServicesListBox(lb *ui.ListBox, servicesItems []ui.Item) {
	lb.Items = servicesItems
	lb.Height = lb.Height + len(servicesItems)
	ui.Body.Align()
	ui.Render(lb)
	ui.Render(ui.Body)
}
