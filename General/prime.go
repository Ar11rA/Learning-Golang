package main

import "fmt"

func printPrimeTill(limit int) []int {
	var numbers []int
	for i := 1; i <= limit; i++ {
		if isPrime(i) {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

func isPrime(number int) bool {
	ctr := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			ctr = false
			break
		}
	}
	return ctr
}

func main() {
	fmt.Println(printPrimeTill(25))
}
