package main

// Go program to concatenate command-line arguments
// Concatenates all command-line arguments into a single string, separated by spaces.

import (
	"fmt"
	"os"
)

func main() {
	// Define the string to hold the concatenated arguments and
	// initialize it with an empty string
    var s, sep string

    // Iterate over all command-line arguments, starting from index 1 (since index 0 is the program name)
    for i := 1; i < len(os.Args); i++ {
        // Append the current argument to the string, separated by a space
        s += sep + os.Args[i]
		// Update the separator to a space for the next iteration
		sep = " "
    }

    // Print the concatenated string
    fmt.Print(s)
}
