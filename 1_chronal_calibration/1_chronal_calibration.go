package one_chronal_calibration

import (
	"strconv"
)

var seen map[int]bool

func ChronalCalibrationPartOne(frequencies []string) int {
	total := 0

	for _, freq := range frequencies {
		num, _ := strconv.Atoi(freq)
		total += num
	}

	return total
}

func ChronalCalibrationPartTwo(frequencies []string) int {
	total := 0
	loops := len(frequencies)
	seen = make(map[int]bool)

	i := 0
	for {
		_, in := seen[total]
		if in == true {
			return total
		}
		seen[total] = true

		num, _ := strconv.Atoi(frequencies[i])
		total += num

		i += 1
		if i >= loops {
			i = 0
		}
	}
}
