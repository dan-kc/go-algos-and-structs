package searchs

import (
	"math"
)

func TwoEggsSearch(breakArr []bool) int {
	// break array example = [true, true, true, true, true]

	var (
		totalFloors int // 5
		inc         int // 2
		i           int // 2
	)

	totalFloors = len(breakArr)
	inc = int(math.Sqrt(float64(totalFloors)))
	i = inc

	if breakArr[totalFloors-1] == false { // Eggs never break
		return -1
	}

	if breakArr[0] == true { // Eggs breaks on ground floor
		return 0
	}

	if !isSorted(breakArr) { // Array unsorted
		return -1
	}

	for {
		if breakArr[i] {
			for j := i - inc + 1; j <= i; j++ {
				if breakArr[j] {
					return j
				}
			}
			//
		} else {
			i = i + inc
		}
	}
}

func isSorted(arr []bool) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if !arr[i] && arr[i-1] {
			return false
		}
	}
	return true
}
