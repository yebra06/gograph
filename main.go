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