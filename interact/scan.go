package interact

import (
	"fmt"
	"strconv"
)

// ScanInputAndCheck scans user input with check.
//
// It returns -1 if input is not a number or is less than 0.
//
// If nothing is input, return -2
func ScanInputAndCheck() int {
	in := ""
	fmt.Print("> ")
	fmt.Scanln(&in)
	println()
	if in == "" {
		return -2
	}
	op, err := strconv.Atoi(in)
	if err != nil {
		return -1
	}
	if op < 0 {
		return -1
	}
	return op
}

// ScantoSting scans user input to string.
func ScantoSting() string {
	input := ""
	fmt.Scanln(&input)
	return input
}
