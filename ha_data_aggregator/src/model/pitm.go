package model

type TemperatureModel struct {
	MachineName string  `json:"machineName"`
	Temperature float64 `json:"temperature"`
	// Datetime    time.Time `json:"datetime"`
	// Timestamp   uint64    `json:"timestamp"`
}
