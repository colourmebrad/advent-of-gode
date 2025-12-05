package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
)

func ThreeOne() {
	/* 987654321111111
	   811111111111119
	   234234234234278
	   818181911112111
		 100100000010009
		 100100000080709
	*/

	/*bank := "4122333122212322225322231334325214312232223232222412212214232124422231324324242122124222322242232333"
	fmt.Printf("%s = %d\n", bank, getMaxJoltage(bank))
	return*/

	input := shared.GetPuzzleInput("3.debug.txt", "\n")

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
			i++
			secondDigit, _ = strconv.Atoi(string(bank[i]))
		} else if thisDigit > secondDigit {
			secondDigit = thisDigit
		}
	}

	fmt.Printf("getMaxJoltage(%s) = %d\n", bank, firstDigit*10+secondDigit)

	return firstDigit*10 + secondDigit
}
