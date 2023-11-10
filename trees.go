package datastructs

import (
	"bufio"
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
