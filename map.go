package piscine

func Map(f func(int) bool, a []int) []bool {
	mepn := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		mepn[i] = f(a[i])
	}
	return mepn
}
