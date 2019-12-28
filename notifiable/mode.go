package notifiable

import ui "github.com/mhmoudgmal/termui"

// Mode the application modes.
type Mode struct {
	Code  string
	Color ui.Attribute
}

var (
	// ModeA - auto-ack mode
	ModeA = Mode{"Auto-Ack", ui.ColorYellow}

	// ModeN the normal mode
	ModeN = Mode{"Normal", ui.ColorUndef}
)
