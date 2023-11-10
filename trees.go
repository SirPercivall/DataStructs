package datastructs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TreeNode implements a node of a generic tree
type TreeNode struct {
	firstChild  *TreeNode
	nextBrother *TreeNode
	val         any
}

// NewNode returns a pointer to a TreeNode initialized with val.
func NewNode(val any) (n *TreeNode) {
	return &TreeNode{nil, nil, val}
}

// GetVal returns the value of the node.
func (node *TreeNode) GetVal() (val any) {
	return node.val
}

// SetVal set the value of the nodes to the value passed as input.
func (node *TreeNode) SetVal(val any) {
	node.val = val
}

// AddChild add passed child to the node.
func (node *TreeNode) AddChild(other *TreeNode) {
	child := node.firstChild
	if child != nil {
		for child.nextBrother != nil {
			child = child.nextBrother
		}
		child.nextBrother = other
	} else {
		node.firstChild = other
	}
}

// Tree implement a tree data types.
type Tree struct {
	root *TreeNode
}

// TraversePreorder perform a deep-first research and returns a slice of *TreeNode that satisfies the condition.
func (t *Tree) TraversePreorder(condition func(n *TreeNode) bool) (nodes []*TreeNode) {
	t.root.traversePreorder(condition, &nodes)
	return
}

func (node *TreeNode) traversePreorder(condition func(n *TreeNode) bool, nodes *[]*TreeNode) {
	if node != nil {
		// vist to the root
		if condition(node) {
			*nodes = append(*nodes, node)
		}
		child := node.firstChild
		for child != nil {
			child.traversePreorder(condition, nodes)
			child = child.nextBrother
		}
	}
}

// TraversePostorder perform a deep-first research and returns a slice of *TreeNode that satisfies the condition.
func (t *Tree) TraversePostorder(condition func(n *TreeNode) bool) (nodes []*TreeNode) {
	t.root.traversePostorder(condition, &nodes)
	return
}

func (node *TreeNode) traversePostorder(condition func(n *TreeNode) bool, nodes *[]*TreeNode) {
	if node != nil {
		child := node.firstChild
		for child != nil {
			child.traversePostorder(condition, nodes)
			child = child.nextBrother
		}
		// vist to the root
		if condition(node) {
			*nodes = append(*nodes, node)
		}
	}
}

// BreadthFirstSearch performs a breadth-first search on the tree and returns a list of the nodes that satysfies the condiiton.
func (t *Tree) BreadthFirstSearch(condition func(n *TreeNode) bool) (nodes []*TreeNode) {
	var q Queue = new(DoubleLinkedListQueue)
	q.Enqueue(t.root)
	for !q.IsEmpty() {
		n := q.Dequeue().(*TreeNode)
		if n != nil {
			if condition(n) {
				nodes = append(nodes, n)
			}

			child := n.firstChild
			for child != nil {
				q.Enqueue(child)
				child = child.nextBrother
			}
		}
	}
	return
}

// ParseTreeFromFile returns a tree from a file in the correct format.
// The value are parsed with function 'parseFromStr' passed as input.
func ParseTreeFromFile[T any](filename string, parseFromStr func(s string) (T, error)) (t *Tree) {
	// key is height
	// value is last-visited node of corresponding height
	nodes := make(map[int]*TreeNode)
	lines := readLines(filename)

	var node *TreeNode
	for _, line := range lines {
		height := strings.Count(line, " ")
		line = strings.TrimSpace(line)

		val, err := parseFromStr(line)
		if err != nil {
			panic(err)
		}
		node = NewNode(val)
		nodes[height] = node
		if height > 0 {
			nodes[height-4].AddChild(node)
		}
	}

	t = &Tree{nodes[0]} // the root have heght equal zero.
	return
}

func readLines(filename string) (res []string) {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return
}

// String returns a string representation of the tree.
func (t *Tree) String() string {
	var s string
	t.root.auxString(0, &s)
	return s
}

// Auxiliary function for String method.
func (node *TreeNode) auxString(spaces int, s *string) {
	if node != nil {
		*s += strings.Repeat("  ", spaces) + fmt.Sprintf("- %v\n", node.val)
		child := node.firstChild
		for child != nil {
			child.auxString(spaces+1, s)
			child = child.nextBrother
		}
	}
}
