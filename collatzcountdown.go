package piscine

func CollatzCountdown(start int) int {
	a := 0
	if start <= 0 {
		return -1
	}
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = start*3 + 1
		}
		a++
	}
	return a
}
