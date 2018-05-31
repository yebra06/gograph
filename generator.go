package main

func generateUndirectedGraph(numNodes int) Graph {
	g := Graph {
		numNodes: numNodes,
		adjMatrix: createSquareMatrix(numNodes),
		nodes: make([]Node, 0, numNodes),
		edges: make([]Edge, 0, numNodes * (numNodes - 1) / 2),
	}

	weights := randomArray(0, 2, numNodes)
	for i := range weights {
		g.nodes = append(g.nodes, Node{i, []int{}})
		for j := range weights {
			if i != j {
				weight := weights[abs(i - j)]
				g.adjMatrix[i][j] = weight
				if weight > 0 {
					g.nodes[i].neighbors = append(g.nodes[i].neighbors, j)
					if !g.hasEdge(Edge{j, i, g.adjMatrix[j][i]}) {
						g.edges = append(g.edges, Edge{i, j, g.adjMatrix[i][j]})
					}
				}
			}
		}
	}

	return g
}
