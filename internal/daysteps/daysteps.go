package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	sliceData := strings.Split(data, ",")

	if len(sliceData) != 2 {
		return 0, 0, fmt.Errorf("slice is not equal to 0")
	}

	stepUser, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("step conversion error")
	}

	if stepUser <= 0 {
		return 0, 0, fmt.Errorf("steps value error")
	}

	durationUser, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("time conversion error '%s': %v", durationUser, err)
	}
	if durationUser <= 0 {
		return 0, 0, errors.New("duration is 0")
	}

	return stepUser, durationUser, nil

}

func DayActionInfo(data string, weight, height float64) string {
	stepUser, durationUser, err := parsePackage(data)
	if err != nil {
		log.Println("parsing error")
	}

	distanceUser := float64(stepUser) * stepLength
	distanceUserKm := distanceUser / float64(mInKm)
	if durationUser <= 0 {
		return ""
	}
	caloriesUser, err := spentcalories.WalkingSpentCalories(stepUser, weight, height, durationUser)
	if err != nil {
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepUser, distanceUserKm, caloriesUser)

	return result
}
