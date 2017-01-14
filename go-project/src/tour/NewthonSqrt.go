package tour

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 20; i++ {
		z = z - (z * z - x) / (2 * z)
	}
	return z
}