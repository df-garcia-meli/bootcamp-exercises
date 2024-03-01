package exercise

import "fmt"

type Category string

var (
	CategoryA Category = "A"
	CategoryB Category = "B"
	CategoryC Category = "C"
)

func main() {
	fmt.Println(multipleAnimals())
}

func salaryDeduction(salary float32) (tax float32) {
	if salary >= 50000 {
		tax = 0.17
	}

	if salary >= 150000 {
		tax += 0.1
	}

	return
}

func ellipsisAverage(numbers ...float32) float32 {
	var totalSum float32
	for _, number := range numbers {
		totalSum += number
	}
	return totalSum / float32(len(numbers))
}

func salary(totalMinutes int, category Category) (totalSalary float64) {
	var totalHours float64 = float64(totalMinutes) / 60
	switch category {
	case CategoryA:
		totalSalary = totalHours * 1000
	case CategoryB:
		totalSalary = (totalHours * 1500) * 1.2
	case CategoryC:
		totalSalary = (totalHours * 3000) * 1.5
	}

	return
}

const (
	minimum = "minimum"

	average = "average"

	maximum = "maximum"
)

func minFunc(numbers ...float32) float32 {
	var minValue float32 = numbers[0]
	for _, number := range numbers {
		if number < minValue {
			minValue = number
		}
	}
	return minValue
}

func maxFunc(numbers ...float32) float32 {
	var maxValue float32
	for _, number := range numbers {
		if number > maxValue {
			maxValue = number
		}
	}
	return maxValue
}

func averageFunc(numbers ...float32) float32 {
	var totalSum float32
	for _, number := range numbers {
		totalSum += number
	}
	return totalSum / float32(len(numbers))
}

func operation(operation string) (func(numbers ...float32) float32, error) {
	switch operation {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return nil, fmt.Errorf("operation %s not found", operation)
	}
}

func multipleFuncs() (float32, float32, float32) {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
	}

	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	return minValue, averageValue, maxValue
}

const (
	dog = "dog"

	cat = "cat"

	hamster   = "hamster"
	tarantula = "tarantula"
)

func animalDog(totalAnimals int) float32 {
	return float32(totalAnimals) * 10
}
func animalCat(totalAnimals int) float32 {
	return float32(totalAnimals) * 5
}
func animalHamster(totalAnimals int) float32 {
	return float32(totalAnimals) * 0.25
}
func animalTarantula(totalAnimals int) float32 {
	return float32(totalAnimals) * 0.15
}

func animal(animalText string) (func(totalAnimals int) float32, string) {
	switch animalText {
	case dog:
		return animalDog, ""
	case cat:
		return animalCat, ""
	case hamster:
		return animalHamster, ""
	case tarantula:
		return animalTarantula, ""
	default:
		return nil, "animal not found"
	}
}

func multipleAnimals() float32 {
	animalDog, msg := animal(dog)
	if msg != "" {
		fmt.Println(msg)
	}

	animalCat, msg := animal(cat)
	if msg != "" {
		fmt.Println(msg)
	}

	animalHamster, msg := animal(hamster)
	if msg != "" {
		fmt.Println(msg)
	}

	animalTarantula, msg := animal(tarantula)
	if msg != "" {
		fmt.Println(msg)
	}

	var amount float32
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(10)
	amount += animalTarantula(10)

	return amount
}
