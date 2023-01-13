package graph

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGraph_BellmanFord(t *testing.T) {
	graph := buildGraph()
	r := graph.FindArbitrageLoop(0)
	fmt.Println(r)
}

func buildGraph() Graph {
	nodes := Nodes{
		&Node{ID: 0, UUID: "A"},
		&Node{ID: 1, UUID: "B"},
		&Node{ID: 2, UUID: "C"},
		&Node{ID: 3, UUID: "D"},
		&Node{ID: 4, UUID: "E"},
	}
	edges := Edges{
		&Edge{ID: 0, UUID: "A-B", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(-1)},
		&Edge{ID: 1, UUID: "A-C", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(4)},
		&Edge{ID: 2, UUID: "B-C", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(3)},
		&Edge{ID: 3, UUID: "B-D", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(2)},
		&Edge{ID: 4, UUID: "B-E", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(2)},
		&Edge{ID: 5, UUID: "B-E", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(5)},
		&Edge{ID: 6, UUID: "D-B", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(1)},
		&Edge{ID: 7, UUID: "E-D", Source: nodes[0], Target: nodes[1], Weight: new(big.Float).SetInt64(-3)},
	}

	return Graph{Edges: edges, Nodes: nodes}
}
