package calc

import "errors"

//Add should return stuff
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("Provide more than two numbers"), sum
	}

	for _, num := range numbers {
		sum = sum + num
	}

	return nil, sum
}
