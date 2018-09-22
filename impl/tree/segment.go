package tree

import (
	"log"
)

type Comp func(interface{}, interface{}) interface{}
type Mutate func(interface{}) interface{}

type SegmentTree struct {
	tree []interface{}
	comp Comp
}

func NewSegmentTree(arr []interface{}, comp Comp) *SegmentTree {
	backing := 1
	for backing < len(arr) {
		backing <<= 1
	}

	tree := &SegmentTree{
		make([]interface{}, 2*backing-1),
		comp,
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
