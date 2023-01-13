package main

import (
	"fmt"
	"graph/graph"
)

const source = "A"

func main() {
	g := graph.GetMockGraph()
	loop := g.FindArbitrageLoop(g.GetTokenId(source))

	for _, key := range loop {
		fmt.Printf("%v -> ", g.GetTokenName(key))
	}
}
