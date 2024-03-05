package piscine

import "github.com/01-edu/z01"

func Printiintrune(n int) {
	c := '0'
	for i := 1; i <= n%10; i++ {
		c++
	}
	for i := -1; i >= n%10; i-- {
		c++
	}
	if n/10 != 0 {
		Printiintrune(n / 10)
	}
	z01.PrintRune(c)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	Printiintrune(n)
}
