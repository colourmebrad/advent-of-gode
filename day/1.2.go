package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
)

func OneTwo() {
	input := shared.GetPuzzleInput("1.txt", "\n")
	var lockPosition = 50
	var zeroCount = 0
	var newZeroCount = 0

	for _, instruction := range input {
		direction := string(instruction[0])
		numberStr := string(instruction[1:])

		fmt.Print(direction + ":" + numberStr + " => ")

		if by, err := strconv.Atoi(numberStr); err == nil {
			if direction == "R" {
				lockPosition, newZeroCount = oneTwoIncrement(lockPosition, by)
			} else {
				lockPosition, newZeroCount = oneTwoDecrement(lockPosition, by)
			}

			zeroCount += newZeroCount

			fmt.Printf("lockPosition is: %d\n", lockPosition)
		} else {
			fmt.Println("error atoi'ing")
		}
	}

	fmt.Printf("zero count: %d", zeroCount)
}

// there's probably a clever math way with modulus or some bullshit, but I minored in Drama
func oneTwoIncrement(lockPosition int, by int) (int, int) {
	var zeroCount = 0

	for i := 0; i < by; i++ {
		if lockPosition == 99 {
			lockPosition = 0
			zeroCount++

			fmt.Print("[zero!] ")

		} else {
			lockPosition++
		}
	}

	return lockPosition, zeroCount
}

func oneTwoDecrement(lockPosition int, by int) (int, int) {
	var zeroCount = 0

	for i := 0; i < by; i++ {
		if lockPosition == 0 {
			lockPosition = 99
		} else {
			lockPosition--
			if lockPosition == 0 {
				zeroCount++
				fmt.Print("[zero!] ")
			}
		}
	}

	return lockPosition, zeroCount
}
