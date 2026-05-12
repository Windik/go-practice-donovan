package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("Argument %d: %s\n", index+1, arg)
	}
}
