package main

import (
	"math/rand"
	"fmt"
)

// Generate random integer number.
func random(min, max int) int {
	return rand.Intn(max - min) + min
}

// Find absolute value of number.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Return true if 2 integer arrays are equal.
func areEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// Display matrix.
func displayMatrix(matrix [][]int) {
	for row := range matrix {
		fmt.Println(matrix[row])
	}
}

// Create and return new size X size matrix.
func createSquareMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return matrix
}
