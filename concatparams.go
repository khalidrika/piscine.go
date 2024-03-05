package piscine

func ConcatParams(args []string) string {
	x := ""
	for i, y := range args {
		x += y
		if i < len(args)-1 {
			x += "\n"
		}
	}
	return x
}
