package main

import (
	"advent-of-gode/day"
	"advent-of-gode/shared"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var dayNum int = -1
	var err error = nil

	if len(os.Args) > 1 {
		if dayNum, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Printf("Argument '%s' is not a valid day number", os.Args[1])
			return
		}
	} else {
		fmt.Printf("Use \"go run main.go [1-12]\" for running the exercise for the specified day")
	}

	input := shared.GetPuzzleInput()

	switch dayNum {
	case 1:
		fmt.Printf("Running day 1")
		day.One(input)
	default:
		fmt.Printf("Day %d is not implemented yet", dayNum)
	}
}
