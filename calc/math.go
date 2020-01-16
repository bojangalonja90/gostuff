package calc

import (
	"errors"

	"github.com/fatih/color"
)

//Add should return stuff
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		errorString := color.RedString("Provide more than two numbers")
		return errors.New(errorString), sum
	}

	for _, num := range numbers {
		sum = sum + num
	}

	return nil, sum
}
