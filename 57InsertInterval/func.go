package insertPackage

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var resInterval [][]int
	left, right := newInterval[0], newInterval[1]
	for index := 0; index < len(intervals); index++ {
		interval := intervals[index]
		if left < interval[0] {
			if right < interval[0] {
				resInterval = append(resInterval, []int{left, right})
				resInterval = append(resInterval, intervals[index:]...)
				return resInterval
			} else if right <= interval[1] {
				right = interval[1]
			} else { //right>interval[1]
				//do nothing
			}
		} else if left <= interval[1] {
			if right <= interval[1] {
				resInterval = append(resInterval, intervals[index:]...)
				return resInterval
			} else { //right>interval[1]
				left = interval[0]
			}
		} else { //left>interval[1]
			resInterval = append(resInterval, intervals[index])
		}
		if index == len(intervals)-1 {
			resInterval = append(resInterval, []int{left, right})
		}
	}
	return resInterval
}
