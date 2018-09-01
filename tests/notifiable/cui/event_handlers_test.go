package cui_test

import (
	ui "github.com/pdevine/termui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pdcli/i"

	"pdcli/backend/pd"
	"pdcli/notifiable/cui"
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

func testCommand(ctx *AppContext, commandOpts map[string]interface{}, expect func()) {
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
	updatingChan := make(chan UpdateIncidentInfo)

	pdbe := pd.Backend{
		pd.Config{
			Token: "token",
			Email: "foo@bar.baz",
		},
	}

	ctx := AppContext{
		Backend:                &pdbe,
		UpdateBackendChannel:   &updatingChan,
		TerminateChannel:       &termChan,
		StopFrequestingChannel: &stopFreqChan,
		GetIncidentChannel:     &pdGetIChan,
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
		"msg":  "sends ACKNOWLEDGED message to UpdateBackendChannel",
	}, func() {
		Eventually(*ctx.UpdateBackendChannel).
			Should(Receive(Equal(UpdateIncidentInfo{
				ID:     "Item1",
				Status: ACKNOWLEDGED,
				Config: ctx.Backend,
			})))
	})

	testCommand(&ctx, commandOpts{
		"cmd":  "C-r",
		"path": "/sys/kbd/C-r",
		"exec": "async",
		"msg":  "sends RESOLVED message to UpdateBackendChannel",
	}, func() {
		Eventually(*ctx.UpdateBackendChannel).
			Should(Receive(Equal(UpdateIncidentInfo{
				ID:     "Item1",
				Status: RESOLVED,
				Config: ctx.Backend,
			})))
	})

	testCommand(&ctx, commandOpts{
		"cmd":  "C-v",
		"path": "/sys/kbd/C-v",
		"exec": "async",
		"msg":  "sends message to GetIncidentChannel",
	}, func() {
		Eventually(*ctx.GetIncidentChannel).Should(Receive(Equal("Item1")))
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
		Eventually(*ctx.TerminateChannel).Should(Receive(Equal(true)))
		Eventually(*ctx.StopFrequestingChannel).Should(Receive(Equal(true)))
	})
})
