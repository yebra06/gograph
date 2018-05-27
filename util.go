package main

import "math/rand"

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}