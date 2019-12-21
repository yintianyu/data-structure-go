package vector

import (
	_ "errors"
	"fmt"
)

type Datatype int

type Vector struct {
	data []Datatype
}

// Size, return the size of vector
func (v *Vector) Size() int {
	return len(v.data)
}

// Get, return the rth element of vector
func (v *Vector) Get(r int) Datatype {
	return v.data[r]
}

// Put, replace rth element with e
func (v *Vector) Put(r int, e Datatype) error {
	if v.Size() <= r {
		return fmt.Errorf("no %dth element", r)
	}
	v.data[r] = e
	return nil
}

// Insert, insert e at r, and push back all the elements after r
func (v *Vector) Insert(r int, e Datatype) error {
	if v.Size() < r {
		return fmt.Errorf("no %dth element", r)
	}
	if v.Size() == r {
		v.data = append(v.data, e)
	} else {
		v.data = append(v.data[:r+1], v.data[r:]...)
		v.data[r] = e
	}
	return nil
}

// Remove: 删除下标为r的元素
func (v *Vector) Remove(r int) error {
	if v.Size() <= r {
		return fmt.Errorf("no %dth element", r)
	}
	v.data = append(v.data[:r], v.data[r+1:]...)
	return nil
}

// Disordered：判断所有元素是否已按非降序排列
func (v *Vector) Disordered() bool {
	length := v.Size()
	if length == 0 || length == 1 {
		return false
	}
	last := v.data[0]
	for i := 1; i < length; i += 1 {
		if v.data[i] >= last {
			return true
		}
		last = v.data[i]
	}
	return false
}

// Sort: 使用快速排序对向量进行排序，使其按降序排列
func (v *Vector) Sort() {
	length := v.Size()
	if length == 0 || length == 1 {
		return
	}
	v.quickSort(0, length-1)
}

// quickSort：从start到end之间，选取start作为pivot进行快速排序（降序）
func (v *Vector) quickSort(start int, end int) {
	if start >= end {
		return
	}
	pivot := v.data[start]
	pivotIdx := start
	for i := start + 1; i <= end; i++ {
		if pivot < v.data[i] {
			v.data[pivotIdx] = v.data[i]
			v.data[i] = pivot
			pivotIdx = i
		}
	}
	v.quickSort(start, pivotIdx-1)
	v.quickSort(pivotIdx+1, end)
}

// Find: 查找等于e且秩最大的元素，返回其下标，未找到返回-1
func (v *Vector) Find(e Datatype) int {
	length := v.Size()
	var i int
	for i = length - 1; i >= 0; i-- {
		if v.data[i] == e {
			break
		}
	}
	return i
}

// Search: 查找目标元素e，返回不大于e且秩最大的元素，如果没有不大于e的元素，则返回error
func (v *Vector) Search(e Datatype) (Datatype, error) {
	length := v.Size()
	var i int
	for i = length - 1; i >= 0; i-- {
		if v.data[i] <= e {
			return v.data[i], nil
		}
	}
	return 0, fmt.Errorf("not found")
}

// Deduplicate: 剔除向量中的重复元素（不要求是有序向量）
func (v *Vector) Deduplicate() {
	set := make(map[Datatype]bool)
	length := v.Size()
	for i := 0; i < length; i++ {
		_, exist := set[v.data[i]]
		if exist {
			err := v.Remove(i)
			i -= 1
			length -= 1
			if err != nil {
				panic(fmt.Errorf("duplicate unexpected error"))
			}
		} else {
			set[v.data[i]] = true
		}
	}
}

// Uniquify： 剔除向量中的重复元素（要求是有序向量）
func (v *Vector) Uniquify() {
	length := v.Size()
	last := v.data[0]
	for i := 1; i < length; i++ {
		if v.data[i] == last {
			err := v.Remove(i)
			if err != nil {
				panic(fmt.Errorf("uniquify unexpected error"))
			}
			i -= 1
			length -= 1
		} else {
			last = v.data[i]
		}
	}
}

type T func(i *Datatype)

// Traverse：遍历每个元素，执行某个函数
func (v *Vector) Traverse(f T) {
	length := v.Size()
	for i := 0; i < length; i++ {
		f(&v.data[i])
	}
}

func (v *Vector) Print() {
	length := v.Size()
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", v.data[i])
	}
	fmt.Printf("\n")
}
