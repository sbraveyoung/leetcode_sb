package findMinArrowShots

import (
	"sort"
)

type interval [][]int

func (i interval) Len() int {
	return len(i)
}
func (i interval) Less(x, y int) bool {
	if i[x][0] < i[y][0] {
		return true
	} else if i[x][0] > i[y][0] {
		return false
	} else {
		return i[x][1] < i[y][1]
	}
}
func (i interval) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func findMinArrowShots(points interval) int {
	if len(points) <= 1 {
		return len(points)
	}
	sort.Sort(points)
	count := 1
	_, right := points[0][0], points[0][1]
	for index := 1; index < len(points); index++ {
		if points[index][0] <= right {
			if points[index][1] <= right {
				_, right = points[index][0], points[index][1]
			} else {
				//	_, right = points[index][0], right
			}
		} else {
			count++
			_, right = points[index][0], points[index][1]
		}
	}
	return count
}
