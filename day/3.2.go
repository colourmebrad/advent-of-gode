package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
)

func ThreeTwo() {
	input := shared.GetPuzzleInput("3.txt", "\n")

	totalJoltage := 0

	for _, bankStr := range input {
		totalJoltage += getMaxJoltage12(convertStringToBank(bankStr))
	}

	fmt.Printf("Total Maximum Joltage is: %d\n", totalJoltage)
}

func convertStringToBank(bankStr string) []int {
	var bank []int
	for i := 0; i < len(bankStr); i++ {
		thisDigit, _ := strconv.Atoi(string(bankStr[i]))
		bank = append(bank, thisDigit)
	}

	return bank
}

func getMaxJoltage12(bank []int) int {
	// {1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4}
	//fmt.Printf("getMaxJoltage(%v)\n", bank)

	digits := []int{}
	sliceWindowStart := 0

	for len(digits) < 12 {
		currentDigit := len(digits)
		sliceWindowEnd := len(bank) + currentDigit - 11

		//fmt.Printf("currentDigit := %d\n", currentDigit)
		//fmt.Printf("sliceWindowStart := %d\n", sliceWindowStart)
		//fmt.Printf("sliceWindowEnd := %d\n", sliceWindowEnd)

		// find the max value within the slice of the bank that we can use
		indexOfMax := maxValIndex(bank[sliceWindowStart:sliceWindowEnd]) + sliceWindowStart
		sliceWindowStart = indexOfMax + 1

		digits = append(digits, bank[indexOfMax])

		//fmt.Printf("digits is %v\n", digits)
	}

	voltage := 0
	multiplier := 1

	for i := 11; i >= 0; i-- {
		voltage += (digits[i] * multiplier)
		multiplier *= 10
	}

	//fmt.Printf("Voltage = %v\n", voltage)

	return voltage
}

func maxValIndex(in []int) int {
	maxIndex := 0

	for index, val := range in {
		if val > in[maxIndex] {
			maxIndex = index
		}
	}

	return maxIndex
}
