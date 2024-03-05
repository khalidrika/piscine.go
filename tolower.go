package piscine

func ToLower(s string) string {
	char := []rune(s)
	for i := range char {
		if char[i] >= 'A' && char[i] <= 'Z' {
			char[i] = char[i] + 32
		}
	}
	return string(char)
}
