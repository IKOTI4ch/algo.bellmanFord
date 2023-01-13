package graph

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math/big"
	"os"
)

type Edge struct {
	Source   int
	Target   int
	UUID     string
	Weight   *big.Float
	Distance *big.Float
}

type Edges []*Edge

// NewEdge returns a pointer to a new Edge
func NewEdge(UUID string, source, target int, weight *big.Float) *Edge {
	return &Edge{UUID: UUID, Source: source, Target: target, Weight: weight}
}

func Logger(edges Edges) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"i", "UUID", "Source", "Target", "Weight", "Distance"})

	fmt.Printf("Bellman–Ford Algorithm\n")
	fmt.Println(edges)
	for i, edge := range edges {
		if edge.Distance == nil {
			table.Append([]string{
				fmt.Sprint(i),
				edge.UUID,
				fmt.Sprint(edge.Source),
				fmt.Sprint(edge.Target),
				fmt.Sprint(edge.Weight),
				"Inf",
			})
			continue
		}

		table.Append([]string{
			fmt.Sprint(i),
			edge.UUID,
			fmt.Sprint(edge.Source),
			fmt.Sprint(edge.Target),
			fmt.Sprint(edge.Weight),
			edge.Weight.String(),
		})
	}
	table.Render()
}
