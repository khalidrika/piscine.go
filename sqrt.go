package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	} else if nb <= 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			if i*i == nb {
				return i
			}
		}
	}
	return 0
}
