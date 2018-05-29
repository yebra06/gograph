package main

import "math/rand"

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
	if len(x) != len(y) {return false}
	for i := range(x) {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}