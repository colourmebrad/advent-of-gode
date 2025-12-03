package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
	"strings"
)

type IDRange struct {
	start int
	end   int
}

/* The ranges are separated by commas (,); each range gives its first ID and last ID separated by a dash (-).

Since the young Elf was just doing silly patterns,
	you can find the invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice.
	So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.

None of the numbers have leading zeroes; 0101 isn't an ID at all. (101 is a valid ID that you would ignore.) */

func (idRange IDRange) getRangeAsStringArray() []string {
	var out []string
	// TODO: this should definitely be done with Go ranges
	for x := idRange.start; x <= idRange.end; x++ {
		out = append(out, strconv.Itoa(x))
	}

	return out
}

func (idRange IDRange) getDoubleCycles() []int {
	var doubles []int

	var r = idRange.getRangeAsStringArray()

	for _, stringVal := range r {
		if isDoubleCycle(stringVal) {
			intVal, _ := strconv.Atoi(stringVal)
			doubles = append(doubles, intVal)
		}
	}

	return doubles
}

func isDoubleCycle(intString string) bool {
	if len(intString)%2 == 0 {
		firstHalf := intString[0:(len(intString) / 2)]
		secondHalf := intString[len(intString)/2:]

		//fmt.Printf("%s => %s == %s\n", intString, firstHalf, secondHalf)

		return firstHalf == secondHalf
	}
	return false
}

func TwoOne() {
	input := shared.GetPuzzleInput("2.txt", ",")

	sum := 0

	for _, rangeStr := range input {
		var idRange IDRange
		parts := strings.Split(rangeStr, "-")
		idRange.start, _ = strconv.Atoi(parts[0])
		idRange.end, _ = strconv.Atoi(parts[1])

		cycles := idRange.getDoubleCycles()

		for _, val := range cycles {
			sum += val
		}
	}

	fmt.Printf("Sum is: %d\n", sum)
}
