package main

import "fmt"

type Matrix [][]int

type Node struct {
	index int
	neighbors []int
}

// Display node information.
func (n Node) display() {
	fmt.Printf("%v\t", n.index)
	for i := range n.neighbors {
		fmt.Printf("%v ", n.neighbors[i])
	}
	fmt.Println()
}

type Edge struct {
	node1, node2, weight int
}

// Display edge information.
func (e Edge) display() {
	fmt.Println(e.node1, " ", e.node2, " ", e.weight)
}

type Graph struct {
	numNodes int
	adjMatrix Matrix
	nodes []Node
	edges []Edge
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

// Determine if a graph contains an edge.
// Return true if graph has it.
func (g Graph) hasEdge(e Edge) bool {
	for i := range g.edges {
		if g.edges[i] == e {
			return true
		}
	}
	return false
}

// Display graph information.
func (g Graph) display() {
	fmt.Println("\nAdjacency Matrix")
	displayMatrix(g.adjMatrix)
	fmt.Println("\nis empty: ", g.isEmpty())
	fmt.Printf("is connected: %t\n\n", g.isConnected())
	for i := range g.nodes {
		g.nodes[i].display()
	}
	fmt.Println()
	for i := range g.edges {
		g.edges[i].display()
	}
}
