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
	fmt.Printf("Checking %s\n", intString)
	sequence := ""
	for y := 0; y < len(intString)-1; y++ {
		sequence = sequence + string(intString[y]) // why do I have to string this, that's so annoying (oh, I bet it's unicode bullshit)

		// the problem I'm running into is that this code works for 11 and 12341234 but erroneously returns true for 110
		// I need to ensure there aren't any non-repeated characters
		// 121212 is valid
		// 12125 is not
		// 11 is valid
		// 211 is not and 112 is not

		// maybe here is where we want another loop?
		// we basically have to take sequence and then ensure that the rest of the strong is ONLY 1+ loops of that sequence

		// is there enough string left for another of this sequence?
		nextSequenceStart := y + 1
		nextSequenceEnd := y + 1 + len(sequence)

		if nextSequenceEnd < len(intString)+1 {
			if sequence == intString[nextSequenceStart:nextSequenceEnd] {
				fmt.Printf("%s == %s\n", sequence, intString[nextSequenceStart:nextSequenceEnd])
				return true
			} else {
				//fmt.Printf("%s != %s\n", sequence, intString[nextSequenceStart:nextSequenceEnd])
			}
		} else {
			//fmt.Printf("%d >= %d\n", nextSequenceEnd, len(intString))
		}
	}

	return false
}

func TwoTwo() {
	input := shared.GetPuzzleInput("2.example.txt", ",")

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
