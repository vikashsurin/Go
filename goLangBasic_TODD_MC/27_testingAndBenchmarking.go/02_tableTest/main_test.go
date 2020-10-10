package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{1, 2, 3}, 6},
		test{[]int{21, 1}, 22},
		test{[]int{21, 21, 21}, 63},
		test{[]int{21, 100}, 121},
	}

	for _, v := range tests {
		xx := sum(v.data...)
		if xx != v.answer {
			t.Error("Expected", v.answer, "Got", xx)
		}
	}

}
