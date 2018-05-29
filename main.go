package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Display adjacency matrix.
func displayMatrix(matrix [][]int) {
	for row := range matrix {
		fmt.Println(matrix[row])
	}
}

// Create adjacency matrix representation of graph.
func createConnections(adjMatrix [][]int) {
	connections := make([]int, len(adjMatrix[0]))
	for i := range connections {
		connections[i] = random(0, 2)
	}

	for i := range connections {
		for j := range connections {
			if i != j {
				adjMatrix[i][j] = connections[abs(i-j)]
			}
		}
	}
}

// Determine if a graph is connected.
// Return true if connected.
func isConnected(matrix [][]int) bool {
	zeroRow := make([]int, len(matrix[0]))
	for i := range matrix {
		if areEqual(zeroRow, matrix[i]) {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().Unix())

	numNodes := 4
	adjMatrix := make([][]int, numNodes)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, numNodes)
	}

	createConnections(adjMatrix)
	displayMatrix(adjMatrix)
}