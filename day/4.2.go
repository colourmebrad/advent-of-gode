package day

import (
	"advent-of-gode/shared"
	"fmt"
)

func (stack StackOfPaper) getForked() {
	for row := 0; row < len(stack.rolls); row++ {
		for col := 0; col < len(stack.rolls[row]); col++ {
			if stack.rolls[row][col] == "x" {
				stack.rolls[row][col] = "."
			}
		}
	}

	stack.Print()
}

func FourTwo() {
	input := shared.GetPuzzleInput("4.txt", "\n")

	/*
		..@@.@@@@.
		@@@.@.@.@@
		@@@@@.@.@@
		@.@@@@..@.
		@@.@@@@.@@
		.@@@@@@@.@
		.@.@.@.@@@
		@.@@@.@@@@
		.@@@@@@@@.
		@.@.@@@.@.
	*/

	stack := initFromStringArray(input)

	stack.Print()

	totalForked := 0

	forkable := stack.countForkable(4)
	for forkable > 0 {
		totalForked += forkable
		stack.getForked()
		forkable = stack.countForkable(4)
	}

	fmt.Printf("Removed a total of %d rolls\n", totalForked)

	// The forklifts can only access a roll of paper if there are fewer than four rolls of paper in the eight adjacent positions.

}
