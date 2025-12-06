package main

import (
	"advent-of-gode/day"
	"fmt"
	"os"
)

func main() {
	var exercise string = ""

	if len(os.Args) > 1 {
		exercise = os.Args[1]
	} else {
		fmt.Printf("Use \"go run main.go [day].[part]\" for running an exercise")
	}

	switch exercise {
	case "1.1":
		fmt.Printf("Running day 1.1\n")
		day.OneOne()
	case "1.2":
		fmt.Printf("Running day 1.2\n")
		day.OneTwo()
	case "2.1":
		fmt.Printf("Running day 2.1\n")
		day.TwoOne()
	case "2.2":
		fmt.Printf("Running day 2.2\n")
		day.TwoTwo()
	case "3.1":
		fmt.Printf("Running day 3.1\n")
		day.ThreeOne()
	case "3.2":
		fmt.Printf("Running day 3.2\n")
		day.ThreeTwo()
	default:
		fmt.Printf("Exercise %s is not implemented yet", exercise)
	}
}
