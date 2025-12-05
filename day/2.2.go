package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
	"strings"
)

func (idRange IDRange) getRepeatedCycles() []int {
	var repeats []int

	var r = idRange.getRangeAsStringArray()

	for _, stringVal := range r {
		if hasRepeatedCycle(stringVal) {
			intVal, _ := strconv.Atoi(stringVal)
			repeats = append(repeats, intVal)
		}
	}

	return repeats
}

func hasRepeatedCycle(intString string) bool {
	//fmt.Printf("Checking %s\n", intString)
	sequence := ""
	for y := 0; y < len(intString)-1; y++ {
		sequence = sequence + string(intString[y]) // why do I have to string this, that's so annoying (oh, I bet it's unicode bullshit)

		if isRepeatedCycleOf(sequence, intString[y+1:]) {
			return true
		}
	}

	return false
}

func isRepeatedCycleOf(sequence string, input string) bool {
	//fmt.Printf("isRepeatedCycleOf(%s, %s)\n", sequence, input)

	if sequence == input {
		return true
	}

	end := min(len(sequence), len(input))

	if sequence == input[0:end] {
		return isRepeatedCycleOf(sequence, input[len(sequence):])
	}
	return false
}

func TwoTwo() {
	input := shared.GetPuzzleInput("2.txt", ",")

	sum := 0

	for _, rangeStr := range input {
		var idRange IDRange
		parts := strings.Split(rangeStr, "-")
		idRange.start, _ = strconv.Atoi(parts[0])
		idRange.end, _ = strconv.Atoi(parts[1])

		cycles := idRange.getRepeatedCycles()

		for _, val := range cycles {
			sum += val
		}
	}

	fmt.Printf("Sum is: %d\n", sum)
}
