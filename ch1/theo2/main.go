package theo2

// Go program to concatenate command-line arguments
// uses a for loop with range to iterate over command-line arguments
import (
	"fmt"
	"os"
)

func main() {
	// Initialize an empty string and a separator (space) for the concatenated string.
	s, sep := "", " "

	// Iterate over all command-line arguments using a for loop with range.
	for _, arg := range os.Args[1:] {
        s += sep + arg
		sep = " "
    }

	// Print the concatenated string.
	fmt.Print(s)
}