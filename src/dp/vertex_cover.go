package dp

import (
	"errors"
	"fmt"
)

var EmptyGraphErr = errors.New("graph is empty")

type Graph struct {
	AdjList [][]int
}

// NewGraph initialize the graph as disconnected
func NewGraph(vertices int) *Graph {
	adjList := make([][]int, vertices)
	for i := range adjList {
		adjList[i] = make([]int, 0)
	}
	return &Graph{adjList}
}

// AddEdge created undirected edge between u and w
func (g *Graph) AddEdge(u, w int) error {
	if u <= 0 || u > len(g.AdjList) {
		return fmt.Errorf("vertix %v can not be less than 1 or greater than %v", u, len(g.AdjList))
	}

	if w <= 0 || w > len(g.AdjList) {
		return fmt.Errorf("vertix %v can not be less than 1 or greater than %v", w, len(g.AdjList))
	}

	g.AdjList[u-1] = append(g.AdjList[u-1], w-1)
	g.AdjList[w-1] = append(g.AdjList[w-1], u-1)
	return nil
}

func (g *Graph) VertexCoverNaive() ([]int, error) {
	for i := range g.AdjList {
		if len(g.AdjList[i]) <= 0 {
			return nil, EmptyGraphErr
		}
	}

	res := []int{}
	visited := map[int]bool{}

	// visit the graph and create the result
	for i, l := range g.AdjList {
		if visited[i] {
			continue
		}
		res = append(res, i)
		visited[i] = true
		for _, v := range l {
			visited[v] = true
		}
	}
	return res, nil
}

func (g *Graph) IsTree() (bool, error) {
	for i := range g.AdjList {
		if len(g.AdjList[i]) <= 0 {
			return false, EmptyGraphErr
		}
	}

	visited := map[int]bool{}
	if g.isCyclic(0, -1, visited) {
		return false, nil
	}

	for i := range g.AdjList {
		if _, ok := visited[i]; !ok {
			return false, nil
		}
	}

	return true, nil
}

func (g *Graph) isCyclic(curVertex, parentVertex int, visited map[int]bool) bool {
	// visit the current vertex
	visited[curVertex] = true

	// iterate over neighbours in DFS and check if subgraph contains cycle
	for _, neighbour := range g.AdjList[curVertex] {
		_, ok := visited[neighbour]
		if ok && neighbour != parentVertex {
			return true
		} else if !ok {
			if g.isCyclic(neighbour, curVertex, visited) {
				return true
			}
		}
	}
	return false
}
