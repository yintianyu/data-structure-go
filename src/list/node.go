package list

// node为链表中的节点
type Node struct {
	data Datatype
	pred *Node
	next *Node
}

// createNode：创建节点，data为e，pred和next均为nil
func createNode(e Datatype) *Node {
	newNode := Node{
		data: e,
		pred: nil,
		next: nil,
	}
	return &newNode
}

// Data：返回该节点的数据data
func (n *Node) Data() Datatype {
	return n.data
}

// Pred：返回该节点前驱节点的指针
func (n *Node) Pred() *Node {
	return n.pred
}

// Next：返回该节点后继节点的指针
func (n *Node) Next() *Node {
	return n.next
}

// InsertAsPred：将值e作为该节点的前驱节点插入
func (n *Node) InsertAsPred(e Datatype) {
	newNode := createNode(e)
	predNode := n.Pred()
	predNode.next = newNode
	newNode.pred = predNode
	n.pred = newNode
	newNode.next = n
}

// InsertAsNext：将值e作为该节点的后继节点插入
func (n *Node) InsertAsNext(e Datatype) {
	newNode := createNode(e)
	nextNode := n.Next()
	nextNode.pred = newNode
	newNode.next = nextNode
	n.next = newNode
	newNode.pred = n
}
