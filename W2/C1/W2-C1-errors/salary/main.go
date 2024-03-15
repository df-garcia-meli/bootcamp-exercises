package main

import "fmt"

func main() {
	var salary int = 20000

	result, err := salaryExercise(salary)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func salaryExercise(salary int) (string, error) {

	if salary <= 150000 {
		return "", &SalaryError{Msg: "Error: the salary entered does not reach the taxable minimum"}
	}
	return "Must pay tax", nil
}

type SalaryError struct {
	Msg string
}

func (e *SalaryError) Error() string {
	return e.Msg
}
