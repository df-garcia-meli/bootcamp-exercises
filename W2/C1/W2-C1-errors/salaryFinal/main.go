package main

import (
	"errors"
	"fmt"
)

func main() {
	salary, err := salaryByHours(80, 1.0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(salary)
}

func salaryByHours(totalHours int, baseSalary float64) (float64, error) {
	salary := float64(totalHours) * baseSalary

	if salary >= 150000 {
		salary = salary * 0.9
	}

	if totalHours < 80 {
		return 0.0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	return salary, nil
}
