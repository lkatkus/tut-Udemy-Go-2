package main

import (
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestMySum(t *testing.T) {
	tests := []test{
		test{
			data:   []int{1, 2, 3},
			answer: 6,
		},
		test{
			data:   []int{4, 4},
			answer: 8,
		},
		test{
			data:   []int{5, 95},
			answer: 100,
		},
		test{
			data:   []int{1, 1000},
			answer: 1,
		},
	}

	for _, test := range tests {
		x := mySum(test.data...)
		if x != test.answer {
			t.Error("Expected 6, but got", x)
		}
	}
}
