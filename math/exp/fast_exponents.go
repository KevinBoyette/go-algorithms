package exp

// FastExponents uses the squaring method to
// calculate the exponentiation faster than
// the naive method.
func FastExponents(x, y int) int {
	result := 1
	if y == 0 {
		return 1
	} else if y == 1 {
		return x
	}
	for y >= 1 {
		if y%2 == 1 {
			result *= x
		}
		y /= 2
		x *= x
	}

	return result
}
