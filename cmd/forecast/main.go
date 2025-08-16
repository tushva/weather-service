package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tushva/weather-service/internal"
)

var configs *internal.Config

func main() {
	logger := internal.GetLoggerInstance()
	var err error
	configs, err = internal.LoadConfig()
	if err != nil {
		logger.Error("Exiting server, error reading configs, %s", err.Error(), nil)
		os.Exit(-1)
	}
	port := configs.Port
	if port == "" {
		logger.Warn("No port defined in environment, defaulting to 7800")
		port = "7800"
	}
	logger.Info("Defaulting to FWD weather station")
	http.HandleFunc("/forecast/{gridPoints}", forecastHandler) //expect comma separated grid points
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":"+port, nil)
}

func forecastHandler(resp http.ResponseWriter, req *http.Request) {
	logger := internal.GetLoggerInstance()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if req.Method == http.MethodPost {
		logger.Error("User initiated a POST method")
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("POST method not allowed"))
		return
	}
	gridPointsArr := strings.Split(req.PathValue("gridPoints"), ",") //comma separated to match inputs
	if len(gridPointsArr) < 2 || gridPointsArr[0] == "" || gridPointsArr[1] == "" {
		logger.Warn("Bad request, both x and y grid points as integer type needed")
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	x, err := strconv.Atoi(gridPointsArr[0])
	if err != nil {
		logger.Warn("Bad request, both x and y grid points should be integers")
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(gridPointsArr[1])
	if err != nil {
		logger.Error("Bad request, both x and y grid points should be integers")
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	weatherURL := fmt.Sprintf(configs.WeatherHost+configs.WeatherEndpoint+"%s,%s"+"/forecast", strconv.Itoa(x), strconv.Itoa(y))
	forecastResponse, err := internal.GetForecastedWeater(ctx, weatherURL)
	if err != nil {
		logger.Error("Error fetching response from weather service: %s", err.Error(), nil)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error fetching response from weather service"))
	}
	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(forecastResponse)
}

func healthHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("alive and kicking"))
}
