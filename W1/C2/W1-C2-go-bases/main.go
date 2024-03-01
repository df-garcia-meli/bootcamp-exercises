package main

import "fmt"

func main() {
	employeesFunc()
}

// Variables and types
func names() {
	var firstName string = "Daniel"
	var lastName string = "Garcia"
	var age int = 24

	fmt.Println(firstName, lastName, age)
}

func weather() {
	var temperature, humidity, pression int = 30, 20, 1000

	fmt.Println(temperature, humidity, pression)
}

// Data structures
func countCharacters() {
	var word string = "Hello"

	fmt.Println(len(word))
	for _, character := range word {
		fmt.Println(string(character))
	}
}

func lending() {
	var age int = 30
	var employee bool = true
	var experience int = 2
	var salary int = 1000

	if age >= 22 {
		fmt.Println("You are old enough")
	}
	if employee == true {
		fmt.Println("You are an employee")
	}
	if experience > 1 {
		fmt.Println("You have experience")
	}

	if salary >= 100000 {
		fmt.Println("You have a high salary")
	} else {
		fmt.Println("You don't have a high salary")
	}
}

func months() {
	var months = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	var month int = 6

	if month >= 12 || month < 1 {
		fmt.Println("Invalid month")
	} else {
		fmt.Println(months[month-1])
	}
}

func employeesFunc() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])

	var counter int
	for _, age := range employees {
		if age >= 21 {
			counter++
		}
	}
	fmt.Println(counter)

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}

func sum() {
	fmt.Println(1 + 1)
}
