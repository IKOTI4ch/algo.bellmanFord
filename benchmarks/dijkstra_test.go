package benchmarks

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"graph/dijkstra"
	"graph/graph"
	"log"
	"testing"
	"time"
)

func TestShortestAll(t *testing.T) {
	psqlURL := "postgresql://koti@localhost:5432/safeblock.sniper_development.ruby9?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), psqlURL)

	gr := dijkstra.NewGraph()
	nodes := make(map[string]int)
	edges := make([][2]string, 0)

	maxCount := 1_000_000
	rows, err := conn.Query(context.Background(), "SELECT id, token1_id, token2_id FROM pairs WHERE network = 'bsc' LIMIT $1", maxCount)
	if err != nil {
		panic(err)
	}

	nodes, edges = graph.GetDBData(rows)
	gr.CreateData(nodes, edges)

	src := nodes["1460a9f3-e37a-455a-99e2-1a9480b8d5a1"]
	dest := nodes["b695cad1-c888-44a6-b72d-b41aacd32657"]
	limit := int64(3)

	_ = gr.AddArc(src, dest, 1) // For test
	_ = gr.AddArc(dest, src, 1) // For test

	exp := 400 * time.Millisecond // max time for complete

	for i := 0; i < 10; i++ {
		start := time.Now()
		_, _ = gr.ShortestAll(src, dest, limit)
		delta := time.Now().Sub(start)

		log.Printf("i:%d | delta:%s\n", i, delta)
		assert.LessOrEqual(t, delta, exp)
	}
}
