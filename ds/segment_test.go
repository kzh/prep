package tree

import (
	"testing"
)

func sum(a interface{}, b interface{}) interface{} {
	ax, bx := a.(int), b.(int)
	return ax + bx
}

func min(a interface{}, b interface{}) interface{} {
	ax, bx := a.(int), b.(int)
	if ax > bx {
		return bx
	}
	return ax
}

func max(a interface{}, b interface{}) interface{} {
	ax, bx := a.(int), b.(int)
	if ax > bx {
		return ax
	}
	return bx
}

func modify(delta int) Mutate {
	return func(num interface{}) interface{} {
		x := num.(int)
		return x + delta
	}
}

func TestSegmentTreeQuery(t *testing.T) {
	type query struct {
		start, end int
		res        interface{}
	}

	tests := []struct {
		arr     []interface{}
		comp    Comp
		queries []query
	}{
		{
			[]interface{}{1, 3, 5, 7, 9, 11},
			sum,
			[]query{
				{0, 2, 9},
				{0, 5, 36},
				{2, 3, 12},
			},
		},
		{
			[]interface{}{5},
			sum,
			[]query{
				{0, 0, 5},
				{0, 2, 5},
			},
		},
		{
			[]interface{}{69, 89, 96, 38, 59, 42},
			min,
			[]query{
				{0, 5, 38},
				{4, 5, 42},
				{0, 2, 69},
			},
		},
		{
			[]interface{}{67, 50, 52, 26, 84, 11, 200},
			max,
			[]query{
				{0, 6, 200},
				{0, 5, 84},
				{1, 3, 52},
				{0, 2, 67},
			},
		},
	}

	for _, test := range tests {
		tree := NewSegmentTree(test.arr, test.comp)

		for _, q := range test.queries {
			got := tree.Query(q.start, q.end)
			if got != q.res {
				t.Errorf(
					"Failed Query:\nArr: %#v\nRange: %d - %d | Expected: %#v, Got: %#v",
					test.arr, q.start, q.end, q.res, got,
				)
			}
		}
	}
}

func TestSegmentTreeUpdate(t *testing.T) {
	tree := NewSegmentTree([]interface{}{1, 2, 3}, min)
	tree.Update(0, 1, func(a interface{}) interface{} {
		ax := a.(int)
		return ax + 5
	})
	tree.Print()
}