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

func FibonacciRecursive(n int) int {
	if n <= 0 {
		return -1
	} else if n > 2 {
		return fibonacci(n-2) + fibonacci(n-1)
	}

	return n - 1
}

func FibonacciCloser() func() int {
	first, second := 0, 1
	return func() int {
		n := first
		first, second = second, first+second
		return n
	}
}

func WordCount(s string) map[string]int {// Считает количество слов в строке
	wordsMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		wordsMap[word]++
	}
	return wordsMap
}

func MultiFunc() func(int) int {
	multiper := 1
	return func(number int) int {
		multiper *= number
		return multiper
	}

}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (read MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// ROT13 Cipher
func Rot13Cipher(c byte) byte {
	switch {
	case c >= 'A' && c <= 'Z': // Для заглавных букв
		return 'A' + (c-'A'+13)%26
	case c >= 'a' && c <= 'z': // Для строчных букв
		return 'a' + (c-'a'+13)%26
	default:
		return c // Неизменные символы
	}
}

func IsPalindrom(str string) bool{
    n := len(str)
    for i:= 0; i < n / 2; i++{
        if str[i] != str[n - 1 - i]{
            return false
        } 
    }
    return true
}


func SumOfNum(number int) int {
	number = abs(number)
	str := strconv.Itoa(number)
	s := 0
	for _, ch := range str {
		s += int(ch - '0')
	}
	return s
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
//Func to move value x in a slice to the last position
func ShiftEnd[T any](s []T, x int) []T {
    if x < 0 {
        return s
    }
    if x >= len(s)-1 {
        return s
    }
    tmp := s[x]
    s = append(s[:x], s[x+1:]...)
    s = append(s, tmp)
    return s
}
