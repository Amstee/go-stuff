package tour

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(param []byte) (int, error) {
	si, err := reader.r.Read(param)

	for i := 0; i < si; i++ {
		if ((param[i] + 13) <= 'z') {
			param[i] = param[i] + 13
		} else {
			param[i] = param[i] - 13
		}
	}
	return si, err
}
