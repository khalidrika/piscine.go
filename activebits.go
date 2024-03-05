package piscine

// import "fmt"

func ActiveBits(n int) int {
	a := 0
	for n > 0 {
		if n%2 != 0 {
			a++
		}
		n = n / 2
	}
	return a
}

// func main() {
// 	fmt.Println(ActiveBits(8))
// }
