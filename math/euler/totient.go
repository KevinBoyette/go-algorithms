package euler

// Totient (Euler's Totient Function) counts the positive integers up
// to a given integer n that are 'relatively' prime to n.
// Usually denoted with the greek symbol for phi.
//
// Example:
//
//	Totient(10) == 4
//		(1, 3, 7, 9) are less than 10,
//		but have no common factors with 10
//
// Example:
//
//	Totient(9) == 6
//		(1, 2, 4, 5, 7, 8) are less than 10,
//		but have no common factors with 9
//
// Example:
//
//	Totient(5) == 4
//		(1, 2, 3, 4) are less than 5 and since 5 is prime,
//		all the numbers less than 5 are relatively prime to 5.
func Totient(n uint) uint {
	//  In the spirit of math,
	//  r == resulting int,
	//  n == is a positive int,
	//  and p is a prime.
	r := float64(n)
	p := uint(2)
	for p*p <= n {
		if n%p == 0 {
			for n%p == 0 {
				n /= p
			}
			r *= (1.0 - (1.0 / float64(p)))
		}
		p++
	}
	if n > 1 {
		r *= (1.0 - (1.0 / float64(n)))
	}
	return uint(r)
}
