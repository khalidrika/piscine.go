package piscine

func CountIf(f func(string) bool, tab []string) int {
	a := 0
	for y := 0; y < len(tab); y++ {
		if f(tab[y]) {
			a++
		}
	}
	return a
}
