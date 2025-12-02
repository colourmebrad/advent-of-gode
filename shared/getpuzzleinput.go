package shared

import (
	"strings"
)

func GetPuzzleInput(filename string) []string {
	s := ReadFile(filename)

	array := strings.Split(s, "\n")

	return array
}
