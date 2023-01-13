package graph

import (
	"fmt"
	"math/big"
)

type Graph struct {
	Edges Edges
	Nodes Nodes
}

func (g *Graph) BellmanFord(source int) ([]*Edge, []*big.Float) {
	seen := make([]*Edge, len(g.Nodes)-1)
	dist := make([]*big.Float, len(g.Nodes))

	inf := new(big.Float).SetInf(false)
	for _, node := range g.Nodes {
		dist[node.ID] = inf
	}

	dist[source] = new(big.Float)

	for i := 1; i < len(g.Nodes); i++ {
		for _, edge := range g.Edges {
			s := edge.Source
			t := edge.Target
			weight := edge.Weight

			if dist[s.ID].Cmp(inf) != 0 && new(big.Float).Add(dist[s.ID], weight).Cmp(dist[t.ID]) < 0 {
				dist[t.ID] = new(big.Float).Add(dist[s.ID], weight)
				edge.Distance = dist[t.ID]
				// sees[edge.Target] = edge.Source
				// Seen - последний результирующий слайс
				seen[t.ID-1] = edge
			}
		}
	}

	return seen, dist
}

func (g *Graph) FindArbitrageLoop(source int) []int { // deprecated
	pred, dist := g.BellmanFord(source)
	return g.FindNegativeWeightCycle(pred, dist, source)
}

func (g *Graph) FindNegativeWeightCycle(seen []*Edge, dist []*big.Float, source int) []int {
	for _, edge := range seen {
		fmt.Println(edge)
		weight := edge.Weight
		if new(big.Float).Add(dist[edge.Source.ID], weight).Cmp(dist[edge.Target.ID]) < 0 {
			//return arbitrageLoop(seen, source)
		}
	}

	return nil
}

//func arbitrageLoop(predecessors []int, source int) []int { // deprecated
//	size := len(predecessors)
//	loop := make([]int, size)
//	loop[0] = source
//
//	exists := make([]bool, size)
//	exists[source] = true
//
//	indices := make([]int, size)
//
//	var index, next int
//	for index, next = 1, source; ; index++ {
//		next = predecessors[next]
//		loop[index] = next
//		if exists[next] {
//			return loop[indices[next] : index+1]
//		}
//		indices[next] = index
//		exists[next] = true
//	}
//}
