package main

import (
	"fmt"
	"strings"
	"time"
)

// Concatination with operator +=
func ConcatWithPlus(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

// Concatination with strings.Join
func ConcatWithJoin(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	// Create big slice with 50 000 elements inside
	bigArgs := make([]string, 50000)
	for i := range bigArgs {
		bigArgs[i] = "item"
	}

	// Time for +=
	start := time.Now()
	_ = ConcatWithPlus(bigArgs)
	time_plus := time.Since(start)
	fmt.Printf("Loop with +: %v\n", time_plus)

	// Time for strings.Join
	start = time.Now()
	_ = ConcatWithJoin(bigArgs)
	time_join := time.Since(start)
	fmt.Printf("strings.Join: %v\n", time_join)

	// Compare the two times
	if time_plus < time_join {
		result_speed := float64(time_join.Nanoseconds()) / float64(time_plus.Nanoseconds())
		fmt.Printf("Using += is faster than strings.Join %.2f times.\n", result_speed)
	} else {
		result_speed := float64(time_plus.Nanoseconds()) / float64(time_join.Nanoseconds())
		fmt.Printf("Using strings.Join is faster than += %.2f times.\n", result_speed)
	}
}
