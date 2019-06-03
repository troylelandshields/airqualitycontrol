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

const (
	moreInfoURL = ""
)

func (s Summary) Message() string {
	var msg string

	switch s.AirQualityLevel {
	case Good:
		msg = "Air quality is good today, enjoy a nice day in the sun."
	case Moderate:
		msg = "Scientists are forecasting moderately poor air quality today. Consider carpooling or taking public transit to work."
	case SomewhatUnhealthy:
		msg = "According to science, today's forecast predicts the air quality to be somewhat unhealthy. Consider carpooling or taking public transit to work."
	case Unhealthy:
		msg = "People with PHDs are predicting today's air quality to be unhealthy. Please work from home if possible."
	case VeryUnhealthy:
		msg = "People with PHDs are predicting today's air quality to be VERY unhealthy. Please work from home if possible."
	case Hazardous:
		msg = "The outlook on air quality is as bad as it gets today. Avoid driving and work at home. Stay inside as much as you can."
	default:
		msg = "Uh oh, air quality seems to be unknown."
	}

	if moreInfoURL != "" {
		msg = msg + " " + moreInfoURL
	}

	return msg
}

func (s Summary) ShouldSend() bool {
	return s.AirQualityLevel > Good
}
