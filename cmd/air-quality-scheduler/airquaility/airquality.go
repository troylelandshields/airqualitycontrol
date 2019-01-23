package airquaility

import "time"

type AirQualityLevel int

const (
	Unknown AirQualityLevel = iota
	Good
	Moderate
	SomewhatUnhealthy
	Unhealthy
	VeryUnhealthy
	Hazardous
)

type Summary struct {
	DateForecast    time.Time
	ActionDay       bool
	AirQualityLevel AirQualityLevel
}

func (s Summary) Message() string {
	switch s.AirQualityLevel {
	case Good:
		return "Air quality is good today, enjoy a nice day in the sun."
	case Moderate, SomewhatUnhealthy:
		return "The air quality today is fairly poor. Consider carpooling or taking public transit to work."
	case Unhealthy, VeryUnhealthy:
		return "The air quality today is very poor. Please work from home if possible."
	case Hazardous:
		return "Air quality is as bad as it gets today. Avoid driving and work at home. Stay inside as much as you can."
	default:
		return "Uh oh, air quality seems to be unknown."
	}
}

func (s Summary) ShouldSend() bool {
	return s.AirQualityLevel > Good
}
