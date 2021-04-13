package maxCounters

func Solution(N int, A []int) []int {
	counters := make([]int, N)
	maxValue := 0
	base := 0
	for _, value := range A {
		if value <= N {
			index := value - 1
			if counters[index] < base {
				counters[index] = base + 1
			} else {
				counters[index]++
			}
			maxValue = max(maxValue, counters[index])
		} else {
			base = maxValue
		}
	}

	for index, counter := range counters {
		if counter >= base {
			continue
		}

		counters[index] = base
	}
	return counters
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
