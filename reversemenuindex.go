package piscine

func ReverseMenuIndex(menu []string) []string {
	arg := len(menu) - 1
	var arey []string
	for i := arg; i >= 0; i-- {
		arey = append(arey, menu[i])
	}
	return arey
}
