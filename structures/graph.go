package structures

import "fmt"

const (
	undiscovered = -1
	discovered   = iota
	processed    = iota
)

// Implements a Graph data structure using an adjacency list.
type Graph struct {
	edges     []*Edge
	degree    []int // how many connections a given vertex has
	parent    []int // the parent of any given vertex
	status    []int // the status of a given vertex, used during traversal
	entryTime []int // the entry time of a given vertex during traversal
	exitTime  []int // the exit time of a given vertex during traversal
	directed  bool  // directed vs undirected graph
}

type Edge struct {
	y      int   // vertex point
	weight int   // for weighted graphs
	next   *Edge // the next edge in the adjacency list (similar to linked list)
}

func InitGraph(directed bool, size int) *Graph {
	edges := make([]*Edge, size)
	degree := make([]int, size)
	parent := make([]int, size)
	status := make([]int, size)
	entryTime := make([]int, size)
	exitTime := make([]int, size)

	return &Graph{edges, degree, parent, status, entryTime, exitTime, directed}
}

func (g *Graph) Insert(x, y int) {
	insertHelper(g, x, y, g.directed)
}

func (g *Graph) Print() {
	for i, _ := range g.edges {
		fmt.Println(i)
		e := g.edges[i]
		for e != nil {
			fmt.Printf(" %d", e.y)
			e = e.next
		}
		fmt.Printf("\n Degree: %d\n", g.degree[i])
	}
}

func (g *Graph) Bfs(start int) {
	g.initSearch()

	queue := []int{start}
	g.status[start] = discovered

	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		processVertexEarly(v)
		g.status[v] = processed

		p := g.edges[v]
		for p != nil {
			y := p.y
			if g.status[y] != processed || g.directed {
				processEdge(v, y)
			}
			if g.status[y] == undiscovered {
				queue = append(queue, y)
				g.status[y] = discovered
				g.parent[y] = v
			}
			p = p.next
		}
		processVertexLate(v)
	}
}

func (g *Graph) Dfs(v int) {
	g.initSearch()

	var time int
	var finished bool
	dfsHelper(g, v, time, finished)
}

func dfsHelper(g *Graph, v, time int, finished bool) {
	if finished {
		return
	}

	g.status[v] = discovered
	time++
	g.entryTime[v] = time

	processVertexEarly(v)

	p := g.edges[v]
	for p != nil {
		y := p.y

		if g.status[y] == discovered || g.directed {
			processEdge(v, y)
		}
		if g.status[y] == undiscovered {
			g.parent[y] = v
			processEdge(v, y)
			dfsHelper(g, y, time, finished)
		}

		if finished {
			return
		}

		p = p.next
	}
	processVertexLate(v)
	time++
	g.exitTime[v] = time

	g.status[v] = processed
}

func (g *Graph) ConnectedComponents() int {
	g.initSearch()

	c := 0
	for i, _ := range g.edges {
		if g.status[i] == undiscovered && g.edges[i] != nil {
			c++
			g.Bfs(i)
		}
	}

	return c
}

func (g *Graph) FindPath(start, end int) []int {
	res := []int{}

	for start != end && g.parent[end] != -1 {
		res = append([]int{end}, res...)
		end = g.parent[end]
	}
	if start == end {
		res = append([]int{start}, res...)
	}

	return res
}

func insertHelper(g *Graph, x, y int, directed bool) {
	edge := &Edge{
		y:    y,
		next: g.edges[x],
	}
	g.edges[x] = edge
	g.degree[x]++

	if !directed {
		insertHelper(g, y, x, true)
	}
}

func (g *Graph) initSearch() {
	for i, _ := range g.parent {
		g.status[i] = undiscovered
		g.parent[i] = -1
		g.entryTime[i] = -1
		g.exitTime[i] = -1
	}
}

func processVertexEarly(v int) {
	fmt.Printf("processed vertex %d\n", v)
}

func processEdge(x, y int) {
	fmt.Printf("processed edge (%d, %d)\n", x, y)
}

func processVertexLate(v int) {
	// do something
}
