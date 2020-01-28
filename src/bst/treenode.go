package bst

type DataType int

type TreeNode struct { // 二叉树节点
	data       DataType
	leftChild  *TreeNode
	rightChild *TreeNode
}

// Function: CreateTreeNode
// Description: Create a tree node, return its pointer
// Input: e, the data value of the newly created node
// Output: The pointer of the newly created node
func CreateTreeNode(e DataType) *TreeNode {
	newTreeNode := TreeNode{
		data:       e,
		leftChild:  nil,
		rightChild: nil,
	}
	return &newTreeNode
}

// Function: CreateSonTreeNode
// Description: Create a tree node, and making it the son of specified node, return its pointer
// Input: e, the data value of the newly created node; n, the father of the new node
// Output: The pointer of the newly created node
func CreateSonTreeNode(e DataType, n *TreeNode) *TreeNode {
	newTreeNode := CreateTreeNode(e)
	if n == nil {
		return newTreeNode
	}
	if e < n.data {
		n.leftChild = newTreeNode
	} else {
		n.rightChild = newTreeNode
	}
	return newTreeNode
}

// Function: Successor
// Description: Get the direct successor of the specified node
// Input: n, pointer of the specified node
// Output: n's direct successor or nil if it doesn't exist
func Successor(n *TreeNode) *TreeNode {
	if n.rightChild == nil {
		return nil
	}
	pst := n.rightChild
	for pst.leftChild != nil {
		pst = pst.leftChild
	}
	return pst
}
