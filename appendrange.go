package piscine

func AppendRange(min, max int) []int {
	var x []int
	if min > max {
		return x
	}
	for i := min; i < max; i++ {
		x = append(x, i)
	}
	return x
}
