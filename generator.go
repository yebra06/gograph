package main

func generateUndirectedGraph(numNodes int) Graph {
	adjMatrix := createSquareMatrix(numNodes)
	nodes := make([]Node, numNodes)
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

	for i := range adjMatrix {
		nodes[i].index = i
		for j := range adjMatrix {
			if adjMatrix[i][j] > 0 {
				nodes[i].neighbors = append(nodes[i].neighbors, Node{index:j})
			}
		}
	}

	g := Graph {
		numNodes: numNodes,
		adjMatrix: adjMatrix,
		nodes: nodes,
	}

	return g
}
