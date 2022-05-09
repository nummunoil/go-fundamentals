package cal

func sumOfFirst(n int) int {
	// return n
	// return 6
	sum := 0
	for i := n; i > 0; i-- {
		sum = sum + i
	}
	return sum
}
