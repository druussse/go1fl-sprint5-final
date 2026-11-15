package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные о дневной активности.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает запись вида "678,0h50m".
func (ds *DaySteps) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 2 {
		return fmt.Errorf("неверный формат прогулки: %q", datastring)
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("некорректное количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return fmt.Errorf("некорректная длительность прогулки: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность прогулки должна быть положительной")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует отчет о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)
	return info, nil
}
