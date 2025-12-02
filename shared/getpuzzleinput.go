package shared

import (
	"strings"
)

func GetPuzzleInput() []string {
	// example input
	//	s := `L68
	//L30
	//R48
	//L5
	//R60
	//L55
	//L1
	//L99
	//R14
	//L82`

	s := ReadFile("1.txt")

	array := strings.Split(s, "\n")

	return array
}
