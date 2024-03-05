package piscine

func UltimateDivMod(a *int, b *int) {
	d := *a
	c := *b
	*a = d / c
	*b = d % c
}
