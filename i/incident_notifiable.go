package i

// IncidentNotifiable represents the downstream client,
// that gets notified with incidents.
// Supported notifiable:
//	- Cui: a command line interface
//
// Any other notifiable or client should implement this interface.
type IncidentNotifiable interface {
	// Every client / notifiable should clean after itself.
	Clean()
	// Initialze the client or the notifiable.
	Init(*AppContext)
	// Notify the client or the notifiable.
	Notify(string, interface{})
}
