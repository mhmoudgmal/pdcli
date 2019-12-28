package cui_test

import (
	ui "github.com/pdevine/termui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pdcli/backend/pd"

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

var terminateChan chan struct{}
var stopWorkerChan chan struct{}
var inspectIncidentChan chan string
var updateIncidentStatusChan chan struct {
	ID     string
	Status string
}

func testCommand(cmdOpts map[string]interface{}, expect func()) {
	Describe(cmdOpts["cmd"].(string), func() {
		It(cmdOpts["msg"].(string), func() {
			wdgts.IncidentsWidget.Selected = 0

			cui.HandleEvents(
				wdgts,
				&terminateChan,
				&stopWorkerChan,
				&inspectIncidentChan,
				&updateIncidentStatusChan,
			)

			evtHandlers := ui.DefaultEvtStream.Handlers
			handler := evtHandlers[cmdOpts["path"].(string)]

			if cmdOpts["before"] != nil {
				cmdOpts["before"].(func())()
			}

			if cmdOpts["exec"].(string) == "async" {
				go handler(ui.Event{})
			} else {
				handler(ui.Event{})
			}

			expect()
		})
	})
}

var _ = Describe("Events", func() {
	terminateChan = make(chan struct{})
	stopWorkerChan = make(chan struct{})
	inspectIncidentChan = make(chan string)
	updateIncidentStatusChan = make(chan struct {
		ID     string
		Status string
	})

	BeforeSuite(func() {
		mockUI()
	})

	AfterSuite(func() {
		resetUI()
	})

	type cmdOpts map[string]interface{}

	testCommand(cmdOpts{
		"cmd":  "C-k",
		"path": "/sys/kbd/C-k",
		"exec": "async",
		"msg":  "sends ACKNOWLEDGED message to UpdateIncidentStatusChan",
	}, func() {
		Eventually(updateIncidentStatusChan).
			Should(Receive(Equal(struct {
				ID     string
				Status string
			}{
				ID:     "Item1",
				Status: ACKNOWLEDGED,
			})))
	})

	testCommand(cmdOpts{
		"cmd":  "C-r",
		"path": "/sys/kbd/C-r",
		"exec": "async",
		"msg":  "sends RESOLVED message to UpdateIncidentStatusChan",
	}, func() {
		Eventually(updateIncidentStatusChan).
			Should(Receive(Equal(struct {
				ID     string
				Status string
			}{
				ID:     "Item1",
				Status: RESOLVED,
			})))
	})

	testCommand(cmdOpts{
		"cmd":  "C-v",
		"path": "/sys/kbd/C-v",
		"exec": "async",
		"msg":  "sends message to InspectIncidentChan",
	}, func() {
		Eventually(inspectIncidentChan).Should(Receive(Equal("Item1")))
	})

	testCommand(cmdOpts{
		"cmd":    "k",
		"path":   "/sys/kbd/k",
		"msg":    "navigate up and select the previous incident item",
		"exec":   "sync",
		"before": func() { wdgts.IncidentsWidget.Selected = 2 },
	}, func() {
		Expect(wdgts.IncidentsWidget.Selected).Should(Equal(1))
	})

	testCommand(cmdOpts{
		"cmd":    "j",
		"path":   "/sys/kbd/j",
		"msg":    "navigate down and select the next incident item",
		"exec":   "sync",
		"before": func() { wdgts.IncidentsWidget.Selected = 1 },
	}, func() {
		Expect(wdgts.IncidentsWidget.Selected).Should(Equal(2))
	})

	testCommand(cmdOpts{
		"cmd":  "C-c",
		"path": "/sys/kbd/C-c",
		"msg":  "sends TERM & stopWorkerChan messages",
		"exec": "async",
	}, func() {
		Eventually(terminateChan).Should(BeClosed())
		Eventually(stopWorkerChan).Should(BeClosed())
	})
})
