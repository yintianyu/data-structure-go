package list

import (
	"fmt"
	"testing"
)

func TestList(*testing.T) {
	list := CreateList()
	size := list.Size()
	if size != 0 {
		panic(fmt.Errorf("size error"))
	}
	list.InsertAsFirst(0)
	list.InsertAsFirst(1)
	list.InsertAsFirst(2)
	list.InsertAsLast(3)
	list.InsertAsLast(4)
	list.InsertAsLast(5)
	size = list.Size()
	if size != 6 {
		panic(fmt.Errorf("insert error"))
	}
	list.Print()
	first, err := list.First()
	if first != 2 || err != nil {
		panic(fmt.Errorf("first error"))
	}
	last, err := list.Last()
	if last != 5 || err != nil {
		panic(fmt.Errorf("last error"))
	}
	pstNode2 := list.header.next.next.next
	err = InsertBefore(pstNode2, 6)
	if err != nil || list.header.next.next.next.Data() != 6 {
		panic(fmt.Errorf("insertBefore error"))
	}
	pstNode2 = list.header.next.next.next
	err = InsertAfter(pstNode2, 7)
	if err != nil || list.header.next.next.next.next.Data() != 7 {
		panic(fmt.Errorf("insertAfter error"))
	}

	fmt.Printf("TestList Pass!")

}
