package main

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	/*
	   A|B,7|C,9|F,20
	   Connect A with B with an edge from A "to" B of weight value 7
	   Connect A with C with an edge from A "to" C of weight value 9
	   Connect A with F with an edge from A "to" F of weight value 20

	   F|A,20|C,2|E,9
	   Connect F with A with an edge from F "to" A of weight value 20
	   Connect F with C with an edge from F "to" C of weight value 2
	   Connect F with E with an edge from F "to" E of weight value 9
	*/

	str1 := `
S|A,7|B,6
A|C,-3|T,9
B|A,8|T,-4|C,5
C|B,-2
T|S,2|C,7
`
	// G1 has negative edges
	// therefore this should return false
	G1 := ConstructDirectedGraph(str1)

	fmt.Println("Print out all vertices.")
	vcnt1 := 1
	for vtx := G1.GetVertexList().Front(); vtx != nil; vtx = vtx.Next() {
		fmt.Println(vcnt1, "ID:", vtx.Value.(*Vertex).ID, ",",
			"Timestamp(Distance):",
			vtx.Value.(*Vertex).timestamp_d)
		vcnt1++
	}
	println()

	fmt.Println("Print out all edges.")
	ecnt1 := 1
	for edge := G1.GetEdgeList().Front(); edge != nil; edge = edge.Next() {
		fmt.Println(ecnt1,
			"Source:",
			edge.Value.(*Edge).SourceVertex.ID, ",",
			"Destination:",
			edge.Value.(*Edge).DestinationVertex.ID, ",",
			"Weight:",
			edge.Value.(*Edge).Weight)
		ecnt1++
	}
	println()
	/*
	   Print out all vertices.
	   1 ID: S , Timestamp(Distance): 9999999999
	   2 ID: A , Timestamp(Distance): 9999999999
	   3 ID: B , Timestamp(Distance): 9999999999
	   4 ID: C , Timestamp(Distance): 9999999999
	   5 ID: T , Timestamp(Distance): 9999999999

	   Print out all edges.
	   1 Source: S , Destination: A , Weight: 7
	   2 Source: S , Destination: B , Weight: 6
	   3 Source: A , Destination: C , Weight: -3
	   4 Source: A , Destination: T , Weight: 9
	   5 Source: B , Destination: A , Weight: 8
	   6 Source: B , Destination: T , Weight: -4
	   7 Source: B , Destination: C , Weight: 5
	   8 Source: C , Destination: B , Weight: -2
	   9 Source: T , Destination: S , Weight: 2
	   10 Source: T , Destination: C , Weight: 7
	*/

	fmt.Println(G1.CheckBellmanFordShortestPath(G1.GetVertexByID("S")))
	// false

	println()
	fmt.Println("After: Print out all vertices.")
	vcnt1 = 1
	for vtx := G1.GetVertexList().Front(); vtx != nil; vtx = vtx.Next() {
		fmt.Println(vcnt1, "ID:", vtx.Value.(*Vertex).ID, ",",
			"Timestamp(Distance):",
			vtx.Value.(*Vertex).timestamp_d)
		vcnt1++
	}
	println()

	fmt.Println("After: Print out all edges.")
	ecnt1 = 1
	for edge := G1.GetEdgeList().Front(); edge != nil; edge = edge.Next() {
		fmt.Println(ecnt1,
			"Source:",
			edge.Value.(*Edge).SourceVertex.ID, ",",
			"Destination:",
			edge.Value.(*Edge).DestinationVertex.ID, ",",
			"Weight:",
			edge.Value.(*Edge).Weight)
		ecnt1++
	}
	println()
	/*
	   After: Print out all vertices.
	   1 ID: S , Timestamp(Distance): 0
	   2 ID: A , Timestamp(Distance): 7
	   3 ID: B , Timestamp(Distance): 2
	   4 ID: C , Timestamp(Distance): 4
	   5 ID: T , Timestamp(Distance): -2

	   After: Print out all edges.
	   1 Source: S , Destination: A , Weight: 7
	   2 Source: S , Destination: B , Weight: 6
	   3 Source: A , Destination: C , Weight: -3
	   4 Source: A , Destination: T , Weight: 9
	   5 Source: B , Destination: A , Weight: 8
	   6 Source: B , Destination: T , Weight: -4
	   7 Source: B , Destination: C , Weight: 5
	   8 Source: C , Destination: B , Weight: -2
	   9 Source: T , Destination: S , Weight: 2
	   10 Source: T , Destination: C , Weight: 7
	*/

	fmt.Println("Prev")
	for _, targetprev := range G1.GetVertexByID("T").Prev {
		fmt.Printf("%v ", targetprev.ID)
	}
	// A B

	println()
	fmt.Println("-----------------------------")

	str2 := `
S|A,11|B,17|C,9
A|S,11|B,5|D,50|T,500
B|S,17|D,30
C|S,9
D|A,50|B,30|E,3|F,11
E|B,18|C,27|D,3|T,19
F|D,11|E,6|T,77
T|A,500|D,10|F,77|E,19
`
	// G2 has no negative edges
	// therefore this should return true
	G2 := ConstructDirectedGraph(str2)
	fmt.Println(G2.CheckBellmanFordShortestPath(G2.GetVertexByID("S")))
	// true

	fmt.Println("Prev")
	for _, targetprev := range G2.GetVertexByID("T").Prev {
		fmt.Printf("%v ", targetprev.ID)
	}
	// A E F
	println()

	fmt.Println("Print out all vertices.")
	vcnt2 := 1
	for vtx := G2.GetVertexList().Front(); vtx != nil; vtx = vtx.Next() {
		fmt.Println(vcnt2, "ID:", vtx.Value.(*Vertex).ID, ",",
			"Timestamp(Distance):",
			vtx.Value.(*Vertex).timestamp_d)
		vcnt2++
	}
	println()
	/*
	   Print out all vertices.
	   1 ID: S , Timestamp(Distance): 0
	   2 ID: A , Timestamp(Distance): 11
	   3 ID: B , Timestamp(Distance): 16
	   4 ID: C , Timestamp(Distance): 9
	   5 ID: D , Timestamp(Distance): 46
	   6 ID: T , Timestamp(Distance): 68
	   7 ID: E , Timestamp(Distance): 49
	   8 ID: F , Timestamp(Distance): 57
	*/
}

/////////// BellmanFord ///////////
/*
	CheckBellmanFordShortestPath(G, source)
		// Initialize-Single-Source(G,s)
		for each vertex v ∈ G.V
			v.d = ∞
			v.π = nil
		source.d = 0

		// for each vertex
		for  i = 1  to  |G.V| - 1
			for  each edge (u, v) ∈ G.E
				Relax(u, v, w)

		for  each edge (u, v) ∈ G.E
			if  v.d > u.d + w(u, v)
				if v.d > u.d + w(u,v)
					return FALSE

		return TRUE
*/
// Return true if there is any negative edge.
func (G *Graph) CheckBellmanFordShortestPath(source *Vertex) bool {

	// for each vertex u ∈ G.V
	for vtx := G.GetVertexList().Front(); vtx != nil; vtx = vtx.Next() {
		vtx.Value.(*Vertex).timestamp_d = 9999999999
		vtx.Value.(*Vertex).Predecessor.Init()
	}
	source.timestamp_d = 0

	for v := G.GetVertexList().Front(); v != nil; v = v.Next() {
		// Relax
		for edge := G.GetEdgeList().Front(); edge != nil; edge = edge.Next() {
			a := edge.Value.(*Edge).DestinationVertex.timestamp_d
			b := edge.Value.(*Edge).SourceVertex.timestamp_d
			c := G.GetEdgeWeight(edge.Value.(*Edge).SourceVertex, edge.Value.(*Edge).DestinationVertex)

			if a > b+c {
				edge.Value.(*Edge).DestinationVertex.timestamp_d = b + c
			}

			edge.Value.(*Edge).DestinationVertex.Predecessor.Init()
			edge.Value.(*Edge).DestinationVertex.Predecessor.PushBack(edge.Value.(*Edge).SourceVertex)

			if len(edge.Value.(*Edge).DestinationVertex.Prev) == 0 {
				edge.Value.(*Edge).DestinationVertex.Prev = append(edge.Value.(*Edge).DestinationVertex.Prev, edge.Value.(*Edge).SourceVertex)
				continue
			} else {
				exist := false
				for _, targetprev := range edge.Value.(*Edge).DestinationVertex.Prev {
					if targetprev == edge.Value.(*Edge).SourceVertex {
						exist = true
					}
				}

				if exist == false {
					edge.Value.(*Edge).DestinationVertex.Prev = append(edge.Value.(*Edge).DestinationVertex.Prev, edge.Value.(*Edge).SourceVertex)
				}
			}
		}
	}

	for edge := G.EdgeList.Front(); edge != nil; edge = edge.Next() {
		a := edge.Value.(*Edge).DestinationVertex.timestamp_d
		b := edge.Value.(*Edge).SourceVertex.timestamp_d
		c := G.GetEdgeWeight(edge.Value.(*Edge).SourceVertex, edge.Value.(*Edge).DestinationVertex)

		if a > b+c {
			return false
		}
	}
	return true
}

/////////////////////////////////

// ///////// Graph ///////////
type Graph struct {
	VertexList *list.List
	EdgeList   *list.List
}

// Construct a new bmf and return it.
func NewGraph() *Graph {
	return &Graph{
		list.New(),
		list.New(),
	}
}

type Vertex struct {
	ID    string
	Color string

	// Vertex's outgoing edges
	EdgesFromThisVertex *list.List

	// vertices with edges that goes into Vertex
	// Vertex's incoming vertices
	Predecessor *list.List

	// distance from source vertex
	timestamp_d int64

	// another timestamp to be used in another algorithm
	timestamp_f int64

	// for BellmanFord's shortest path
	Prev []*Vertex
}

// Construct a new vertex and return it.
func NewVertex(input_id string) *Vertex {
	return &Vertex{
		ID:                  input_id,
		Color:               "white",
		EdgesFromThisVertex: list.New(),
		Predecessor:         list.New(),
		timestamp_d:         9999999999,
		timestamp_f:         9999999999,
		Prev:                nil,
	}
}

type Edge struct {
	SourceVertex      *Vertex
	DestinationVertex *Vertex
	Weight            int64
}

// Construct a new edge from source to destination vertex.
// Not from destination to source vertex.
func NewEdge(source, destination *Vertex, weight int64) *Edge {
	return &Edge{
		source,
		destination,
		weight,
	}
}

// Connect the vertex A with edges.
func (A *Vertex) ConnectEdgeWithVertex(edges ...*Edge) {
	for _, edge := range edges {
		A.EdgesFromThisVertex.PushBack(edge)
	}
}

// Return adjacent edges that comes "out of" vertex A.
func (A *Vertex) GetEdgesChannelFromThisVertex() chan *Edge {
	edgechan := make(chan *Edge)

	go func() {
		defer close(edgechan)
		for e := A.EdgesFromThisVertex.Front(); e != nil; e = e.Next() {
			edgechan <- e.Value.(*Edge)
		}
	}()
	return edgechan
}

/*
Comments on this pattern from Alan Donovan

While tempting, it's not idiomatic Go style to use channels
simply for the ability to iterate over them.

It's not efficient
, and it can easily lead to an accumulation of idle goroutines:

consider what happens when the caller of GetEdgesChannelFromThisVertex
discards the channel before reading to the end.

---

It's better to use container/list rather than channel

Return an adjacent edge list that comes "out of" vertex A.
*/
func (A *Vertex) GetEdgeListFromThisVertex() *list.List {
	return A.EdgesFromThisVertex
}

// Return adjacent vertices from a vertex.
func (A *Vertex) GetAdjacentVertices() *list.List {
	result := list.New()
	for edge := A.GetEdgeListFromThisVertex().Front(); edge != nil; edge = edge.Next() {
		result.PushBack(edge.Value.(*Edge).DestinationVertex)
	}
	return result
}

// Construct a bmf from input string data.
func ConstructDirectedGraph(input_str string) *Graph {
	var validID = regexp.MustCompile(`\t{1,}`)
	newstr := validID.ReplaceAllString(input_str, " ")
	newstr = strings.TrimSpace(newstr)
	lines := strings.Split(newstr, "\n")

	new_graph := NewGraph()

	for _, line := range lines {
		fields := strings.Split(line, "|")

		// SourceID in string format
		SourceID := fields[0]
		edgepairs := fields[1:]

		new_graph.FindOrConstruct(SourceID)

		for _, pair := range edgepairs {
			if len(strings.Split(pair, ",")) == 1 {
				// to skip the lines below
				// and go back to the for-loop
				continue
			}
			DestinationID := strings.Split(pair, ",")[0]
			weight := StrToInt64(strings.Split(pair, ",")[1])

			src_vertex := new_graph.FindOrConstruct(SourceID)
			des_vertex := new_graph.FindOrConstruct(DestinationID)

			// This is not constructing the bi-directional edge automatically.
			// We need to input bi-directional bmf data.
			edge := NewEdge(src_vertex, des_vertex, weight)
			src_vertex.ConnectEdgeWithVertex(edge)
			des_vertex.Predecessor.PushBack(src_vertex)

			new_graph.EdgeList.PushBack(edge)
		}
	}
	return new_graph
}

// Convert string to integer.
func StrToInt64(input_str string) int64 {
	result, err := strconv.Atoi(input_str)
	if err != nil {
		panic("failed to convert string")
	}
	return int64(result)
}

// Return the vertex with input ID.
func (G *Graph) GetVertexByID(id string) *Vertex {
	for vtx := G.VertexList.Front(); vtx != nil; vtx = vtx.Next() {
		// NOT  vtx.Value.(Vertex).ID
		if vtx.Value.(*Vertex).ID == id {
			return vtx.Value.(*Vertex)
		}
	}
	return nil
}

// Find the node with the ID, or create it.
func (G *Graph) FindOrConstruct(id string) *Vertex {
	vertex := G.GetVertexByID(id)
	if vertex == nil {
		vertex = NewVertex(id)

		// then add this vertex to the bmf
		G.VertexList.PushBack(vertex)
	}
	return vertex
}

// Return the vertex list.
func (G *Graph) GetVertexList() *list.List {
	return G.VertexList
}

// Return the edge list.
func (G *Graph) GetEdgeList() *list.List {
	return G.EdgeList
}

// Delete a Vertex from the bmf.
func (G *Graph) DeleteVertex(A *Vertex) {
	for vtx := G.VertexList.Front(); vtx != nil; vtx = vtx.Next() {
		if vtx.Value.(*Vertex) == A {
			// remove from the bmf
			G.VertexList.Remove(vtx)
		}
	}

	// traverse all the outgoing edge(vertex)
	// remove this vertex A from the predecessor of the vertices,
	// also delete the edges from A
	for edge := A.GetEdgeListFromThisVertex().Front(); edge != nil; edge = edge.Next() {
		G.DeleteEdgeFrom(A, edge.Value.(*Edge).DestinationVertex)
		for vtx := edge.Value.(*Edge).DestinationVertex.Predecessor.Front(); vtx != nil; vtx = vtx.Next() {
			edge.Value.(*Edge).DestinationVertex.Predecessor.Remove(vtx)
		}
	}
}

// Delete the edge from the vertex A to B.
func (G *Graph) DeleteEdgeFrom(A, B *Vertex) {

	for edge := G.EdgeList.Front(); edge != nil; edge = edge.Next() {

		// if the edge is from A to B
		if edge.Value.(*Edge).SourceVertex == A && edge.Value.(*Edge).DestinationVertex == B {

			// don't do this
			// edge.Value.(*Edge).SourceVertex = nil
			// edge.Value.(*Edge).DestinationVertex = nil

			// remove this edge from the bmf's edge list
			G.EdgeList.Remove(edge)

			// remove this edge from the vertex's predecessor list
			for vtx := edge.Value.(*Edge).DestinationVertex.Predecessor.Front(); vtx != nil; vtx = vtx.Next() {
				if vtx.Value.(*Vertex) == A {
					edge.Value.(*Edge).DestinationVertex.Predecessor.Remove(vtx)
				}
			}
		}
	}
}

// Return the size of vertex list in a bmf.
func (G *Graph) GetVertexSize() int {
	return G.VertexList.Len()
}

// Return the size of edge list in a bmf.
func (G *Graph) GetEdgeSize() int {
	return G.EdgeList.Len()
}

// Return the weight value from source to target.
func (G *Graph) GetEdgeWeight(source, target *Vertex) int64 {
	for edge := source.GetEdgeListFromThisVertex().Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).DestinationVertex == target {
			return edge.Value.(*Edge).Weight
		} else {
			continue
		}
	}
	return 0
}
