package piscine

func AlphaCount(s string) int {
	C := 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			C++
		}
	}
	return C
}
