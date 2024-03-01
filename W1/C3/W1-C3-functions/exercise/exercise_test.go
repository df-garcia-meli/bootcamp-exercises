package exercise

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSalaryDeduction(t *testing.T) {
	t.Run("Below 50.000", func(t *testing.T) {
		// Given
		var salary float32 = 40000.0
		var expectedResult float32 = 0.0

		// When
		result := salaryDeduction(salary)

		// Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("More than 50.000", func(t *testing.T) {
		// Given
		var salary float32 = 60000.0
		var expectedResult float32 = 0.17

		// When
		result := salaryDeduction(salary)

		// Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("More than 150.000", func(t *testing.T) {
		// Given
		var salary float32 = 1500000.0
		var expectedResult float32 = 0.27

		// When
		result := salaryDeduction(salary)

		// Then
		require.Equal(t, expectedResult, result)
	})
}

func TestEllipsisAverage(t *testing.T) {
	//Given
	numbers := []float32{1, 2, 3, 4, 5}
	var expectedResult float32 = 3.0

	//When
	result := ellipsisAverage(numbers...)

	//Then
	require.Equal(t, expectedResult, result)
}

func TestSalary(t *testing.T) {
	t.Run("Category A", func(t *testing.T) {
		//Given
		totalMinutes := 60
		category := CategoryA
		expectedResult := 1000.0

		//When
		result := salary(totalMinutes, category)

		//Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("Category B", func(t *testing.T) {
		//Given
		totalMinutes := 60
		category := CategoryB
		expectedResult := 1800.0

		//When
		result := salary(totalMinutes, category)

		//Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("Category C", func(t *testing.T) {
		//Given
		totalMinutes := 60
		category := CategoryC
		expectedResult := 4500.0

		//When
		result := salary(totalMinutes, category)

		//Then
		require.Equal(t, expectedResult, result)
	})
}

func TestMultipleFuncs(t *testing.T) {
	t.Run("Min", func(t *testing.T) {
		//Given
		currentOperation := minimum
		minFunc, _ := operation(currentOperation)
		var expectedResult float32 = 2.0

		//When
		result := minFunc(2, 3, 4, 5)

		//Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("Max", func(t *testing.T) {
		//Given
		currentOperation := maximum
		minFunc, _ := operation(currentOperation)
		var expectedResult float32 = 5.0

		//When
		result := minFunc(2, 3, 4, 5)

		//Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("Avg", func(t *testing.T) {
		//Given
		currentOperation := average
		minFunc, _ := operation(currentOperation)
		var expectedResult float32 = 3.0

		//When
		result := minFunc(2, 3, 4)

		//Then
		require.Equal(t, expectedResult, result)
	})

}

func TestAnimals(t *testing.T) {
	t.Run("Dogs", func(t *testing.T) {
		//Given
		totalAnimals := 10
		currentAnimal := dog
		animalFunc, _ := animal(currentAnimal)
		var expectedResult float32 = 100.0

		//When
		result := animalFunc(totalAnimals)

		//Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("Cats", func(t *testing.T) {
		//Given
		totalAnimals := 10
		currentAnimal := cat
		animalFunc, _ := animal(currentAnimal)
		var expectedResult float32 = 50.0

		//When
		result := animalFunc(totalAnimals)

		//Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("Hamsters", func(t *testing.T) {
		//Given
		totalAnimals := 10
		currentAnimal := hamster
		animalFunc, _ := animal(currentAnimal)
		var expectedResult float32 = 2.5

		//When
		result := animalFunc(totalAnimals)

		//Then
		require.Equal(t, expectedResult, result)
	})

	t.Run("Tarantulas", func(t *testing.T) {
		//Given
		totalAnimals := 10
		currentAnimal := tarantula
		animalFunc, _ := animal(currentAnimal)
		var expectedResult float32 = 1.5

		//When
		result := animalFunc(totalAnimals)

		//Then
		require.Equal(t, expectedResult, result)
	})

}
