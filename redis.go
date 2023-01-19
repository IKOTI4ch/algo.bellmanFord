package main

import (
	"fmt"
	rg "github.com/thinkdata-works/redisgraph-go"
)

func CreateEdgeWithNodes(rg *rg.Graph, node1, node2, edgeID string) (*rg.QueryResult, error) {
	query := fmt.Sprintf(`
				MERGE (t1:Token{uuid: '%[1]s'})
				MERGE (t2:Token{uuid: '%[2]s'})
				MERGE (t1)-[:pair{uuid: '%[3]s'}]->(t2)
			`, node1, node2, edgeID)
	return rg.Query(query)
}

func DisplayToRedisGraph(rg *rg.Graph, nodes map[string]int, edges [][2]string, limit int) {
	for i := 0; i < limit; i++ {
		edge := edges[i]
		id := fmt.Sprintf("%d-%d", nodes[edge[0]], nodes[edge[1]])
		_, err := CreateEdgeWithNodes(rg, edge[0], edge[1], id)
		if err != nil {
			panic(err)
		}
	}
}
