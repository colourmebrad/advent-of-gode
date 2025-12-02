package day

import (
	"fmt"
	"strconv"
)

func One(input []string) {
	var lockPosition = 50
	zeroCount := 0

	for _, instruction := range input {
		direction := string(instruction[0])
		numberStr := string(instruction[1:])

		fmt.Printf(direction + ":" + numberStr + " => ")

		if by, err := strconv.Atoi(numberStr); err == nil {
			if direction == "R" {
				lockPosition = increment(lockPosition, by)
			} else {
				lockPosition = decrement(lockPosition, by)
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
func increment(lockPosition int, by int) int {
	for i := 0; i < by; i++ {
		if lockPosition == 99 {
			lockPosition = 0
		} else {
			lockPosition++
		}
	}

	return lockPosition
}

func decrement(lockPosition int, by int) int {
	for i := 0; i < by; i++ {
		if lockPosition == 0 {
			lockPosition = 99
		} else {
			lockPosition--
		}
	}

	return lockPosition
}
