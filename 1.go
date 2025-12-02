package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getPuzzleInput()
	fmt.Println(input)

	var lockPosition = 50
	zeroCount := 0

	for _, instruction := range input {
		//for j := 0; j < len(input); j++ {
		direction := string(instruction[0])
		numberStr := string(instruction[1:len(instruction)])

		fmt.Println(direction + ":" + numberStr)

		if by, err := strconv.Atoi(numberStr); err == nil {
			if direction == "R" {
				lockPosition = increment(lockPosition, by)
			} else {
				lockPosition = decrement(lockPosition, by)
			}

			fmt.Printf("lockPosition is: %d\n", lockPosition)

			if lockPosition == 0 {
				zeroCount++
			}
		} else {
			fmt.Println("error atoi'ing")
		}
	}

	fmt.Printf("zero count: %d", zeroCount)
}

// there's probably a clever math way with modulus or some bullshit, but I minored in Drama
func increment(lockPosition int, by int) int {
	for i := 0; i < by; i++ {
		if lockPosition == 99 {
			lockPosition = 0
		} else {
			lockPosition++
		}
	}

	return lockPosition
}

func decrement(lockPosition int, by int) int {
	for i := 0; i < by; i++ {
		if lockPosition == 0 {
			lockPosition = 99
		} else {
			lockPosition--
		}
	}

	return lockPosition
}

func getPuzzleInput() []string {
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

	s := readInput()

	array := strings.Split(s, "\n")

	return array
}

// yes this is from a google search, why do you ask?
// I learned c-style file buffer reading decades ago, leave me alone!
func readInput() string {
	// Attempt to open the file
	file, err := os.Open("1.txt")
	if err != nil {
		// Handle the error appropriately
		fmt.Println("Error opening file:", err)
		return ""
	}
	// Ensure the file is closed after reading
	defer file.Close()

	var sb strings.Builder
	sb.Grow(20000)

	// Create a buffer to hold the file's content
	buffer := make([]byte, 1024) // Adjust buffer size as needed
	for {
		// Read from the file into the buffer
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// End of file reached
				break
			}
			// Handle other potential errors
			fmt.Println("Error reading file:", err)
			return ""
		}
		sb.WriteString(string(buffer[:bytesRead]))
	}
	return sb.String()
}
