package tree

import (
	"log"
)

type Comp func(interface{}, interface{}) interface{}
type Mutate func(interface{}) interface{}

type SegmentTree struct {
	tree []interface{}
	lazy [][]Mutate
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
		make([][]Mutate, 2*backing-1),
		comp,
		len(arr),
	}
	tree.build(arr, 0, 0, tree.len-1)

	return tree
}

func (seg *SegmentTree) build(arr []interface{}, node, start, end int) {
	if start == end {
		seg.tree[node] = arr[start]
		return
	}

	mid := (start + end) / 2
	seg.build(arr, 2*node+1, start, mid)
	seg.build(arr, 2*node+2, mid+1, end)
	seg.tree[node] = seg.comp(seg.tree[2*node+1], seg.tree[2*node+2])
}

func (seg *SegmentTree) Print() {
	seg.print(0, 0, seg.len-1)
}

func (seg *SegmentTree) print(node, start, end int) {
	log.Printf(
		"Node: %d Value: %d Range: %d-%d Lazy: %d\n",
		node, seg.tree[node], start, end, len(seg.lazy[node]),
	)

	if start == end {
		return
	}

	mid := (start + end) / 2
	seg.print(2*node+1, start, mid)
	seg.print(2*node+2, mid+1, end)
}

func (seg *SegmentTree) propagate(node int, mut Mutate, leaf bool) {
	seg.tree[node] = mut(seg.tree[node])

	if !leaf {
		seg.lazy[2*node+1] = append(seg.lazy[2*node+1], mut)
		seg.lazy[2*node+2] = append(seg.lazy[2*node+2], mut)
	}
}

func (seg *SegmentTree) Update(start, end int, mut Mutate) {
	seg.update(0, 0, seg.len-1, start, end, mut)
}

func (seg *SegmentTree) update(node, left, right, start, end int, mut Mutate) {
	for _, lazy := range seg.lazy[node] {
		seg.propagate(node, lazy, left == right)
	}
	seg.lazy[node] = nil

	if start > right || end < left {
		return
	}

	if left >= start && right <= end {
		seg.propagate(node, mut, left == right)
		return
	}

	mid := (left + right) / 2
	seg.update(2*node+1, left, mid, start, end, mut)
	seg.update(2*node+2, mid+1, right, start, end, mut)
	seg.tree[node] = seg.comp(seg.tree[2*node+1], seg.tree[2*node+2])
}

func (seg *SegmentTree) Query(start, end int) interface{} {
	return seg.query(0, 0, seg.len-1, start, end)
}

func (seg *SegmentTree) query(node, left, right, start, end int) interface{} {
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
