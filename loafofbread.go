package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	y := ""
	for i := 0; i < len(str); i++ {
		if len(str)-i < 5 {
			break
		}
		if str[i] == ' ' {
			i++
		}
		for j := i; j < i+5; j++ {
			if j < len(str) && str[j] == ' ' {
				j++
			}
			if j < len(str) {
				y += string(str[j])
			}
		}
		y += " "
		i = i + 5
	}
	return y
}
