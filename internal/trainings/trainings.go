package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training описывает одну тренировку.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает запись вида "3456,Ходьба,3h00m".
func (t *Training) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 3 {
		return fmt.Errorf("неверный формат тренировки: %q", datastring)
	}

	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("некорректное количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(fields[2])
	if err != nil {
		return fmt.Errorf("некорректная длительность тренировки: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность тренировки должна быть положительной")
	}

	// Сохраняем поля только после успешной проверки всех значений.
	t.Steps = steps
	t.TrainingType = fields[1]
	t.Duration = duration
	return nil
}

// ActionInfo формирует отчет о тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories)
	return info, nil
}
