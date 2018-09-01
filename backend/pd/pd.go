package pd

// Config the pagerduty configuration
type Config struct {
	Token string
	Email string
}

// Backend the PD backend
type Backend struct {
	Config
}

// GetConfig ..
func (be Backend) GetConfig() interface{} {
	return be.Config
}
