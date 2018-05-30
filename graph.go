package main

import "fmt"

type Matrix [][]int

type Node struct {
	index int
	neighbors []Node
}

// Display node information.
func (n Node) display() {
	fmt.Printf("%v\t", n.index)
	for i := range n.neighbors {
		fmt.Printf("%v ", n.neighbors[i].index)
	}
	fmt.Println()
}

type Edge struct {
	node1, node2 Node
	weight int
}

type Graph struct {
	numNodes int
	nodes []Node
	edges []Edge
	adjMatrix Matrix
}

// Determine if a graph is connected.
// Return true if connected.
func (g Graph) isConnected() bool {
	zeroRow := make([]int, g.numNodes)
	for i := range g.adjMatrix {
		if areEqual(zeroRow, g.adjMatrix[i]) {
			return false
		}
	}
	return true
}

// Determine if a graph contains 0 nodes and 0 edges.
// Return true if empty.
func (g Graph) isEmpty() bool {
	zeroRow := make([]int, g.numNodes)
	for i := range g.adjMatrix {
		if !areEqual(zeroRow, g.adjMatrix[i]) {
			return false
		}
	}
	return true
}

// Display graph information.
func (g Graph) display() {
	fmt.Println("\nAdjacency Matrix")
	displayMatrix(g.adjMatrix)
	fmt.Println("\nis empty: ", g.isEmpty())
	fmt.Println("is connected: ", g.isConnected())
	for i := range g.nodes {
		g.nodes[i].display()
	}
}
