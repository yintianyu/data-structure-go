package vector

import (
	"fmt"
	"testing"
)

func TestVector(*testing.T) {
	var v Vector

	err := v.Insert(0, 3)

	if err != nil {
		panic(fmt.Errorf("insert error"))
	}

	if v.Size() != 1 {
		panic(fmt.Errorf("size error"))
	}

	err = v.Insert(0, 4)
	if err != nil {
		panic(fmt.Errorf("insert error"))
	}
	err = v.Insert(0, 5)
	if err != nil {
		panic(fmt.Errorf("insert error"))
	}
	err = v.Insert(0, 6)
	if err != nil {
		panic(fmt.Errorf("insert error"))
	}
	err = v.Insert(0, 7)
	if err != nil {
		panic(fmt.Errorf("insert error"))
	}
	err = v.Insert(0, 8)
	if err != nil {
		panic(fmt.Errorf("insert error"))
	}

	err = v.Remove(4)
	if err != nil || v.data[4] != 3 || v.Size() != 5 {
		panic(fmt.Errorf("remove error"))
	}

	v.Print()

	disordered := v.Disordered()
	if disordered == true {
		panic(fmt.Errorf("disordered error"))
	}

	err = v.Insert(0, 1)
	err = v.Insert(0, -3)

	v.Print()

	v.Sort()

	v.Print()

	disordered = v.Disordered()
	if disordered == true {
		panic(fmt.Errorf("sort error"))
	}

	position := v.Find(5)
	if position != 3 {
		panic(fmt.Errorf("find error"))
	}

	result, err := v.Search(10)
	if err != nil || result != -3 {
		panic(fmt.Errorf("search error"))
	}

	result, err = v.Search(-5)
	if err == nil {
		panic(fmt.Errorf("find error 2"))
	}

	err = v.Insert(3, 1)
	err = v.Insert(3, 8)

	if v.Size() != 9 {
		panic(fmt.Errorf("insert error 2"))
	}
	v.Print()

	v.Deduplicate()
	v.Print()

	if v.Size() != 7 {
		panic(fmt.Errorf("deduplicate error"))
	}

	err = v.Insert(2, 6)
	err = v.Insert(2, 6)
	err = v.Insert(6, 5)

	v.Uniquify()

	if v.Size() != 7 {
		panic(fmt.Errorf("uniquify error"))
	}

	v.Traverse(add1)
	if v.data[3] != 2 {
		panic(fmt.Errorf("traverse error"))
	}

	fmt.Printf("TestVector Pass!")

}

func add1(i *Datatype) {
	*i += 1
}
