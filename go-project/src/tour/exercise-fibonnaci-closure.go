package tour

func fibonacci() func() int {
	a, b := -1, 1
	return func() int {
		if a == -1 {
			a = 0
			return a
		}
		a, b = b, a + b
		return a
	}
}
