package dp

import (
	"os"
	"reflect"
	"sort"
	"testing"
)

var g *Graph

func TestMain(m *testing.M) {
	g = NewGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)

	exitCode := m.Run()

	g = nil
	os.Exit(exitCode)
}

func TestGraph_VertexCoverNaive(t *testing.T) {
	exp := []int{0, 3, 5}
	actual, err := g.VertexCoverNaive()
	if err != nil {
		t.Fatal(err)
	}
	sort.Ints(actual)
	if !reflect.DeepEqual(exp, actual) {
		t.Errorf("expected %v, but got %v instead", exp, actual)
	}
}
