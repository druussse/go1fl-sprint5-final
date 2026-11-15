package spentenergy

import (
	"fmt"
	"time"
)

// Коэффициенты и единицы измерения для расчета активности.
const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance возвращает пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

// MeanSpeed вычисляет среднюю скорость в километрах в час.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// RunningSpentCalories вычисляет калории, израсходованные при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные параметры бега")
	}
	speed := MeanSpeed(steps, height, duration)
	calories := weight * speed * duration.Minutes() / minInH
	return calories, nil
}

// WalkingSpentCalories вычисляет калории, израсходованные при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные параметры ходьбы")
	}
	speed := MeanSpeed(steps, height, duration)
	calories := weight * speed * duration.Minutes() / minInH
	return calories * walkingCaloriesCoefficient, nil
}
