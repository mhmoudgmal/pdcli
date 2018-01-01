package cui_test

import (
	"github.com/mhmoudgmal/pdcli/config"
	cui "github.com/mhmoudgmal/pdcli/cui"
	. "github.com/mhmoudgmal/pdcli/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ui "github.com/pdevine/termui"
)

func test_command(ctx *config.AppContext, commandOpts map[string]string, expect func()) {
	Describe(commandOpts["cmd"], func() {
		incidentsWdgtMock := &ui.ListBox{
			Items: []ui.Item{ui.Item{
				ItemVal: "Item1",
				Text:    "Item1 Text",
			}},
		}
		incidentsWdgtMock.Selected = 0

		wdgts := cui.Widgets{nil, nil, incidentsWdgtMock, nil}

		It(commandOpts["msg"], func() {
			cui.HandleEvents(ctx, wdgts)

			go ui.DefaultEvtStream.Handlers[commandOpts["path"]](ui.Event{})
			expect()

		})
	})
}

var _ = Describe("Events", func() {
	var ctx config.AppContext
	var pdcfg config.PDConfig

	BeforeSuite(func() {
		updatingChan := make(chan IncidentUpdateInfo)
		termChan := make(chan bool)
		stopFreqChan := make(chan bool)

		pdcfg = config.PDConfig{Token: "token", Email: "foo@bar.baz"}
		ctx = config.AppContext{
			PDUpdatingChannel:      &updatingChan,
			PDConfig:               &pdcfg,
			TermChannel:            &termChan,
			StopFrequestingChannel: &stopFreqChan,
		}
	})

	test_command(&ctx, map[string]string{"cmd": "C-k", "path": "/sys/kbd/C-k", "msg": "sends ACKNOWLEDGED message to PDUpdating chan"}, func() {
		Eventually(*ctx.PDUpdatingChannel).Should(Receive(Equal(IncidentUpdateInfo{
			ID:     "Item1",
			From:   pdcfg.Email,
			Status: ACKNOWLEDGED,
		})))
	})

	test_command(&ctx, map[string]string{"cmd": "C-r", "path": "/sys/kbd/C-r", "msg": "sends RESOLVED message to PDUpdating chan"}, func() {
		Eventually(*ctx.PDUpdatingChannel).Should(Receive(Equal(IncidentUpdateInfo{
			ID:     "Item1",
			From:   pdcfg.Email,
			Status: RESOLVED,
		})))
	})

	test_command(&ctx, map[string]string{"cmd": "C-c", "path": "/sys/kbd/C-c", "msg": "sends TERM & STOPFREQUESTING messages"}, func() {
		Eventually(*ctx.TermChannel).Should(Receive(Equal(true)))
		Eventually(*ctx.StopFrequestingChannel).Should(Receive(Equal(true)))
	})
})
