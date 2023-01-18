package main

import (
	"fmt"
	"graph/dijkstra"
	"time"
)

type DBPair struct {
	ID       string
	Token1ID string
	Token2ID string
}

const MaxCount = 1_000_000

func main() {
	//conn, err := pgx.Connect(context.Background(), "postgresql://safeblock:safeblock@localhost:5432/safeblock.sniper_development.ruby8?sslmode=disable")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close(context.Background())
	//
	//rows, err := conn.Query(context.Background(), "SELECT id, token1_id, token2_id FROM pairs WHERE network = 'bsc' LIMIT $1", MaxCount)
	//if err != nil {
	//	panic(err)
	//}

	//bar := pb.StartNew(MaxCount)
	gr := dijkstra.NewGraph()

	nodes := make(map[string]int)           // uuid, i++
	edges := make([][2]string, 0, MaxCount) // [[uuid, uuid]]

	nodes["A"] = 0
	nodes["B"] = 1
	nodes["C"] = 2
	nodes["D"] = 3
	nodes["E"] = 4

	edges = append(
		edges,
		[2]string{"A", "B"},
		[2]string{"A", "C"},
		[2]string{"B", "C"},
		[2]string{"B", "D"},
		[2]string{"B", "E"},
		[2]string{"C", "D"},
		[2]string{"E", "D"},
		[2]string{"D", "A"}, //For dev
	)

	//for rows.Next() {
	//	var dbPair DBPair
	//
	//	err = rows.Scan(&dbPair.ID, &dbPair.Token1ID, &dbPair.Token2ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	nodes[dbPair.Token1ID] = 0
	//	nodes[dbPair.Token2ID] = 0
	//	edges = append(edges, [2]string{dbPair.Token1ID, dbPair.Token2ID})
	//
	//	bar.Increment()
	//}

	//var i int
	for _, val := range nodes {
		//i++
		//nodes[uuid] = i
		gr.AddVertex(val)
	}
	for _, ints := range edges {
		err := gr.AddArc(nodes[ints[0]], nodes[ints[1]], 1)
		if err != nil {
			panic(err)
		}
		err = gr.AddArc(nodes[ints[1]], nodes[ints[0]], 1)
		if err != nil {
			panic(err)
		}
	}
	//bar.Finish()

	fmt.Println("Edges: ", len(edges))
	fmt.Println("Nodes: ", len(nodes))

	start := time.Now()
	best, err := gr.ShortestAll(0, 3, 3)
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
