package graph

import "math/big"

type Graph struct {
	Edges Edges
	Nodes Nodes
}

// BMF
// A-(-1)-B-(5)-E=4(weight)
// A-(-1)-B-(3)-C-(1)-E=3(weight)

// BFS
// A-(-1)-B=1(edge)
// A-(4)-C=1(edge)
// A-(-1)-B-(3)-D=2(edge)
// A-(4)-C-(5)-D=2(edge)
// A-(-1)-B-(5)-E=2(edge)

func BuildGraph() *Graph {
	nodes := Nodes{
		&Node{Index: 0, UUID: "A"},
		&Node{Index: 1, UUID: "B"},
		&Node{Index: 2, UUID: "C"},
		&Node{Index: 3, UUID: "D"},
		&Node{Index: 4, UUID: "E"},
	}
	edges := Edges{
		&Edge{Index: 0, UUID: "A-B", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(-1)},
		&Edge{Index: 1, UUID: "A-C", Source: nodes[0], Target: nodes[2], Weight: new(big.Float).SetInt64(4)},
		&Edge{Index: 2, UUID: "B-C", Source: nodes[1], Target: nodes[2], Weight: new(big.Float).SetInt64(3)},
		&Edge{Index: 3, UUID: "B-D", Source: nodes[1], Target: nodes[3], Weight: new(big.Float).SetInt64(2)},
		&Edge{Index: 4, UUID: "B-E", Source: nodes[1], Target: nodes[4], Weight: new(big.Float).SetInt64(5)},
		&Edge{Index: 5, UUID: "C-D", Source: nodes[2], Target: nodes[3], Weight: new(big.Float).SetInt64(5)},
		&Edge{Index: 6, UUID: "D-B", Source: nodes[3], Target: nodes[1], Weight: new(big.Float).SetInt64(1)},
		&Edge{Index: 7, UUID: "D-E", Source: nodes[3], Target: nodes[4], Weight: new(big.Float).SetInt64(3)},
		&Edge{Index: 8, UUID: "C-E", Source: nodes[2], Target: nodes[4], Weight: new(big.Float).SetInt64(1)},
	}

	return &Graph{Edges: edges, Nodes: nodes}
}
