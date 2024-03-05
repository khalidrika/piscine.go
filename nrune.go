package piscine

func NRune(s string, n int) rune {
	t := []rune(s)
	for i := range t {
		if i == n-1 {
			return rune(t[i])
		}
	}
	return 0
}
