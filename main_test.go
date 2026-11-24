package main

import (
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	cases := []struct {
		points [][2]int
		k      int
		want   [][2]int
	}{
		{[][2]int{{1, 3}, {-2, 2}}, 1, [][2]int{{-2, 2}}},
		{[][2]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][2]int{{3, 3}, {-2, 4}}},
		{[][2]int{{0, 0}}, 1, [][2]int{{0, 0}}},
		{[][2]int{{1, 0}, {0, 1}}, 2, [][2]int{{0, 1}, {1, 0}}}, // tie broken by x
		{[][2]int{{2, 2}}, 5, [][2]int{{2, 2}}},                 // k beyond size: everything
		{nil, 3, [][2]int{}},
	}
	for _, c := range cases {
		if got := KClosest(c.points, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("KClosest(%v,%d) = %v, want %v", c.points, c.k, got, c.want)
		}
	}
}
