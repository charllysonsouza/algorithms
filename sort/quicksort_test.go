package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Quicksort(t *testing.T) {
	testCases := []struct {
		name    string
		numbers []int
	}{
		{
			name:    "already sorted",
			numbers: []int{1, 2, 3, 4, 5, 6, 9, 20, 22, 23, 28, 32, 34, 39, 40, 42, 76, 87, 99, 112},
		},
		{
			name:    "inversed",
			numbers: []int{117, 90, 88, 83, 81, 77, 74, 69, 64, 63, 51, 50, 49, 42, 41, 34, 32, 29, 28, 22, 16, 8, 6, 5, 3, 1},
		},
		{
			name:    "repeated",
			numbers: []int{7, 7, 7, 7, 7, 1, 1, 9, 9, 0, 4, 4, 4, 5, 4, 5, 7, 1},
		},
		{
			name:    "aleatory",
			numbers: rand.Perm(1000000),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fmt.Printf("numbers to be sorted: %v\n", testCase.numbers)
			sortedNumbers := Quicksort(testCase.numbers)
			fmt.Printf("sorted numbers: %v\n", sortedNumbers)
		})
	}
}
