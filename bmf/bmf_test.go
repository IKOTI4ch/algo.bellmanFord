package bmf

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"graph/graph"
	"os"
	"testing"
)

func TestFindPathWithBMF(t *testing.T) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Edge"})
	g := graph.BuildGraph()
	path := FindPathWithBMF(g, g.Nodes[0], g.Nodes[4])

	for i, edge := range path {
		table.Append([]string{fmt.Sprint(i), fmt.Sprint(edge)})
	}
	table.Render()
}
