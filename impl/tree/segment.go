package tree

import (
	"log"
)

type Comp func(interface{}, interface{}) interface{}
type Mutate func(interface{}) interface{}

type SegmentTree struct {
	tree []interface{}
	comp Comp
	len  int
}

func NewSegmentTree(arr []interface{}, comp Comp) *SegmentTree {
	backing := 1
	for backing < len(arr) {
		backing <<= 1
	}

	tree := &SegmentTree{
		make([]interface{}, 2*backing-1),
		comp,
		len(arr),
	}

	log.Printf("%#v", arr)
	tree.build(arr, 0, 0, len(arr)-1)
	log.Printf("%#v", tree.tree)

	return tree
}

func (seg *SegmentTree) build(arr []interface{}, node, start, end int) {
	log.Printf("%d %d %d \n", node, start, end)
	if start == end {
		seg.tree[node] = arr[start]
	} else {
		mid := (start + end) / 2
		seg.build(arr, 2*node+1, start, mid)
		seg.build(arr, 2*node+2, mid+1, end)
		seg.tree[node] = seg.comp(seg.tree[2*node+1], seg.tree[2*node+2])
	}

	log.Printf("Set %d to %#v\n", node, seg.tree[node])
}

func (seg *SegmentTree) Query(start, end int) interface{} {
	return seg.query(0, 0, seg.len, start, end)
}

func (seg *SegmentTree) query(node, left, right, start, end int) interface{} {
	log.Printf(
		"Node: %d, Val: %#v, Left: %d, Right: %d, Start: %d, End: %d",
		node, seg.tree[node], left, right, start, end,
	)

	if start <= left && end >= right {
		return seg.tree[node]
	} else if start > right || end < left {
		return nil
	}

	mid := (left + right) / 2
	l := seg.query(2*node+1, left, mid, start, end)
	r := seg.query(2*node+2, mid+1, right, start, end)

	if r == nil {
		return l
	} else if l == nil {
		return r
	}

	return seg.comp(l, r)
}
