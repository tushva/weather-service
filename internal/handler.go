package internal

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	model "github.com/tushva/weather-service/api"
)

func GetForecastedWeater(ctx context.Context, url string) (model.ForecastResponse, error) {
	todaysPeriod := model.Period{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		mylogger.Logger.Error("Error in creating request: ", err.Error(), nil)
		return model.ForecastResponse{}, err
	}
	req.Header.Set("Accept", "application/ld+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		mylogger.Logger.Error("Error request Weather URL: ", err.Error(), nil)
		return model.ForecastResponse{}, err
	}

	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	var weatherForecast model.WeatherForecast
	err = json.Unmarshal(responseBody, &weatherForecast)
	if err != nil {
		mylogger.Logger.Error("Unmarshal error: ", err.Error(), nil)
	}
	var charaterization string
	todaysPeriod, charaterization = GetTodaysForecast(weatherForecast.Periods, time.Now())
	forecastResponse := model.ForecastResponse{
		ShortForecast:    todaysPeriod.ShortForecast,
		Characterization: charaterization,
	}

	return forecastResponse, nil
}

// best to push this function into a service
func GetTodaysForecast(periods []model.Period, today time.Time) (model.Period, string) {
	//2025 August 15
	var todaysForecast []model.Period
	for _, period := range periods {
		layout := "2006-01-02T15:04:05-05:00"
		parsedTime, err := time.Parse(layout, period.StartTime)
		if err != nil {
			mylogger.Logger.Error("Error parsing time: ", err.Error(), nil)
			return model.Period{}, ""
		}
		if parsedTime.Truncate(24 * time.Hour).Equal(today.Truncate(24 * time.Hour)) {
			todaysForecast = append(todaysForecast, period)
		}
	}
	if len(todaysForecast) < 1 {
		mylogger.Logger.Error("Error comparing dates")
		return model.Period{}, ""
	}
	mylogger.Logger.Info("Found instances in forecase: ", strconv.Itoa(len(todaysForecast)), nil)
	characterization := TempCharacterization(periods[0].Temperature)
	return periods[0], characterization
}

func TempCharacterization(temperature int64) string {
	if temperature >= 90 {
		return model.Hot
	} else if temperature >= 65 {
		return model.Moderate
	}
	return model.Cold
}
