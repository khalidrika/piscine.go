package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
	a := '0'
	b := '0'
	c := '0'
	d := '0'

	for i := 0; i < ptr.x/10; i++ {
		a++
	}
	for i := 0; i < ptr.x%10; i++ {
		b++
	}
	for i := 0; i < ptr.y/10; i++ {
		c++
	}
	for i := 0; i < ptr.y%10; i++ {
		d++
	}
	r := []rune{'x', ' ', '=', ' ', a, b, ',', ' ', 'y', ' ', '=', ' ', c, d}
	for i := 0; i < 14; i++ {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}

func main() {
	points := &point{}
	setPoint(points)
}

// 	/*slis := make([]rune, 2)

// 	slis[0] = rune(ptr.x / 10)
// 	slis[1] = rune(ptr.x % 10)
// 	slis2 := make([]rune, 2)
// 	slis2[0] = rune(ptr.y / 10)
// 	slis2[1] = rune(ptr.y % 10)*/
// 	for i := 0; i < len(slis); i++ {
// 		z01.PrintRune(slis[i] + '0')
// 	}
// 	for i := 0; i < 2; i++ {
// 		z01.PrintRune(slis2[i] + '0')
// 	}
// 	r := []rune{'x', ' ', '=', ' ', slis[0] + '0', slis[1] + '0', ',', ' ', 'y', ' ', '=', ' ', slis2[0] + '0', slis2[1] + '0'}
// 	for i := 0; i < len(r); i++ {
// 		z01.PrintRune(r[i])
// 	}
// 	z01.PrintRune('\n')
// }
