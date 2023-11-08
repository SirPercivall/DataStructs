package datastructs

// TreeNode represent a node of a Tree.
type BinaryTreeNode struct {
	val    any
	father *BinaryTreeNode
	sx     *BinaryTreeNode
	dx     *BinaryTreeNode
}

// NewBinaryTreeNode returns a pointer to a BinaryTreeNode initialized with value passed as input.
func NewBinaryTreeNode(val any) *BinaryTreeNode {
	if val == nil {
		return nil
	} else {
		return &BinaryTreeNode{val, nil, nil, nil}
	}
}

// GetVal returns the value of the node.
func (node *BinaryTreeNode) GetVal() any {
	return node.val
}

// Left add left child initialized as value passed as input.
func (node *BinaryTreeNode) AddLeft(val any) {
	newNode := NewBinaryTreeNode(val)
	newNode.father = node
	node.sx = newNode
}

// AddRight add right child initialized as value passed as input.
func (node *BinaryTreeNode) AddRight(val any) {
	newNode := NewBinaryTreeNode(val)
	newNode.father = node
	node.dx = newNode
}

// BinaryTree represent a tree data type recursively defined as a root and two secondary-trees.
type BinaryTree struct {
	root *BinaryTreeNode
}

// NewBinaryTree returns a pointer to a BinaryTree
func NewBinaryTree(n *BinaryTreeNode) *BinaryTree {
	return &BinaryTree{n}
}

// BreadFirstSearch perform a breadth-first search on the tree and returns the nodes that satisfy the passed function.
func (t *BinaryTree) BreadthFirstSearch(condition func(n *BinaryTreeNode) bool) (nodes []*BinaryTreeNode) {
	var q Queue = new(DoubleLinkedListQueue)
	q.Enqueue(t.root)
	for !q.IsEmpty() {
		n := q.Dequeue().(*BinaryTreeNode)
		if n != nil {
			if condition(n) {
				nodes = append(nodes, n)
			}
			q.Enqueue(n.sx)
			q.Enqueue(n.dx)
		}
	}

	return
}

// TraversePreorder perform a deep-first search and returns a slice of nodes that satisfies the function.
func (t *BinaryTree) TraversePreorder(condition func(n *BinaryTreeNode) bool) (nodes []*BinaryTreeNode) {
	traversePreorder(t.root, &nodes, condition)
	return
}

func traversePreorder(node *BinaryTreeNode, nodes *[]*BinaryTreeNode, condition func(n *BinaryTreeNode) bool) {
	if node != nil {
		if condition(node) {
			(*nodes) = append((*nodes), node)
		}
		traversePreorder(node.sx, nodes, condition)
		traversePreorder(node.dx, nodes, condition)
	}
}

// TraverseInorder perform a deep-first search and returns a slice of nodes that satisfies the function.
func (t *BinaryTree) TraverseInorder(condition func(n *BinaryTreeNode) bool) (nodes []*BinaryTreeNode) {
	traverseInorder(t.root, &nodes, condition)
	return
}

func traverseInorder(node *BinaryTreeNode, nodes *[]*BinaryTreeNode, condition func(n *BinaryTreeNode) bool) {
	if node != nil {
		traverseInorder(node.sx, nodes, condition)
		if condition(node) {
			(*nodes) = append((*nodes), node)
		}
		traverseInorder(node.dx, nodes, condition)
	}
}

// TraverseInorder perform a deep-first search and returns a slice of nodes that satisfies the function.
func (t *BinaryTree) TraversePostorder(condition func(n *BinaryTreeNode) bool) (nodes []*BinaryTreeNode) {
	traversePostorder(t.root, &nodes, condition)
	return
}

func traversePostorder(node *BinaryTreeNode, nodes *[]*BinaryTreeNode, condition func(n *BinaryTreeNode) bool) {
	if node != nil {
		traversePostorder(node.sx, nodes, condition)
		traversePostorder(node.dx, nodes, condition)
		if condition(node) {
			(*nodes) = append((*nodes), node)
		}
	}
}

// BuildToNodeSlice returns a slice of pointers from a slice of any values.
func BuildNodeSlice(values []any) (nodes []*BinaryTreeNode) {
	for _, v := range values {
		if v == nil {
			nodes = append(nodes, nil)
		} else {
			newNode := NewBinaryTreeNode(v)
			nodes = append(nodes, newNode)
		}
	}
	return
}

// BuildTreeFromSlice returns a pointer to an instance of BinaryTree, initialized with the slice of nodes.
// The slice of nodes must be a result of a breadth-first search.
func BuildTreeFromSlice(sl []any) *BinaryTree {
	var nodes []*BinaryTreeNode
	for _, v := range sl {
		n := NewBinaryTreeNode(v)
		nodes = append(nodes, n)
	}

	for i, v := range nodes {
		// checking if node is a leaf
		if 2*i+1 < len(nodes) {
			if v != nil {
				var sx, dx *BinaryTreeNode
				sx = nodes[2*i+1]

				if 2*i+2 < len(nodes) {
					dx = nodes[2*i+2]
				}
				v.sx = sx
				v.dx = dx

				if sx != nil {
					sx.father = v
				}

				if dx != nil {
					dx.father = v
				}

			}
		}
	}

	return NewBinaryTree(nodes[0])
}
