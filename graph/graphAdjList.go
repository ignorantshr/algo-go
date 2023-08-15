package graph

import "fmt"

type Vertex struct {
	val int
}

func NewVertex(val int) Vertex {
	return Vertex{val: val}
}

type graphAdjList struct {
	adjList map[Vertex][]Vertex
}

func NewGraphAdjList(edges [][2]int) *graphAdjList {
	g := new(graphAdjList)
	g.adjList = make(map[Vertex][]Vertex)
	for _, edge := range edges {
		g.AddVertex(edge[0])
		g.AddVertex(edge[1])
		g.AddEdge(edge[0], edge[1])
	}
	return g
}

func (g *graphAdjList) AddVertex(val int) {
	ver := NewVertex(val)
	if _, ok := g.adjList[ver]; ok {
		return
	}
	g.adjList[ver] = make([]Vertex, 0)
}

func (g *graphAdjList) DelVertex(val int) {
	ver := NewVertex(val)
	list := g.adjList[ver]
	delete(g.adjList, ver)
	for _, k := range list {
		delIndex := -1
		for i, tmp := range g.adjList[k] {
			if tmp.val == val {
				delIndex = i
				break
			}
		}
		if delIndex != -1 {
			g.adjList[k] = append(g.adjList[k][:delIndex], g.adjList[k][delIndex+1:]...)
		}
	}
}

func (g *graphAdjList) AddEdge(from, to int) {
	vf := NewVertex(from)
	f, ok := g.adjList[vf]
	if !ok {
		return
	}
	vt := NewVertex(to)
	t, ok := g.adjList[vt]
	if !ok {
		return
	}

	g.adjList[vf] = append(f, vt)
	g.adjList[vt] = append(t, vf)
}

func (g *graphAdjList) DelEdge(from, to int) {
	vf := NewVertex(from)
	f, ok := g.adjList[vf]
	if !ok {
		return
	}
	vt := NewVertex(to)
	t, ok := g.adjList[vt]
	if !ok {
		return
	}

	var i int
	for i = range f {
		if f[i] == vt {
			break
		}
	}
	if i == len(f) {
		return
	}
	g.adjList[vf] = append(f[:i], f[i+1:]...)

	i = 0
	for i = range t {
		if t[i] == vt {
			break
		}
	}
	if i == len(t) {
		return
	}
	g.adjList[vt] = append(t[:i], t[i+1:]...)
}

func (g *graphAdjList) Println() {
	fmt.Println("============================")
	for k, l := range g.adjList {
		fmt.Printf("%d: %v\n", k, l)
	}
}

func GraphBFS(g *graphAdjList, start Vertex) []Vertex {
	if len(g.adjList) == 0 {
		return nil
	}

	var res []Vertex
	visited := make(map[Vertex]bool)
	visited[start] = true

	queue := make([]Vertex, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		first := queue[0]
		res = append(res, first)

		queue = queue[1:]
		for _, v := range g.adjList[first] {
			if visited[v] {
				continue
			}
			queue = append(queue, v)
			visited[v] = true
		}
	}
	return res
}

func GraphDFS(g *graphAdjList, start Vertex) []Vertex {
	if len(g.adjList) == 0 {
		return nil
	}

	return graphDFS(g, make(map[Vertex]bool), start)
}

func graphDFS(g *graphAdjList, visited map[Vertex]bool, start Vertex) []Vertex {
	res := []Vertex{start}
	visited[start] = true
	for _, v := range g.adjList[start] {
		if visited[v] {
			continue
		}

		res = append(res, graphDFS(g, visited, v)...)
	}
	return res
}
