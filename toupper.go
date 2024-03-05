package piscine

func ToUpper(s string) string {
	char := []rune(s)
	for i := range char {
		if char[i] >= 'a' && char[i] <= 'z' {
			char[i] = char[i] - 32
		}
	}
	return string(char)
}
