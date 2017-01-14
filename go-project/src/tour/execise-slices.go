package tour

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
		for u := 0; u < dx; u++ {
			pic[i][u] = uint8((u + i) / 2)
		}
	}
	return pic
}
