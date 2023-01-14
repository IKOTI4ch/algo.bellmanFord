package bmf

import (
	"graph/graph"
	"math"
	"math/big"
)

var Count int64

func BellmanFord(g *graph.Graph, source *graph.Node) (map[string]*graph.Edge, []*big.Float) {
	seen := make(map[string]*graph.Edge, len(g.Nodes)-1)
	dist := make([]*big.Float, len(g.Nodes))

	inf := new(big.Float).SetInf(false)
	for _, node := range g.Nodes {
		dist[node.Index] = inf
		Count++
	}

	dist[source.Index] = new(big.Float)

	for i := 1; i < len(g.Nodes); i++ {
		for _, edge := range g.Edges {
			s := edge.Source
			t := edge.Target
			weight := edge.Weight
			Count++

			if dist[s.Index].Cmp(inf) != 0 && new(big.Float).Add(dist[s.Index], weight).Cmp(dist[t.Index]) < 0 {
				dist[t.Index] = new(big.Float).Add(dist[s.Index], weight)
				edge.Distance = dist[t.Index]
				// sees[edge.Target] = edge.Source
				// Seen - последний результирующий слайс
				seen[t.UUID] = edge
			}
		}
	}

	return seen, dist
}

func FindPathWithBMF(g *graph.Graph, source, target *graph.Node) graph.Edges {
	seen, _ := BellmanFord(g, source)
	return FindPath(seen, source, target)
}

func FindPath(seen map[string]*graph.Edge, source, target *graph.Node) graph.Edges {
	path := graph.Edges{}

	curr := target
	for curr != source {
		path = append(path, seen[curr.UUID])
		curr = seen[curr.UUID].Source
		Count++
	}

	size := len(path)
	for i := 0; float64(i) < math.Floor(float64((size-1)/2)) || (size == 2 && i == 0); i++ {
		path[i], path[size-1-i] = path[size-1-i], path[i]
		Count++
	}

	return path
}
