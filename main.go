package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"graph/dijkstra"
	"graph/graph"
	"os"
	"time"
)

const MaxCount = 1_000_000

func main() {
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, token1_id, token2_id FROM pairs WHERE network = 'bsc' LIMIT $1", MaxCount)
	if err != nil {
		panic(err)
	}

	gr := dijkstra.NewGraph()
	nodes := make(map[string]int)
	edges := make([][2]string, 0, MaxCount)

	//nodes, edges = graph.GetMockData()
	nodes, edges = graph.GetDBData(rows)

	i := 0 // index of node
	for id := range nodes {
		nodes[id] = i
		gr.AddVertex(nodes[id])
		i++
	}

	for _, vertices := range edges {
		err = gr.AddArc(nodes[vertices[0]], nodes[vertices[1]], 1)
		if err != nil {
			panic(err)
		}

		err = gr.AddArc(nodes[vertices[1]], nodes[vertices[0]], 1)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("len(edges)=", len(edges))
	fmt.Println("len(nodes)=", len(nodes))

	//j := 0
	//	//for s, i := range nodes {
	//	//	if j < 10 {
	//	//		fmt.Println("s: ", s, "\ti: ", i)
	//	//	}
	//	//	j++
	//	//}
	//	//fmt.Println()
	//	//for i := 0; i < 10; i++ {
	//	//	s := edges[i]
	//	//	fmt.Println("s: ", s, "\ti: ", i)
	//	//}

	src := nodes["ee129ba5-d4e8-4e59-8183-a32c71b956c6"]
	dest := nodes["e4edf663-72c2-49a7-babc-35b37d5cab4d"]

	start := time.Now()
	best, err := gr.ShortestAll(src, dest, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nlen=%d, path: %+v\n", len(best), best)
	}

	duration := time.Now().Sub(start)
	fmt.Println("duration:", duration)

	//start = time.Now()
	//graph.ShortestPath(g,
	//	"c18baadc-deda-42ea-aa6c-7586df091a0e",
	//	"acb83973-9780-495d-a72b-7e961057ea4a",
	//)
	//duration = time.Now().Sub(start)
	//fmt.Println("path2:", duration)
}
