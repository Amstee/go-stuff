package tour

import (
	"strings"
)

func WordCount(s string) (container map[string]int) {
	fields := strings.Fields(s)
	container = make(map[string]int)

	for i := 0; i < len(fields); i++ {
		if _, ok := container[fields[i]]; ok == true {
			container[fields[i]] += 1
		} else {
			container[fields[i]] = 1
		}
	}
	return
}