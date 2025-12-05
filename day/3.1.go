package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
)

func ThreeOne() {
	input := shared.GetPuzzleInput("3.txt", "\n")

	totalJoltage := 0

	for _, bank := range input {
		totalJoltage += getMaxJoltage(bank)
	}

	fmt.Printf("Total Maximum Joltage is: %d\n", totalJoltage)
}

func getMaxJoltage(bank string) int {
	// 818181911112111

	firstDigit, _ := strconv.Atoi(string(bank[0]))
	secondDigit, _ := strconv.Atoi(string(bank[1]))

	for i := 1; i < len(bank); i++ {
		//fmt.Printf("%d%d\n", firstDigit, secondDigit)

		thisDigit, _ := strconv.Atoi(string(bank[i]))

		// inelegant, but whatever, I need to catch the case of the last digit being the best second digit
		if i == len(bank)-1 {
			if thisDigit > secondDigit {
				secondDigit = thisDigit
			}
			break
		}

		if thisDigit > firstDigit {
			firstDigit = thisDigit
			// grab the next digit as it's our new baseline second digit
			secondDigit, _ = strconv.Atoi(string(bank[i+1]))
		} else if thisDigit > secondDigit {
			secondDigit = thisDigit
		}
	}

	//fmt.Printf("getMaxJoltage(%s) = %d\n", bank, firstDigit*10+secondDigit)

	return firstDigit*10 + secondDigit
}
