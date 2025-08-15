package internal

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tushva/weather-service/api"
)

func GetForecastedWeater(ctx context.Context, url string) (api.ForecastResponse, error) {
	todaysPeriod := api.Period{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		mylogger.Logger.Error("Error in creating request: ", err.Error(), nil)
		return api.ForecastResponse{}, err
	}
	req.Header.Set("Accept", "application/ld+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		mylogger.Logger.Error("Error request Weather URL: ", err.Error(), nil)
		return api.ForecastResponse{}, err
	}

	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	var weatherForecast api.WeatherForecast
	err = json.Unmarshal(responseBody, &weatherForecast)
	if err != nil {
		mylogger.Logger.Error("Unmarshal error: ", err.Error(), nil)
	}
	var charaterization string
	todaysPeriod, charaterization = GetTodaysForecast(weatherForecast.Periods, time.Now())
	forecastResponse := api.ForecastResponse{
		ShortForecast:    todaysPeriod.ShortForecast,
		Characterization: charaterization,
	}

	return forecastResponse, nil
}

// best to push this function into a service
func GetTodaysForecast(periods []api.Period, today time.Time) (api.Period, string) {
	//2025 August 15
	var todaysForecast []api.Period
	for _, period := range periods {
		layout := "2006-01-02T15:04:05-05:00"
		parsedTime, err := time.Parse(layout, period.StartTime)
		if err != nil {
			mylogger.Logger.Error("Error parsing time: ", err.Error(), nil)
			return api.Period{}, ""
		}
		if parsedTime.Truncate(24 * time.Hour).Equal(today.Truncate(24 * time.Hour)) {
			todaysForecast = append(todaysForecast, period)
		}
	}
	if len(todaysForecast) < 1 {
		mylogger.Logger.Error("Error comparing dates")
		return api.Period{}, ""
	}
	mylogger.Logger.Info("Found instances in forecase: ", strconv.Itoa(len(todaysForecast)), nil)
	characterization := tempCharacterization(periods[0].Temperature)
	return periods[0], characterization
}

func tempCharacterization(temperature int64) string {
	if temperature >= 90 {
		return "hot"
	} else if temperature >= 65 {
		return "moderate"
	}
	return "cold"
}
