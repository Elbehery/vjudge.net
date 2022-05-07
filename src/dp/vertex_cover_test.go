package dp

import (
	"reflect"
	"sort"
	"testing"
)

func TestGraph_VertexCoverNaive(t *testing.T) {
	testCases := []struct {
		name   string
		input  *Graph
		expErr error
		exp    []int
	}{
		{
			"graphOne",
			testGraph_One(),
			nil,
			[]int{0, 3, 5},
		},
		{
			"graphTwo",
			testGraph_Two(),
			nil,
			[]int{0, 4},
		},
		{
			"graphThree",
			testGraph_Three(),
			nil,
			[]int{0, 4},
		},
		{
			"emptyGraph",
			testGraph_Empty(),
			EmptyGraphErr,
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.input.VertexCoverNaive()
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("expected error %v, but got nil instead", err)
				}
				if err != tc.expErr {
					t.Errorf("expected error %v, but got %v instead", tc.expErr, err)
				}
				if actual != nil {
					t.Errorf("expected nil, but got %v instead", actual)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil error, but got %v instead", err)
			}
			sort.Ints(actual)
			if !reflect.DeepEqual(actual, tc.exp) {
				t.Errorf("expected %v, but got %v instead", tc.exp, actual)
			}
		})
	}
}

func TestGraph_IsTree(t *testing.T) {
	testCases := []struct {
		name   string
		input  *Graph
		expErr error
		exp    bool
	}{
		{
			"graphOne",
			testGraph_One(),
			nil,
			true,
		},
		{
			"graphTwo",
			testGraph_Two(),
			nil,
			true,
		},
		{
			"graphThree",
			testGraph_Three(),
			nil,
			false,
		},
		{
			"emptyGraph",
			testGraph_Empty(),
			EmptyGraphErr,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.input.IsTree()
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("expected err %v, but got nil instead", tc.expErr)
				}
				if err != tc.expErr {
					t.Errorf("expected err %v, but got %v instead", tc.expErr, err)
				}
				if actual != false {
					t.Errorf("expected false, but got %v instead", actual)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil error, but got %v instead", err)
			}
			if tc.exp != actual {
				t.Errorf("expected %v, but got %v instead", tc.exp, actual)
			}
		})
	}
}

func testGraph_One() *Graph {
	g := NewGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)

	return g
}

func testGraph_Two() *Graph {
	g := NewGraph(5)
	g.AddEdge(2, 1)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(4, 5)

	return g
}

func testGraph_Three() *Graph {
	g := NewGraph(5)
	g.AddEdge(2, 1)
	g.AddEdge(1, 3)
	g.AddEdge(3, 2)
	g.AddEdge(1, 4)
	g.AddEdge(4, 5)

	return g
}

func testGraph_Empty() *Graph {
	return NewGraph(3)
}
