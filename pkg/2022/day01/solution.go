package day01

import (
	"bufio"
	"fmt"
	"strconv"
)

type Elf struct {
	ListOfItems []int
}

func (e *Elf) TotalCalories() int {
	result := 0
	for _, item := range e.ListOfItems {
		result += item
	}
	return result
}

func ReadLineByLine(scanner *bufio.Scanner, elves *[]Elf) {
	newElf := &Elf{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error occurred when converting scanner.Text to integer.")
				continue
			}
			newElf.ListOfItems = append(newElf.ListOfItems, calories)
		} else {
			*elves = append(*elves, *newElf)
			newElf = &Elf{}
		}
	}
	// Append the last elf if the file does not end with a newline
	if len(newElf.ListOfItems) > 0 {
		*elves = append(*elves, *newElf)
	}
}

func HighestCalorieElve(elves *[]Elf) int {
	highestCount := 0

	for _, item := range *elves {
		total := item.TotalCalories()
		if total > highestCount {
			highestCount = total
		}
	}
	return highestCount

}

func Solve(scanner *bufio.Scanner) int {
	var elves []Elf
	ReadLineByLine(scanner, &elves)
	highestCount := HighestCalorieElve(&elves)

	return highestCount
}
