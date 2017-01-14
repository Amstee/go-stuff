package tour

type MyReader struct{}

func (m MyReader) Read(param []byte) (int, error) {
	count := 0
	for i := 0; i < len(param); i++ {
		param[i] = 'A'
		count += 1
	}
	return count, nil
}
