package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
)

func OneOne() {
	input := shared.GetPuzzleInput("1.txt", "\n")
	var lockPosition = 50
	zeroCount := 0

	for _, instruction := range input {
		direction := string(instruction[0])
		numberStr := string(instruction[1:])

		fmt.Print(direction + ":" + numberStr + " => ")

		if by, err := strconv.Atoi(numberStr); err == nil {
			if direction == "R" {
				lockPosition = oneOneIncrement(lockPosition, by)
			} else {
				lockPosition = oneOneDecrement(lockPosition, by)
			}

			fmt.Printf("lockPosition is: %d\n", lockPosition)

			if lockPosition == 0 {
				zeroCount++
			}
		} else {
			fmt.Println("error atoi'ing")
		}
	}

	fmt.Printf("zero count: %d", zeroCount)
}

// there's probably a clever math way with modulus or some bullshit, but I minored in Drama
func oneOneIncrement(lockPosition int, by int) int {
	for i := 0; i < by; i++ {
		if lockPosition == 99 {
			lockPosition = 0
		} else {
			lockPosition++
		}
	}

	return lockPosition
}

func oneOneDecrement(lockPosition int, by int) int {
	for i := 0; i < by; i++ {
		if lockPosition == 0 {
			lockPosition = 99
		} else {
			lockPosition--
		}
	}

	return lockPosition
}
