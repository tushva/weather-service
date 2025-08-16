package test

import (
	"testing"

	model "github.com/tushva/weather-service/api"
	"github.com/tushva/weather-service/internal"
)

func TestTempCharacterization(t *testing.T) {
	var temperature int64 = 75
	messageModerate := internal.TempCharacterization(temperature)
	if messageModerate == "" || messageModerate != model.Moderate {
		t.Errorf(`Temp %d Characterization has a missed condition, expected %s, received %s`, temperature, model.Moderate, messageModerate)
	}
	temperature = 30
	messageCold := internal.TempCharacterization(temperature)
	if messageCold == "" || messageCold != model.Cold {
		t.Errorf(`Temp %d Characterization has a missed condition, expected %s, received %s`, temperature, model.Cold, messageCold)
	}
	temperature = 110
	messageHot := internal.TempCharacterization(temperature)
	if messageHot == "" || messageHot != model.Hot {
		t.Errorf(`Temp %d Characterization has a missed condition, expected %s, received %s`, temperature, model.Hot, messageHot)
	}
}
