package collatz_conjecture

import "errors"

func Conjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}
	var passos int = 0
	var result int = n

	for result != 1 {
		passos++

		if result%2 == 0 {
			result = result / 2
		} else {
			result = 3*result + 1
		}
	}

	return passos, nil
}
