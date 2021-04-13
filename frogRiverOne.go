package solution

func Solution(X int, A []int) int {
	leaves := make(map[int]bool)
	for second, pos := range A {
		if pos > X {
			continue
		}
		leaves[pos] = true
		if len(leaves) == X {
			return second
		}
	}

	return -1
}
