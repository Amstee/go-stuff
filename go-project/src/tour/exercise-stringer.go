package tour

import "fmt"

type IPAddr [4]byte

func (this IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", this[0], this[1], this[2], this[3])
}