package arrays

import (
	"testing"
)

func Test_EvenOdd(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{6, 2, 4, 5, 3, 1},
		},
	}

	compare := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}

		for i, num := range a {
			if b[i] != num {
				return false
			}
		}

		return true
	}

	for _, test := range tests {
		got := EvenOdd(test.arr)
		if !compare(test.want, got) {
			t.Errorf("Want: %#v Got: %#v", test.want, got)
		}
	}
}
