package main

import (
	"errors"
	"fmt"
	"os"
)

var customersSlice map[int]any = make(map[int]any)

func main() {
	defer fmt.Println("Ejecucion Finalizada")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Several errors were detected at runtime")
		}
	}()

	registerCustomers_3("File", "Name1", 1, 123, "Home1")
	registerCustomers_3("File", "Name1", 1, 123, "Home1")
}

func costumers_1() {
	data, err := os.ReadFile("customers.txt")

	defer fmt.Println("Ejecucion Finalizada")

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	fmt.Println(string(data))
}

func data_2() {
	data, err := os.ReadFile("customers_2.txt")

	defer fmt.Println("Ejecucion Finalizada")

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	fmt.Println(string(data))
}

func registerCustomers_3(file string, name string, id int, phone int, home string) {

	_, ok := customersSlice[id]
	if ok == true {
		panic("Error: client already exists")
	}

	_, err := validateNewCustomer(file, name, id, phone, home)
	if err != nil {
		fmt.Println(err)
	}

	customersSlice[id] = name
}

func validateNewCustomer(file string, name string, id int, phone int, home string) (bool, error) {

	if file == "" || name == "" || home == "" || id == 0 || phone == 0 {
		return false, errors.New("Invalid data")
	}

	return true, nil
}
