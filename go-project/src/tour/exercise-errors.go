package tour

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", int(e))
}

func Sqrt(x float64) (z float64, err error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, err
	}
	z = 1.0
	for i := 0; i < 20; i++ {
		z = z - (z * z - x) / (2 * z)
	}
	return z, nil
}