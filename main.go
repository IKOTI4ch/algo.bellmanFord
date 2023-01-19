package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"graph/dijkstra"
	"graph/graph"
	"log"
	"time"
)

const MaxCount = 1_000_000

func main() {
	env := GetENV()
	conn := CreatePGConnection(env.PostgresURL)
	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, context.Background())

	gr := dijkstra.NewGraph()
	nodes := make(map[string]int)
	edges := make([][2]string, 0, MaxCount)

	rows, err := conn.Query(context.Background(), "SELECT id, token1_id, token2_id FROM pairs WHERE network = 'bsc' LIMIT $1", MaxCount)
	if err != nil {
		panic(err)
	}

	nodes, edges = graph.GetDBData(rows)
	gr.CreateData(nodes, edges)

	//rgConn, err := redis.DialURL(env.RedisGraphURL)
	//if err != nil {
	//	panic(err)
	//}
	//defer func(redisConn redis.Conn) {
	//	_ = redisConn.Close()
	//}(rgConn)
	//
	//redisGraph := rg.GraphNew("dev", rgConn)
	//DisplayToRedisGraph(&redisGraph, nodes, edges, 1_000)

	src := nodes["1460a9f3-e37a-455a-99e2-1a9480b8d5a1"]
	dest := nodes["b695cad1-c888-44a6-b72d-b41aacd32657"]
	_ = gr.AddArc(src, dest, 1) // For test
	_ = gr.AddArc(dest, src, 1) // For test

	log.Println("len(edges)=", len(edges))
	log.Println("len(nodes)=", len(nodes))

	start := time.Now()
	best, err := gr.ShortestAll(src, dest, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("len=%d, path:%+v\n", len(best), best)
	}

	delta := time.Now().Sub(start)
	log.Println("delta:", delta)
}
