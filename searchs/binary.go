package searchs

func BinarySearch(stack []int, target int) int {
	low := 0
	high := len(stack) - 1

	for low <= high {
		median := (high + low) / 2
		if stack[median] < target {
			low = median + 1
		} else if stack[median] > target {
			high = median - 1
		} else {
			return median
		}
	}
	return -1
}
