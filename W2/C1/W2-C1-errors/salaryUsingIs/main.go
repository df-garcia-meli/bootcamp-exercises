package main

import (
	"errors"
	"fmt"
)

var ErrLowSalaryVar error = &CustomError{"Error: salary is less than 10000"}

func main() {
	var salary int = 10000

	err := salaryExercise(salary)

	if errors.Is(err, ErrLowSalaryVar) {
		fmt.Println(err)
	}
}

func salaryExercise(salary int) error {

	if salary <= 10000 {
		return ErrLowSalaryVar
	}

	return nil
}

type CustomError struct {
	Msg string
}

func (e *CustomError) Error() string {
	return e.Msg
}
