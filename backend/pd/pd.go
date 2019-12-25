package pd

import "time"

// Config the pagerduty configuration
type Config struct {
	Token string
	Email string
}

// Backend the PD backend
type Backend struct {
	Config
}

func (be Backend) Worker(reqInterval time.Duration, notify func(...Incident)) {
	params := map[string]string{"since": ""}

	for range time.Tick(reqInterval) {
		incidents, _ := be.GetIncidents(params)

		t, _ := time.Now().MarshalText()
		params["since"] = string(t)

		if len(incidents) == 0 {
			continue
		}

		notify(incidents...)
	}
}
