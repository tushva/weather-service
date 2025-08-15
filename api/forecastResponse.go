package api

type Period struct {
	Number           int64  `json:"number"`
	IsDaytime        bool   `json:"isDaytime"`
	Temperature      int64  `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	ShortForecast    string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
	WindSpeed        string `json:"windSpeed"`
	StartTime        string `json:"startTime"`
}

type WeatherForecast struct {
	UpdateTime  string   `json:"updateTime"`
	GeneratedAt string   `json:"generatedAt"`
	Periods     []Period `json:"periods"`
}

type ForecastResponse struct {
	ShortForecast    string `json:"shortForecast"`
	Characterization string `json:"characterization"`
}
