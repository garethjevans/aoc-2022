package iterator

func Fill(start int, end int) []int {
	var a []int

	for x := start; x <= end; x++ {
		a = append(a, x)
	}

	return a
}
