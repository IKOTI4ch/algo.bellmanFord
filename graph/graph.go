package graph

import (
	"math/big"
)

type Graph struct {
	Edges         Edges
	Nodes         Nodes
	TokenIdToName map[int]string //  0 -> eth, 1 -> wbtc
	TokenNameToId map[string]int //	 eth -> 0, wbtc -> 1
}

func NewGraph(
	edges []*Edge,
	nodes []int,
	idToName map[int]string,
	nameToId map[string]int) *Graph {

	return &Graph{
		Edges:         edges,
		Nodes:         nodes,
		TokenIdToName: idToName,
		TokenNameToId: nameToId,
	}
}

func GetMockGraph() *Graph {
	var g *Graph
	n1 := "A"
	n2 := "B"
	n3 := "C"
	n4 := "D"
	n5 := "E"

	nameToId := map[string]int{
		n1: 0,
		n2: 1,
		n3: 2,
		n4: 3,
		n5: 4,
	}
	idToName := map[int]string{
		nameToId[n1]: n1,
		nameToId[n2]: n2,
		nameToId[n3]: n3,
		nameToId[n4]: n4,
		nameToId[n5]: n5,
	}
	n := []int{
		nameToId[n1],
		nameToId[n2],
		nameToId[n3],
		nameToId[n4],
		nameToId[n5],
	}
	e := Edges{
		// add edge 0-1 (or A-B in above figure)
		NewEdge("A-B", nameToId[n1], nameToId[n2], new(big.Float).SetInt64(-1)),
		// add edge 0-2 (or A-C in above figure)
		NewEdge("A-C", nameToId[n1], nameToId[n3], new(big.Float).SetInt64(4)),
		// add edge 1-2 (or B-C in above figure)
		NewEdge("B-C", nameToId[n2], nameToId[n3], new(big.Float).SetInt64(3)),
		// add edge 1-3 (or B-D in above figure)
		NewEdge("B-D", nameToId[n2], nameToId[n4], new(big.Float).SetInt64(2)),
		// add edge 1-4 (or B-E in above figure)
		NewEdge("B-E", nameToId[n2], nameToId[n5], new(big.Float).SetInt64(2)),
		// add edge 3-2 (or D-C in above figure)
		NewEdge("B-E", nameToId[n4], nameToId[n3], new(big.Float).SetInt64(5)),
		// add edge 3-1 (or D-B in above figure)
		NewEdge("D-B", nameToId[n4], nameToId[n2], new(big.Float).SetInt64(1)),
		// add edge 4-3 (or E-D in above figure)
		NewEdge("E-D", nameToId[n5], nameToId[n4], new(big.Float).SetInt64(-3)),
	}

	g = NewGraph(e, n, idToName, nameToId)

	return g
}

func (g *Graph) GetTokenName(id int) string {
	return g.TokenIdToName[id]
}

func (g *Graph) GetTokenId(name string) int {
	return g.TokenNameToId[name]
}

func (g *Graph) BellmanFord(source int) ([]int, []*big.Float) {
	size := len(g.Nodes)
	edges := g.Edges
	predecessors := make([]int, size)
	distances := make([]*big.Float, size)

	inf := new(big.Float).SetInf(false)
	for _, node := range g.Nodes {
		distances[node] = inf
	}

	distances[source] = new(big.Float)

	for i := 1; i < size; i++ {
		for _, edge := range edges {
			s := edge.Source
			t := edge.Target
			weight := edge.Weight

			if distances[s].Cmp(inf) != 0 && new(big.Float).Add(distances[s], weight).Cmp(distances[t]) < 0 {
				predecessors[t] = s
				distances[t] = new(big.Float).Add(distances[s], weight)
			}
		}
	}

	for _, edge := range edges {
		s := edge.Source
		t := edge.Target
		weight := edge.Weight

		if distances[s].Cmp(inf) != 0 && new(big.Float).Add(distances[s], weight).Cmp(distances[t]) < 0 {
			return []int{}, distances
		}
	}

	return predecessors, distances
}

func (g *Graph) FindArbitrageLoop(source int) []int {
	pred, dist := g.BellmanFord(source)
	return g.FindNegativeWeightCycle(pred, dist, source)
}

func (g *Graph) FindNegativeWeightCycle(pred []int, dist []*big.Float, source int) []int {
	for _, edge := range g.Edges {
		tmpDist := new(big.Float).Add(dist[edge.Source], edge.Weight)
		if tmpDist.Cmp(dist[edge.Target]) < -1 {
			return arbitrageLoop(pred, source)
		}
	}

	return nil
}

func arbitrageLoop(predecessors []int, source int) []int {
	size := len(predecessors)
	loop := make([]int, size)
	loop[0] = source

	exists := make([]bool, size)
	exists[source] = true

	indices := make([]int, size)

	var index, next int
	for index, next = 1, source; ; index++ {
		next = predecessors[next]
		loop[index] = next
		if exists[next] {
			return loop[indices[next] : index+1]
		}
		indices[next] = index
		exists[next] = true
	}
}
