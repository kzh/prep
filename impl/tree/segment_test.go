package tree

import (
	"log"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	arr := []interface{}{1, 3, 5, 7, 9, 11}
	min := func(a interface{}, b interface{}) interface{} {
		ax, bx := a.(int), b.(int)
		return ax + bx
	}

	tree := NewSegmentTree(arr, min)
	tree.Print()

	res := tree.Query(0, 6)
	log.Println(res)
}
