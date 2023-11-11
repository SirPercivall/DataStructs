package datastructs

// Graph implements a graph data structure.
type Graph interface {
	AddVertex(v *Vertex)
	AddEdge(key1, key2 any)
	RemoveEdge(key1, key2 any)
	GetVertex(key any) *Vertex
	GetVertices() []*Vertex
	AreAdjacent(key1, key2 any) bool
}

// UndirectedGraph implements an undirected graph data structure (where edges are bidirectional);
// keys are the vertices' keys, unique identifiers.
type UndirectedGraph map[any]*Vertex

// DirectedGraph implements a directed graph data structure (digraph, where edges are unidirectional);
// keys are the vertices' keys, unique identifiers.
type DirectedGraph map[any]*Vertex

// Vertex implements a vertex of a graph;
// it has a key, a value and a set of adjacent (neighboring) vertices.
type Vertex struct {
	key any
	val any
	adj []*Vertex
}

// NewUndirectedGraph returns a pointer to a new UndirectedGraph.
func NewUndirectedGraph() (g *UndirectedGraph) {
	return &UndirectedGraph{}
}

// NewDirectedGraph returns a pointer to a new DirectedGraph.
func NewDirectedGraph() (g *DirectedGraph) {
	return &DirectedGraph{}
}

// NewVertex returns a pointer to a Vertex initialized with val and key.
func NewVertex(key, val any) (v *Vertex) {
	return &Vertex{key, val, nil}
}

// AddEdge adds an edge between the vertices with keys key1 and key2.
func (g UndirectedGraph) AddEdge(key1, key2 any) {
	v1 := g[key1]
	v2 := g[key2]
	// check if vertices exist
	if v1 == nil || v2 == nil {
		panic("vertex not present")
	}
	// check if edge already exists; if not, add it
	// g.AreAdjacent(key1, key2) is always equal to g.AreAdjacent(key2, key1)
	if g.AreAdjacent(key1, key2) {
		return
	}
	v1.adj = append(v1.adj, v2)
	v2.adj = append(v2.adj, v1)
}

// AddEdge adds an edge between the vertices with keys key1 and key2.
func (g DirectedGraph) AddEdge(key1, key2 any) {
	v1 := g[key1]
	v2 := g[key2]
	// check if vertices exist
	if v1 == nil || v2 == nil {
		panic("vertex not present")
	}
	// check if edge already exists; if not, add it
	if g.AreAdjacent(key1, key2) || g.AreAdjacent(key2, key1) {
		return
	}
	v1.adj = append(v1.adj, v2)
}

// RemoveEdge removes the edge between the vertices with keys key1 and key2 (if any).
func (g UndirectedGraph) RemoveEdge(key1, key2 any) {
	v1 := g[key1]
	v2 := g[key2]
	for i, v := range v1.adj {
		if v.key == v2.key {
			v1.adj = append(v1.adj[:i], v1.adj[i+1:]...)
		}
	}
	for i, v := range v2.adj {
		if v.key == v1.key {
			v2.adj = append(v2.adj[:i], v2.adj[i+1:]...)
		}
	}
}

// RemoveEdge removes the edge between the vertices with keys key1 and key2 (if any).
func (g DirectedGraph) RemoveEdge(key1, key2 any) {
	v1 := g[key1]
	v2 := g[key2]
	for i, v := range v1.adj {
		if v.key == v2.key {
			v1.adj = append(v1.adj[:i], v1.adj[i+1:]...)
		}
	}
}

// AddVertex adds a vertex to the graph. If the key is already present, it panics.
func (g UndirectedGraph) AddVertex(v *Vertex) {
	if g[v.key] != nil {
		panic("key already present")
	}
	g[v.key] = v
}

// AddVertex adds a vertex to the graph. If the key is already present, it panics
func (g DirectedGraph) AddVertex(v *Vertex) {
	if g[v.key] != nil {
		panic("key already present")
	}
	g[v.key] = v
}

// RemoveVertex removes the vertex with key = key from the graph (and all the edges connected to it).
func (g UndirectedGraph) RemoveVertex(key any) {
	v := g[key]
	for _, adj := range v.adj {
		g.RemoveEdge(key, adj.key)
	}
	delete(g, key)
}

// RemoveVertex removes the vertex with key = key from the graph (and all the edges connected to it).
func (g DirectedGraph) RemoveVertex(key any) {
	v := g[key]
	for _, adj := range v.adj {
		g.RemoveEdge(key, adj.key)
	}
	// remove all the edges that point to the vertex (if any)
	for _, v := range g.GetVertices() {
		g.RemoveEdge(v.key, key)
	}
	delete(g, key)
}

// GetVertex returns the vertex with key = key.
func (g UndirectedGraph) GetVertex(key any) (v *Vertex) {
	return g[key]
}

// GetVertex returns the vertex with key = key.
func (g DirectedGraph) GetVertex(key any) (v *Vertex) {
	return g[key]
}

// GetVertices returns a slice of all the vertices in the graph.
func (g UndirectedGraph) GetVertices() (vertices []*Vertex) {
	for _, v := range g {
		vertices = append(vertices, v)
	}
	return vertices
}

// GetVertices returns a slice of all the vertices in the graph.
func (g DirectedGraph) GetVertices() (vertices []*Vertex) {
	for _, v := range g {
		vertices = append(vertices, v)
	}
	return vertices
}

// AreAdjacent returns true if the vertices with keys key1 and key2 are adjacent, false otherwise.
func (g UndirectedGraph) AreAdjacent(key1, key2 any) bool {
	v1 := g[key1]
	v2 := g[key2]
	for _, v := range v1.adj {
		if v.key == v2.key {
			return true
		}
	}
	return false
}

// AreAdjacent returns true if exists an edge from the vertex with key key1 to the vertex with key key2, false otherwise.
func (g DirectedGraph) AreAdjacent(key1, key2 any) bool {
	v1 := g[key1]
	v2 := g[key2]
	for _, v := range v1.adj {
		if v.key == v2.key {
			return true
		}
	}
	return false
}
