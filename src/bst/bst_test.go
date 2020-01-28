package bst

import (
	"fmt"
	"runtime"
	"testing"
)

func fileLine() string {
	_, fileName, fileLine, ok := runtime.Caller(1)
	var s string
	if ok {
		s = fmt.Sprintf("%s:%d", fileName, fileLine)
	} else {
		s = ""
	}
	return s
}

func TestBST(*testing.T) {
	tree := CreateBST()
	tree.Insert(36)
	tree.Insert(27)
	tree.Insert(6)
	tree.Insert(58)
	tree.Insert(53)
	tree.Insert(69)
	tree.Insert(40)
	tree.Insert(46)
	tree.Insert(54)
	pst := tree.Search(5)
	if pst != nil {
		panic(fmt.Errorf("error %s\n", fileLine()))
	}
	pst = tree.Search(69)
	if pst == nil || pst.data != 69 {
		panic(fmt.Errorf("error %s\n", fileLine()))
	}
	tree.Remove(40)
	pst = tree.Search(40)
	if pst != nil {
		panic(fmt.Errorf("error %s\n", fileLine()))
	}
	pst = tree.Search(54)
	if pst == nil || pst.data != 54 {
		panic(fmt.Errorf("error %s\n", fileLine()))
	}
	fmt.Printf("BSTTesting Pass!\n")
}
