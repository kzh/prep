package tree

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	arr := []interface{}{1, 3, 5, 7, 9, 11}
	min := func(a interface{}, b interface{}) interface{} {
		ax, bx := a.(int), b.(int)
		return ax + bx
	}

	NewSegmentTree(arr, min)
}
