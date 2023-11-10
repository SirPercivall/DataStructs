package datastructs

import "testing"

func TestNewUndirectedGraph(t *testing.T) {
	g := NewUndirectedGraph()
	if g == nil {
		t.Error("NewUndirectedGraph() returned nil")
	}
}

func TestNewDirectedGraph(t *testing.T) {
	g := NewDirectedGraph()
	if g == nil {
		t.Error("NewDirectedGraph() returned nil")
	}
}

func TestNewVertex(t *testing.T) {
	v := NewVertex(1, 1)
	if v == nil || v.val != 1 || v.key != 1 {
		t.Error("NewVertex(1, 1) returned nil or wrong value")
	}
}

func TestUndirectedGraphAddVertex(t *testing.T) {
	g := NewUndirectedGraph()
	v := NewVertex(1, 1)
	g.AddVertex(v)
	if g.GetVertex(1) == nil {
		t.Error("AddVertex(1) failed")
	}

	// test panic on duplicate vertex
	defer func() {
		if r := recover(); r == nil {
			t.Error("AddVertex(1) did not panic")
		}
	}()
	g.AddVertex(v)
}

func TestDirectedGraphAddVertex(t *testing.T) {
	g := NewDirectedGraph()
	v := NewVertex(1, 1)
	g.AddVertex(v)
	if g.GetVertex(1) == nil {
		t.Error("AddVertex(1) failed")
	}

	// test panic on duplicate vertex
	defer func() {
		if r := recover(); r == nil {
			t.Error("AddVertex(1) did not panic")
		}
	}()
	g.AddVertex(v)
}

func TestUndirectedGraphRemoveVertex(t *testing.T) {
	g := NewUndirectedGraph()
	v1 := NewVertex(1, 1)
	v2 := NewVertex(2, 2)
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.RemoveVertex(1)
	if g.GetVertex(1) != nil {
		t.Error("RemoveVertex(1) failed")
	}
	// check if the edge was removed
	if len(g.GetVertex(2).adj) != 0 {
		t.Error("RemoveVertex(1) failed")
	}
}

func TestDirectedGraphRemoveVertex(t *testing.T) {
	g := NewDirectedGraph()
	v1 := NewVertex(1, 1)
	v2 := NewVertex(2, 2)
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddEdge(1, 2)
	g.RemoveVertex(1)
	if g.GetVertex(1) != nil {
		t.Error("RemoveVertex(1) failed")
	}
	// check if the edge was removed
	if len(g.GetVertex(2).adj) != 0 {
		t.Error("RemoveVertex(1) failed")
	}
}

func TestUndirectedGraphAddEdge(t *testing.T) {
	g := NewUndirectedGraph()
	v1 := NewVertex(1, 1)
	v2 := NewVertex(2, 2)
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddEdge(1, 2)
	if len(g.GetVertex(1).adj) != 1 || len(g.GetVertex(2).adj) != 1 {
		t.Error("AddEdge(1, 2) failed")
	}
}

func TestDirectedGraphAddEdge(t *testing.T) {
	g := NewDirectedGraph()
	v1 := NewVertex(1, 1)
	v2 := NewVertex(2, 2)
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddEdge(1, 2)
	if len(g.GetVertex(1).adj) != 1 || len(g.GetVertex(2).adj) != 0 {
		t.Error("AddEdge(1, 2) failed")
	}
}

func TestUndirectedGraphRemoveEdge(t *testing.T) {
	g := NewUndirectedGraph()
	v1 := NewVertex(1, 1)
	v2 := NewVertex(2, 2)
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddEdge(1, 2)
	g.RemoveEdge(1, 2)
	if len(g.GetVertex(1).adj) != 0 || len(g.GetVertex(2).adj) != 0 {
		t.Error("RemoveEdge(1, 2) failed")
	}
}

func TestDirectedGraphRemoveEdge(t *testing.T) {
	g := NewDirectedGraph()
	v1 := NewVertex(1, 1)
	v2 := NewVertex(2, 2)
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddEdge(1, 2)
	g.RemoveEdge(1, 2)
	if len(g.GetVertex(1).adj) != 0 || len(g.GetVertex(2).adj) != 0 {
		t.Error("RemoveEdge(1, 2) failed")
	}
}
