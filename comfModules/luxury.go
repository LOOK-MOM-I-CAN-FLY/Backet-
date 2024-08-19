package comfModules

import (
	"fmt"
	"errors"
	"log"
)
func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func Swap(a int, b int) {
	tmp := a
	a = b
	b = tmp
}

func Factorial(number int) int {
	if number == 1 || number == 0 {
		return number
	}
	return number * Factorial(number-1)
}

// Hello returns a greeting for the named person.
func Wassup(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func AverageValue(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}
	summary := 0
	for _, val := range values {
		summary += val
	}
	return float64(summary) / float64(len(values))
}

func Dividers(number int) []int {
	divisors := make([]int, 0)
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func Fibonacci(n int) int {
	if n <= 0 {
		return -1
	} else if n > 2 {
		return fibonacci(n-2) + fibonacci(n-1)
	}

	return n - 1
}
