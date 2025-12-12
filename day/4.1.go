package day

import (
	"advent-of-gode/shared"
	"fmt"
)

type StackOfPaper struct {
	rolls [][]string
}

func (stack StackOfPaper) countForkable(rollLimit int) int {
	forkable := 0

	for row := 0; row < len(stack.rolls); row++ {
		for col := 0; col < len(stack.rolls[row]); col++ {
			if stack.rolls[row][col] != "." && stack.isRollForkable(row, col, rollLimit) {
				forkable++
			}
		}
	}

	fmt.Println("Marked forkable")
	stack.Print()

	return forkable
}

func (stack StackOfPaper) isRollForkable(row int, col int, rollLimit int) bool {
	numRows := len(stack.rolls)
	numColumns := len(stack.rolls[0])

	rowStart := max(row-1, 0)
	rowEnd := min(row+1, numRows-1)

	colStart := max(col-1, 0)
	colEnd := min(col+1, numColumns-1)

	adjacentRolls := 0

	for r := rowStart; r <= rowEnd; r++ {
		for c := colStart; c <= colEnd; c++ {
			if !(row == r && col == c) && stack.rolls[r][c] != "." {
				adjacentRolls++
			}
		}
	}

	if adjacentRolls < rollLimit {
		stack.rolls[row][col] = "x" // really just for debugging, but it wouldn't surprise me if 4.2 needed these one way or another
	}

	return adjacentRolls < rollLimit
}

func (stack StackOfPaper) Print() {
	for row := 0; row < len(stack.rolls); row++ {
		for col := 0; col < len(stack.rolls[row]); col++ {
			fmt.Print(stack.rolls[row][col])
		}
		fmt.Print("\n")
	}
}

func initFromStringArray(input []string) StackOfPaper {
	var stack StackOfPaper

	for _, x := range input {
		roll := []string{}
		stack.rolls = append(stack.rolls, roll)
		for _, y := range x {
			stack.rolls[len(stack.rolls)-1] = append(stack.rolls[len(stack.rolls)-1], string(y))
		}
	}

	return stack
}

func FourOne() {
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

	fmt.Printf("There are %d forkable rolls\n", stack.countForkable(4))

	// The forklifts can only access a roll of paper if there are fewer than four rolls of paper in the eight adjacent positions.

}
