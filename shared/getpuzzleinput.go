package shared

import (
	"log"
	"os"
	"strings"
)

func GetPuzzleInput(filename string, splitOn string) []string {
	// shout Nick for this read file code that is 4 lines instead of the 30 lines that I got from Google. Sheesh.
	input, err := os.ReadFile("data/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), splitOn)
}
