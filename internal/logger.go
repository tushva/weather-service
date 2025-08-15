package internal

import (
	"log/slog"
	"os"
	"sync"
)

type logger struct {
	*slog.Logger
}

func (l logger) GetLoggerInstance() any {
	panic("unimplemented")
}

var mylogger *logger
var once sync.Once

func GetLoggerInstance() *logger {
	once.Do(func() {
		mylogger = createLogger()
	})
	return mylogger
}

func createLogger() *logger {
	return &logger{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}
