package structures

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	g := InitGraph(false, 10)

	g.Insert(1, 6)
	g.Insert(1, 5)
	g.Insert(1, 2)
	g.Insert(2, 3)
	g.Insert(2, 5)
	g.Insert(5, 4)
	g.Insert(3, 4)

	g.Print()

	g.Bfs(1)

	p := g.FindPath(1, 4)
	expected := []int{1, 5, 4}

	if ok := reflect.DeepEqual(p, expected); !ok {
		t.Errorf("Expected %v, got %v", expected, p)
	}

	g.Dfs(1)

	if c := g.ConnectedComponents(); c != 1 {
		t.Errorf("Expected 1, got %d", c)
	}

	g = InitGraph(true, 10)

	g.Insert(1, 6)
	g.Insert(1, 5)
	g.Insert(1, 2)
	g.Insert(2, 3)
	g.Insert(2, 5)
	g.Insert(5, 4)
	g.Insert(3, 4)

	g.Print()

	g.Bfs(1)
	g.Dfs(1)

	g = InitGraph(false, 10)

	g.Insert(1, 2)
	g.Insert(4, 5)

	if c := g.ConnectedComponents(); c != 2 {
		t.Errorf("Expected 2, got %d", c)
	}
}
