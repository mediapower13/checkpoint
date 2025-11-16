package main

// func isPrime(n int) bool {
// 	if n < 2 {
// 		return false
// 	}
// 	if n%2 == 0 {
// 		return n == 2
// 	}
// 	for i := 3; i*i <= n; i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func FindPrevPrime(nb int) int {
	for n := nb; n >= 2; n-- {
		if isPrime(n) {
			return n
		}
	}
	return 0
}
