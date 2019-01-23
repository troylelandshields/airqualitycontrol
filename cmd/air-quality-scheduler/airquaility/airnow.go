package airquaility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	airNowBaseURL = "http://www.airnowapi.org/aq/forecast/zipCode/?format=application/json&distance=25"

	dateFormat = "2006-01-02"
)

type airNowCategory struct {
	Number AirQualityLevel
	Name   string
}

type airNowQuality struct {
	DateForecast  string
	ReportingArea string
	StateCode     string
	Category      airNowCategory
	ActionDay     bool
}

func AirQuality(zipCode string, apiKey string, t time.Time) (Summary, error) {

	url, err := buildURL(airNowBaseURL, zipCode, t, apiKey)
	if err != nil {
		return Summary{}, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return Summary{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return Summary{}, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	var data []airNowQuality
	err = dec.Decode(&data)
	if err != nil {
		return Summary{}, err
	}
	if len(data) < 1 {
		return Summary{}, fmt.Errorf("no data returned")
	}

	forecast := data[0]
	forecastDate, err := time.Parse(dateFormat, strings.TrimSpace(forecast.DateForecast))
	if err != nil {
		return Summary{}, fmt.Errorf("invalid date format")
	}

	return Summary{
		DateForecast:    forecastDate,
		ActionDay:       forecast.ActionDay,
		AirQualityLevel: forecast.Category.Number,
	}, nil
}

// buildURL builds a URL like this:
// "http://www.airnowapi.org/aq/forecast/zipCode/?format=application/json&zipCode=84094&date=2019-01-22&distance=25&API_KEY=the_key"
func buildURL(baseURL string, zipCode string, date time.Time, apiKey string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	val := u.Query()

	val.Add("zipCode", zipCode)
	val.Add("API_KEY", apiKey)
	val.Add("date", date.Format(dateFormat))

	u.RawQuery = val.Encode()

	return u.String(), nil
}
