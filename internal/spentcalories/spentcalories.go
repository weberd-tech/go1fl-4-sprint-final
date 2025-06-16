package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)


func parseTraining(data string) (int, string, time.Duration, error) {
	sliceData := strings.Split(data, ",")

	if len(sliceData) != 3 {
		return 0, "", 0, fmt.Errorf("slice is not equal to 3")
	}
	activeType := sliceData[1]
	stepUser, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, "", 0, errors.New("step conversion error")
	}
	if stepUser <= 0 {
		return 0, "", 0, fmt.Errorf("steps are 0")
	}
	durationUser, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return 0, "", 0, errors.New("time conversion error")
	}
	if durationUser <= 0 {
		return 0, "", 0, fmt.Errorf("steps are 0")
	}

	return stepUser, activeType, durationUser, nil
}

func distance(steps int, height float64) float64 {
	widthStep := height * stepLengthCoefficient

	distance := float64(steps) * widthStep

	resultDistance := distance / float64(mInKm)

	return resultDistance

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)
	meanSpeed := distance / duration.Hours()

	return meanSpeed

}
func TrainingInfo(data string, weight, height float64) (string, error) {
	stepParse, typeParse, durationParse, err := parseTraining(data)
	if err != nil {
		return "", errors.New("error getting parameters")
	}

	switch typeParse {
	case "Бег":
		distRun := distance(stepParse, height)
		speedRun := meanSpeed(stepParse, height, durationParse)
		caloriesRun, err := RunningSpentCalories(stepParse, weight, height, durationParse)
		if err != nil {
			return "", errors.New("calorie counting error")
		}
		hours := durationParse.Hours()
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeParse, hours, distRun, speedRun, caloriesRun), nil

	case "Ходьба":
		distWalk := distance(stepParse, height)
		speedWalk := meanSpeed(stepParse, height, durationParse)
		caloriesWalk, err := WalkingSpentCalories(stepParse, weight, height, durationParse)
		if err != nil {
			return "", errors.New("calorie counting error")
		}
		hours := durationParse.Hours()
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeParse, hours, distWalk, speedWalk, caloriesWalk), nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("negative number of steps or equal to 0")
	}
	if weight <= 0 {
		return 0, errors.New("negative weight value")
	}
	if height <= 0 {
		return 0, errors.New("growth must be a positive number")
	}
	if duration <= 0 {
		return 0, errors.New("growth must be a positive number")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	calories := (weight * averageSpeed * duration.Minutes()) / minInH

	return float64(calories), nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("the number of steps must be non-negative")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be a positive number")
	}
	if height <= 0 {
		return 0, errors.New("growth must be a positive number")
	}
	if duration <= 0 {
		return 0, errors.New("growth must be a positive number")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	calories := (weight * averageSpeed * duration.Minutes()) / minInH

	return float64(calories * walkingCaloriesCoefficient), nil
}