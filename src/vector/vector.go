package vector

import (
	_ "errors"
	"fmt"
)

type Vector struct {
	data []interface{}
}

// Size, return the size of vector
func (v *Vector) Size() int {
	return len(v.data)
}

// Get, return the rth element of vector
func (v *Vector) Get(r int) interface{} {
	return v.data[r]
}

// Put, replace rth element with e
func (v *Vector) Put(r int, e interface{}) error {
	if v.Size() <= r {
		return fmt.Errorf("no %dth element", r)
	}
	v.data[r] = e
	return nil
}

// Insert, insert e at r, and push back all the elements after r
func (v *Vector) Insert(r int, e interface{}) error {
	if v.Size() < r {
		return fmt.Errorf("no %dth element", r)
	} else if v.Size() == r {
		v.data = append(v.data[:r], e)
	} else {
		v.data = append(v.data[:r], e, v.data[r:])
	}
	return nil
}
