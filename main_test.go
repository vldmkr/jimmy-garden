package main

import "testing"

type testCases struct {
	arg  []int
	want int
}

func TestSolution(t *testing.T) {
	cases := []testCases{
		{
			arg:  []int{3, 4, 5, 3, 7},
			want: 3,
		},
		{
			arg:  []int{1, 2, 3, 4},
			want: -1,
		},
		{
			arg:  []int{1, 3, 1, 2},
			want: 0,
		},
		{
			arg:  []int{6, 6, 9, 4},
			want: 2,
		},
		{
			arg:  []int{3, 11, 4, 7, 12},
			want: 3,
		},
		{
			arg:  []int{6, 4, 5, 6, 5},
			want: 2,
		},
		{
			arg:  []int{8, 6, 9, 9, 9},
			want: -1,
		},
		{
			arg:  []int{10, 23, 11, 15, 5, 6, 2, 12},
			want: 0,
		},
	}

	for _, test := range cases {
		if got := Solution(test.arg); got != test.want {
			t.Errorf("%v solution is %d but should be %d", test.arg, got, test.want)
		}
	}
}
