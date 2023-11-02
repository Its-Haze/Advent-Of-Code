package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/its-haze/advent-of-code/pkg/2022/day01"
)

func main() {
	//input, err := os.ReadFile(filepath.Join("input", "2022", "day01.txt"))
	/* 	if err != nil {
		log.Fatal(err)
	} */

	file, err := os.Open(filepath.Join("input", "2022", "day01.txt"))
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	result := day01.Solve(scanner)
	fmt.Println("Result:", result)
}
