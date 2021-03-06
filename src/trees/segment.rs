macro_rules! left {
    ($node:expr) => {
        2 * $node + 1
    };
}

macro_rules! right {
    ($node:expr) => {
        2 * $node + 2
    };
}

#[derive(Debug)]
pub struct SegmentTree<T: Copy + Default> {
    tree: Vec<T>,
    comb: fn(T, T) -> T,
    len: usize,
}

impl<T: Copy + Default> SegmentTree<T> {
    pub fn new(arr: Vec<T>, comb: fn(T, T) -> T) -> SegmentTree<T> {
        let mut backing = 1;
        while backing < arr.len() {
            backing <<= 1;
        }
        backing = 2 * backing - 1;

        let mut s = SegmentTree {
            tree: vec![Default::default(); backing],
            comb,
            len: arr.len() as usize,
        };
        s.build(&arr, 0, 0, s.len - 1);
        s
    }

    fn build(&mut self, arr: &Vec<T>, node: usize, start: usize, end: usize) {
        if start == end {
            self.tree[node] = arr[start];
            return;
        }

        let mid = (start + end) / 2;
        self.build(arr, left!(node), start, mid);
        self.build(arr, right!(node), mid + 1, end);
        self.tree[node] = (self.comb)(self.tree[left!(node)], self.tree[right!(node)]);
    }

    fn range(&self, start: usize, end: usize) -> Option<T> {
        self.query_recurse(0, 0, self.len - 1, start, end)
    }

    fn query_recurse(
        &self,
        node: usize,
        left: usize,
        right: usize,
        start: usize,
        end: usize,
    ) -> Option<T> {
        if start <= left && end >= right {
            return Some(self.tree[node]);
        } else if start > right || end < left {
            return None;
        }

        let mid = (left + right) / 2;
        let l = self.query_recurse(left!(node), left, mid, start, end);
        let r = self.query_recurse(right!(node), mid + 1, right, start, end);

        match (l, r) {
            (_, None) => l,
            (None, _) => r,
            (Some(lv), Some(rv)) => Some((self.comb)(lv, rv)),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn segment_tree() {
        let arr = vec![1, 3, 5, 7, 9, 11];
        println!("Array: {:#?}", arr);

        let tree = SegmentTree::new(arr, |a, b| a + b);
        println!("Tree: {:#?}", tree);

        if let Some(sum) = tree.range(0, 5) {
            println!("{}", sum);
        }
    }
}
