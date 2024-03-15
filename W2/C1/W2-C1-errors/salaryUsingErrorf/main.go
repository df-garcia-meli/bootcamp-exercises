package main

import (
	"errors"
	"fmt"
)

var ErrLowSalaryVar error = errors.New("Error: the minimum taxable amount is 150,000 and the salary entered is:")

func main() {
	var salary int = 10000

	err := salaryExercise(salary)

	if errors.Is(err, ErrLowSalaryVar) {
		fmt.Println(err)
	}
}

func salaryExercise(salary int) error {

	if salary <= 10000 {
		return fmt.Errorf("%w %d", ErrLowSalaryVar, salary)
	}

	return nil
}
