package graph

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/jackc/pgx/v5"
)

type DBPair struct {
	ID       string
	Token1ID string
	Token2ID string
}

const MaxCount = 1_000_000

func GetDBData(rows pgx.Rows) (nodes map[string]int, edges [][2]string) {
	nodes = map[string]int{}
	edges = [][2]string{}

	bar := pb.StartNew(MaxCount)
	for rows.Next() {
		var dbPair DBPair

		err := rows.Scan(&dbPair.ID, &dbPair.Token1ID, &dbPair.Token2ID)
		if err != nil {
			panic(err)
		}

		if nodes[dbPair.Token1ID] == 0 || nodes[dbPair.Token2ID] == 0 {
			nodes[dbPair.Token1ID] = -1
			nodes[dbPair.Token2ID] = -1
			edges = append(edges, [2]string{dbPair.Token1ID, dbPair.Token2ID})
		}

		bar.Increment()
	}
	bar.Finish()

	// return nodes without index(all the index=0)
	return nodes, edges
}
