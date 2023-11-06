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

	file, err := os.Open(filepath.Join("input", "2022", "day01.txt"))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	result, result2 := day01.Solve(scanner)
	fmt.Println("Result1:", result)
	fmt.Println("Result2:", result2)

}
