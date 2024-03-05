package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	v := make(map[string]int)
	var a string
	for _, khalid := range str {
		if khalid == 32 {
			v[a] += 1
			a = ""
		} else if khalid != 32 {
			a += string(byte(khalid))
		}
	}
	v[a] += 1

	return v
}
