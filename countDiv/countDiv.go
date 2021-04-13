package countDiv

func Solution(A int, B int, K int) int {
	count := B/K - A/K
	if A%K == 0 {
		count += 1
	}
	return count
}
