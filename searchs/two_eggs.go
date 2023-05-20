package searchs

import "math"

func TwoEggsSearch(breakArray []bool) int {

	n := len(breakArray)
	i := int(math.Sqrt(float64(n)))
	j := int(math.Sqrt(float64(n)))

	for {
		if breakArray[i] {
			for k := i - j + 1; k <= i; k++ {
				if breakArray[k] {
					return k + 1
				}
			}
			break
		} else {
			i = i + j
		}
	}

	return -1
}
