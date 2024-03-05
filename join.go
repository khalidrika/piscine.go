package piscine

func Join(strs []string, sep string) string {
	st := ""
	for a, valor := range strs {
		if len(strs)-1 == a {
			st += valor
		} else {
			st += valor + sep
		}
	}
	return st
}
