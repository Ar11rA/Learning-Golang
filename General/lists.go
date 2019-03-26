package main

import (
	"errors"
	"fmt"
)

type list []int

func (l list) last() int {
	return l[len(l)-1]
}

func (l list) indexOf(n int) int {
	return l[n-1]
}

func (l list) first() int {
	return l[0]
}

func (l list) reverse() list {
	reversedList := make(list, len(l))
	for i, j := 0, len(l)-1; i <= j; i, j = i+1, j-1 {
		reversedList[i], reversedList[j] = l[j], l[i]
	}
	return reversedList
}

func (l list) removeDuplicates() (list, map[int]int) {
	filteredList := list{}
	elementCount := map[int]int{}
	for i := 0; i < len(l); i++ {
		currentElementCount, hasKey := elementCount[l[i]]
		if hasKey {
			elementCount[l[i]] = currentElementCount + 1
		} else {
			elementCount[l[i]] = 1
			filteredList = append(filteredList, l[i])
		}
	}
	return filteredList, elementCount
}

func (l list) slice(begin, end int) (list, error) {
	slicedList := list{}
	listLength := len(l)
	start := begin
	if begin >= 0 && end < listLength {
		for i := start; i <= end; i++ {
			slicedList = append(slicedList, l[i])
		}
		return slicedList, nil
	}
	return slicedList, errors.New("Starting and ending indexes invalid")
}

func main() {
	numbers := list{1, 2, 3, 5, 4, 6, 5, 6, 7}
	fmt.Println("First element", numbers.first())
	fmt.Println("Last element", numbers.last())
	fmt.Println("Nth element", numbers.indexOf(2))
	fmt.Println("Reversed elements", numbers.reverse())
	numbersWithoutDuplicates, numberCountMap := numbers.removeDuplicates()
	fmt.Println("After removal of elements", numbersWithoutDuplicates)
	fmt.Println("Number count map", numberCountMap)
	slicedNumbers, _ := numbers.slice(3, 7)
	fmt.Println("Numbers sliced from position 3 to 7", slicedNumbers)
}
