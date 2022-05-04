package dp

import (
	"errors"
	"fmt"
)

var EmptyGraphErr = errors.New("graph is empty")

type Graph struct {
	AdjList [][]int
}

func NewGraph(vertices int) *Graph {
	adjList := make([][]int, vertices)
	for i := range adjList {
		adjList[i] = make([]int, 0)
	}
	return &Graph{adjList}
}

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
	res := []int{}
	visited := map[int]bool{}

	if len(g.AdjList) <= 0 {
		return nil, EmptyGraphErr
	}
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
