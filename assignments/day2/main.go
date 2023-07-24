// Goal: Create a package and a
// command-line tool to sort input provided
// by the user.
// Inputs: Number (integer or float) array,
// string array.
// Outputs: Sorted result based on the
// provided input type.

// Create a Go package with functions for sorting integer arrays, float
// arrays, and string arrays.
// Implement sorting logic for each data type using appropriate
// algorithms.

// Create a command-line tool (CLI) to parse the input from the
// command line.
// Determine the type of input (integer array, float array, string array,
// or mixed).
// Utilize the corresponding sorting function from the package to sort
// the elements.
// Output the sorted result.

// Use the flag package to parse command line arguments.
// Create separate functions in the package for sorting each data
// type.
// Consider implementing generic sorting functions using interfaces
// to handle mixed input types.

// Path: assignments\day2\sort\sort.go
// Goal: Create a package and a
// command-line tool to sort input provided
// by the user.
// Inputs: Number (integer or float) array,
// string array.
// Outputs: Sorted result based on the
// provided input type.

// Create a Go package with functions for sorting integer arrays, float
// arrays, and string arrays.
// Implement sorting logic for each data type using appropriate
// algorithms.

// Create a command-line tool (CLI) to parse the input from the
// command line.
// Determine the type of input (integer array, float array, string array,
// or mixed).

// Utilize the corresponding sorting function from the package to sort
// the elements.
// Output the sorted result.

// Use the flag package to parse command line arguments.
// Create separate functions in the package for sorting each data
// type.
// Consider implementing generic sorting functions using interfaces
// to handle mixed input types.

package main

import (
	"github.com/its-rav/dwarves-go-23/ass-d2/cmd"
)

var supportedTypes = []string{"number", "string", "mixed"}

func main() {
	cmd.NewCommand().Execute()
}
