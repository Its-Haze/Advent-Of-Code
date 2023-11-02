package day01

import (
	"bufio"
	"testing"
)

// Yes i am too lazy to actually implement the unit tests
// These are auto generated, who do you think i am, a developer testing code??? HAHAHA... No

func TestElf_TotalCalories(t *testing.T) {
	type fields struct {
		ListOfItems []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elf{
				ListOfItems: tt.fields.ListOfItems,
			}
			if got := e.TotalCalories(); got != tt.want {
				t.Errorf("Elf.TotalCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLineByLine(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
		elves   *[]Elf
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadLineByLine(tt.args.scanner, tt.args.elves)
		})
	}
}

func TestHighestCalorieElve(t *testing.T) {
	type args struct {
		elves *[]Elf
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighestCalorieElve(tt.args.elves); got != tt.want {
				t.Errorf("HighestCalorieElve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.scanner); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
