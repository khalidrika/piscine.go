package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	x := true
	r := true

	for y := 0; y < len(a)-1; y++ {
		if f(a[y], a[y+1]) > 0 {
			x = false
		}
		if f(a[y], a[y+1]) < 0 {
			r = false
		}
	}

	if r {
		return true
	}
	return x
}
