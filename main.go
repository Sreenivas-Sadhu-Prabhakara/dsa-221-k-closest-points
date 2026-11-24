package main

import (
	"fmt"
	"sort"
)

// KClosest returns the k points nearest the origin by Euclidean distance.
// Squared distances avoid the sqrt entirely — ordering is identical.
// Ties and output order are made deterministic by sorting (dist, x, y).
func KClosest(points [][2]int, k int) [][2]int {
	if k > len(points) {
		k = len(points)
	}
	if k <= 0 {
		return [][2]int{}
	}
	sorted := append([][2]int{}, points...)
	sort.Slice(sorted, func(i, j int) bool {
		di := sorted[i][0]*sorted[i][0] + sorted[i][1]*sorted[i][1]
		dj := sorted[j][0]*sorted[j][0] + sorted[j][1]*sorted[j][1]
		if di != dj {
			return di < dj
		}
		if sorted[i][0] != sorted[j][0] {
			return sorted[i][0] < sorted[j][0]
		}
		return sorted[i][1] < sorted[j][1]
	})
	return sorted[:k]
}

func main() {
	fmt.Println(KClosest([][2]int{{1, 3}, {-2, 2}}, 1))
	fmt.Println(KClosest([][2]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
