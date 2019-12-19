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

	fmt.Printf("TestVector Pass!")

}
