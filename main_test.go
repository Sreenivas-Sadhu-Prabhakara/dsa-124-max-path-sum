package main

import "testing"

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		{tn(1, tn(2, nil, nil), tn(3, nil, nil)), 6},
		{tn(-10, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil))), 42}, // 15+20+7
		{tn(-3, nil, nil), -3}, // all negative: best single node
		{tn(-2, tn(-1, nil, nil), nil), -1},
		{tn(2, tn(-1, nil, nil), nil), 2},
		{nil, 0},
	}
	for _, c := range cases {
		if got := MaxPathSum(c.root); got != c.want {
			t.Errorf("MaxPathSum = %d, want %d", got, c.want)
		}
	}
}
