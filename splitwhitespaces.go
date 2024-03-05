package piscine

// import "fmt"

func SplitWhiteSpaces(s string) []string {
	x := []rune(s)
	var y []string
	n := ""
	for i := 0; i < len(x); {
		if x[i] != ' ' && x[i] != '\n' && x[i] != '\t' {
			n += string(x[i])
			i++
		} else {
			if n != "" {
				y = append(y, n)
			}
			n = ""
			for i < len(x) && (x[i] == ' ' || x[i] == '\n' || x[i] == '\t') {
				i++
			}
		}
	}
	if n != "" {
		y = append(y, n)
	}
	return y
}

// func main() {
// 	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))}
