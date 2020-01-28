package bst

// Define Binary Search Tree struct
type BST struct {
	root *TreeNode
	hot  *TreeNode // The last node visited
}

func CreateBST() *BST {
	newTree := new(BST)
	return newTree
}

// Function: Search
// Description: search specified value in the BST
// Input: value, to be searched
// Output: The pointer of the node found or nil when 404
func (tree *BST) Search(value DataType) *TreeNode {
	pst := tree.root
	tree.hot = nil
	for {
		if pst == nil || value == pst.data {
			return pst
		}
		tree.hot = pst
		if value < pst.data {
			pst = pst.leftChild
		} else {
			pst = pst.rightChild
		}
	}
}

// Function: Insert
// Description: Insert specified value in BST
// Input: value, to be inserted
// Output: pointer of the newly inserted node(or existed node of the same value)
func (tree *BST) Insert(value DataType) *TreeNode {
	pst := tree.Search(value)
	if pst != nil {
		return pst
	}
	newTreeNode := CreateSonTreeNode(value, tree.hot)
	if tree.hot == nil {
		tree.root = newTreeNode
	}
	return newTreeNode
}

// Function: Remove
// Description: Remove specified value in BST
// Input: value, to be removed
// Output: (None)
func (tree *BST) Remove(value DataType) {
	pst := tree.Search(value)
	if pst == nil {
		return
	}
	if tree.hot == nil { // Delete root
		if pst.leftChild == nil {
			tree.root = pst.rightChild
		} else if pst.rightChild == nil {
			tree.root = pst.leftChild
		} else {
			succ := Successor(pst)
			succVal := succ.data
			originalHot := tree.hot
			tree.Remove(succVal)
			tree.hot = originalHot
			pst.data = succVal
		}
	}
	var left bool
	if tree.hot.leftChild != nil && tree.hot.leftChild.data == pst.data {
		left = true
	} else {
		left = false
	}
	if pst.leftChild == nil {
		if left {
			tree.hot.leftChild = pst.rightChild
		} else {
			tree.hot.rightChild = pst.rightChild
		}
	} else if pst.rightChild == nil {
		if left {
			tree.hot.leftChild = pst.leftChild
		} else {
			tree.hot.rightChild = pst.leftChild
		}
	} else {
		succ := Successor(pst)
		succVal := succ.data
		originalHot := tree.hot
		tree.Remove(succVal)
		tree.hot = originalHot
		pst.data = succVal
	}
}
