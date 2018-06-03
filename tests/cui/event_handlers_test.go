package cui_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	ui "github.com/pdevine/termui"

	"pdcli/config"
	"pdcli/cui"
	"pdcli/models"
)

var incidentsWdgtMock = ui.ListBox{
	Items: []ui.Item{
		ui.Item{
			ItemVal: "Item1",
			Text:    "Item1 Text",
		},
		ui.Item{
			ItemVal: "Item2",
			Text:    "Item2 Text",
		},
		ui.Item{
			ItemVal: "Item3",
			Text:    "Item3 Text",
		},
	},
}

var wdgts = cui.Widgets{
	IncidentsWidget: &incidentsWdgtMock,
}

func mockUI() {
	cui.Render = func(...ui.Bufferer) {}
}

func resetUI() {
	cui.Render = ui.Render
}

func testCommand(ctx *config.AppContext, commandOpts map[string]interface{}, expect func()) {
	Describe(commandOpts["cmd"].(string), func() {
		It(commandOpts["msg"].(string), func() {
			wdgts.IncidentsWidget.Selected = 0

			cui.HandleEvents(ctx, wdgts)

			evtHandlers := ui.DefaultEvtStream.Handlers
			handler := evtHandlers[commandOpts["path"].(string)]

			if commandOpts["before"] != nil {
				commandOpts["before"].(func())()
			}

			if commandOpts["exec"].(string) == "async" {
				go handler(ui.Event{})
			} else {
				handler(ui.Event{})
			}

			expect()
		})
	})
}

var _ = Describe("Events", func() {
	termChan := make(chan bool)
	pdGetIChan := make(chan string)
	stopFreqChan := make(chan bool)
	updatingChan := make(chan models.IncidentUpdateInfo)

	pdcfg := config.PDConfig{
		Token: "token",
		Email: "foo@bar.baz",
	}
	ctx := config.AppContext{
		PDUpdatingChannel:      &updatingChan,
		PDConfig:               &pdcfg,
		TermChannel:            &termChan,
		StopFrequestingChannel: &stopFreqChan,
		PDGetIncidentChannel:   &pdGetIChan,
	}

	BeforeSuite(func() {
		mockUI()
	})

	AfterSuite(func() {
		resetUI()
	})

	type commandOpts map[string]interface{}

	testCommand(&ctx, commandOpts{
		"cmd":  "C-k",
		"path": "/sys/kbd/C-k",
		"exec": "async",
		"msg":  "sends ACKNOWLEDGED message to PDUpdating chan",
	}, func() {
		Eventually(*ctx.PDUpdatingChannel).
			Should(Receive(Equal(models.IncidentUpdateInfo{
				ID:     "Item1",
				From:   pdcfg.Email,
				Status: models.ACKNOWLEDGED,
			})))
	})

	testCommand(&ctx, commandOpts{
		"cmd":  "C-r",
		"path": "/sys/kbd/C-r",
		"exec": "async",
		"msg":  "sends RESOLVED message to PDUpdating chan",
	}, func() {
		Eventually(*ctx.PDUpdatingChannel).
			Should(Receive(Equal(models.IncidentUpdateInfo{
				ID:     "Item1",
				From:   pdcfg.Email,
				Status: models.RESOLVED,
			})))
	})

	testCommand(&ctx, commandOpts{
		"cmd":  "C-v",
		"path": "/sys/kbd/C-v",
		"exec": "async",
		"msg":  "sends message to PDGetIncident chan",
	}, func() {
		Eventually(*ctx.PDGetIncidentChannel).Should(Receive(Equal("Item1")))
	})

	testCommand(&ctx, commandOpts{
		"cmd":    "k",
		"path":   "/sys/kbd/k",
		"msg":    "navigate up and select the previous incident item",
		"exec":   "sync",
		"before": func() { wdgts.IncidentsWidget.Selected = 2 },
	}, func() {
		Expect(wdgts.IncidentsWidget.Selected).Should(Equal(1))
	})

	testCommand(&ctx, commandOpts{
		"cmd":    "j",
		"path":   "/sys/kbd/j",
		"msg":    "navigate down and select the next incident item",
		"exec":   "sync",
		"before": func() { wdgts.IncidentsWidget.Selected = 1 },
	}, func() {
		Expect(wdgts.IncidentsWidget.Selected).Should(Equal(2))
	})

	testCommand(&ctx, commandOpts{
		"cmd":  "C-c",
		"path": "/sys/kbd/C-c",
		"msg":  "sends TERM & STOPFREQUESTING messages",
		"exec": "async",
	}, func() {
		Eventually(*ctx.TermChannel).Should(Receive(Equal(true)))
		Eventually(*ctx.StopFrequestingChannel).Should(Receive(Equal(true)))
	})
})
