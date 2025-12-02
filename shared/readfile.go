package shared

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// yes this is from a google search, why do you ask?
// I learned C-style file buffer reading decades ago, leave me alone!
func ReadFile(filename string) string {
	// Attempt to open the file
	file, err := os.Open("data/" + filename)
	if err != nil {
		// Handle the error appropriately
		fmt.Println("Error opening file:", err)
		return ""
	}
	// Ensure the file is closed after reading
	defer file.Close()

	var sb strings.Builder

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
