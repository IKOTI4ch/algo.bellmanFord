package graph

import (
	"math/big"
)

type Edge struct {
	Distance *big.Float
	Index    int
	Source   *Node
	Target   *Node
	UUID     string
	Weight   *big.Float
}

type Edges []*Edge
