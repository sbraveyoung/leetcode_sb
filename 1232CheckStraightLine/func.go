package checkStraightLine

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}
	var ratio float64
	var horizontal bool = false
	if coordinates[1][0] == coordinates[0][0] {
		horizontal = true
	} else {
		ratio = float64(coordinates[1][1]-coordinates[0][1]) / float64(coordinates[1][0]-coordinates[0][0])
	}
	for i := 2; i < len(coordinates); i++ {
		if horizontal {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		} else {
			var nowRatio float64 = float64(coordinates[i][1]-coordinates[0][1]) / float64(coordinates[i][0]-coordinates[0][0])
			if ratio != nowRatio {
				return false
			}
		}
	}
	return true
}
