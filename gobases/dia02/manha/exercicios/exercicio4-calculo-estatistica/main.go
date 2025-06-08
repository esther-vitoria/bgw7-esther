package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minValue(numbers ...float64) float64 {
	var minValue = numbers[0]
	for _, number := range numbers {
		minValue = math.Min(minValue, number)
	}

	return minValue
}

func maxValue(numbers ...float64) float64 {
	var maxValue float64
	for _, number := range numbers {
		maxValue = math.Max(maxValue, number)
	}

	return maxValue
}

func avgValue(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}

func operation(operation string) (func(numbers ...float64) float64, error) {
	switch operation {
	case minimum:
		return minValue, nil
	case maximum:
		return maxValue, nil
	case average:
		return avgValue, nil
	default:
		return nil, errors.New("operacao invalida")
	}

}

func main() {

	minFunc, err := operation(minimum)
	if err != nil {
		panic(err)
	}
	avgFunc, err := operation(average)
	if err != nil {
		panic(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		panic(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	avgValue := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("O valor mínimo é %.0f: ", minValue)
	fmt.Printf("O valor máximo é %.0f: ", maxValue)
	fmt.Printf("A média é %.0f: ", avgValue)
}
