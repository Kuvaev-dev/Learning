package slice

import (
	"fmt"
)

func GetSlices() {
	testSlice := [][]int{
		{6, 6, 6},
		{7, 7, 7},
		{8, 8, 8},
	}

	// Print
	// PrintSlices(testSlice)

	// Add
	// PrintSlices(AddSlice(testSlice, []int{9, 9, 9}))

	// _
	PrintSlicesChanged(testSlice)
}

func PrintSlices(slices [][]int) {
	for i := range slices {
		fmt.Printf("%v\n", slices[i])
	}
}

func PrintSlicesChanged(slices [][]int) {
	for _, value := range slices {
		fmt.Printf("%v\n", value)
	}
}

func AddSlice(slices [][]int, newSlice []int) [][]int {
	return append(slices, newSlice)
}
