package list

import "fmt"

type Datatype int

type List struct {
	header *Node
	tail   *Node
	//size int
}

// CreateList：创建一个空链表
func CreateList() *List {
	newList := List{
		header: createNode(0),
		tail:   createNode(0),
		//size	:	0,
	}
	newList.header.next = newList.tail
	newList.tail.pred = newList.header
	return &newList
}

// Size：返回该链表的大小
func (l *List) Size() int {
	cnt := 0
	pst := l.header
	for pst.next != l.tail {
		cnt += 1
		pst = pst.next
	}
	return cnt
}

// Empty：判断该链表是否为空，为空返回true，否则返回false
func (l *List) Empty() bool {
	return l.header.next == l.tail
}

// Print：遍历一次，打印所有的值
func (l *List) Print() {
	pst := l.header.next
	for pst != l.tail {
		fmt.Printf("%d ", pst.Data())
		pst = pst.next
	}
	fmt.Printf("\n")
}

// First：返回首节点的值，当链表为空时，返回error
func (l *List) First() (Datatype, error) {
	if l.Empty() {
		return 0, fmt.Errorf("empty list")
	}
	return l.header.Next().Data(), nil
}

// Last：返回尾节点的值，当链表为空时，返回error
func (l *List) Last() (Datatype, error) {
	if l.Empty() {
		return 0, fmt.Errorf("empty list")
	}
	return l.tail.Pred().Data(), nil
}

// InsertAsFirst：将值e作为首节点插入
func (l *List) InsertAsFirst(e Datatype) {
	l.header.InsertAsNext(e)
}

// InsertAsLast：将值e作为尾节点插入
func (l *List) InsertAsLast(e Datatype) {
	l.tail.InsertAsPred(e)
}

// InsertBefore：将值e作为节点p的前驱插入
func InsertBefore(p *Node, e Datatype) error {
	pred := p.Pred()
	if pred == nil { // p的前驱不存在，可能是header虚节点
		return fmt.Errorf("invalid node p")
	}
	p.InsertAsPred(e)
	return nil
}

// InsertAfter：将值e作为节点p的后继插入
func InsertAfter(p *Node, e Datatype) error {
	next := p.Next()
	if next == nil { // p的后继不存在，可能是tailer虚节点
		return fmt.Errorf("invalid node p")
	}
	p.InsertAsNext(e)
	return nil
}

// Remove：删除位置p处的节点，返回其数值
func Remove(p *Node) (Datatype, error) {
	if p == nil {
		return 0, fmt.Errorf("remove error, p is nil")
	}
	if p.Pred() == nil {
		return 0, fmt.Errorf("remove error, p is header node")
	}
	if p.Next() == nil {
		return 0, fmt.Errorf("remove error, p is tail node")
	}
	p.pred.next = p.next
	p.next.pred = p.pred
	return p.data, nil
}

// Disordered：判断所有节点是否已按非降序排列
func (l *List) Disordered() bool {
	if l.header.next == l.tail { // empty list
		return true
	}
	last := l.header.next.data
	for pst := l.header.next.next; pst != l.tail; pst = pst.next {
		if pst.data < last {
			return false
		}
	}
	return true
}

// Sort：调整各节点的位置，使之按非降序排列。目前使用归并排序
func (l *List) Sort() {

}

// mergeSort：使用归并排序对链表进行排序
func (l *List) mergeSort() {
	// 如果链表为空或者长度为1，则直接退出
	if l.header.next == l.tail || l.header.next.next == l.tail {
		return
	}
	// 先用相向而行找到中间的那个
	pst1 := l.header.next
	pst2 := l.tail.pred
	for pst1 != pst2 && pst1.next == pst2 {
		pst1 = pst1.next
		pst2 = pst2.pred
	}
	// 将链表从pst1后面开始切开，切成两个链表
	pst2 = pst1.next
	pst1.next = l.tail
	oldTail := l.tail.pred
	l.tail.pred = pst1.next
	newList := CreateList()
	newList.header.next = pst2
	pst2.pred = newList.header
	newList.tail.pred = oldTail
	oldTail.next = newList.tail

	// 对这两个新链表分别进行归并排序
	l.mergeSort()
	newList.mergeSort()
	// 对这两个链表进行合并

}

// mergeList：合并l和l2，要求l和l2均为非降序排列，最终保证l是完整链表
func (l *List) mergeList(l2 *List) {
	if l2.header.next == l2.tail {
		return
	}
	if l.header.next == l.tail {
		l.header.next = l2.header.next
		l2.header.next.pred = l.header
		l.tail.pred = l2.tail.pred
		l2.tail.pred.next = l.tail
		return
	}
	pst1 := l.header.next
	pst2 := l2.header.next
	for pst1 != l.tail && pst2 != l2.tail {
		if pst1.data <= pst2.data {
			pst1 = pst1.next
		} else { // 把pst2插在pst1的前面
			tmp := pst2.next
			pst1.pred.next = pst2
			pst2.pred = pst1.pred
			pst2.next = pst1
			pst1.pred = pst2
			pst2 = tmp
		}
	}
	if pst2 != l2.tail { // l2还没完，将其插入到后面
		pst1.next = pst2
		pst2.pred = pst1
		l.tail.pred = l2.tail.pred
		l2.tail.pred.next = l.tail
	}
	return
}

// TODO: 进行List的合法性检查，即正向遍历反向遍历结果相同
