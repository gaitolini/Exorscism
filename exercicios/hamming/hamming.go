package hamming

import (
	"fmt"
)

// GAGCCTACTAACGGGAT
// CATCGTAATGACGGCCT
// ^ ^ ^  ^ ^    ^^

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("strings of different lengths")
	}

	counter := 0
	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
