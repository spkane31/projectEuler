package main

import (
	"fmt"
	"math"
	"time"
)

var primes map[int]bool

// IsGeneratingInteger :
func IsGeneratingInteger(n int) bool {
	for d := 1; d < int(math.Sqrt(float64(n))+1); d++ {
		if n%d == 0 {
			fmt.Println(d, n/d)
			if !IsPrime(d + n/d) {
				return false
			} else {
				r := n / d
				if !IsPrime(r + n/r) {
					return false
				}
			}
		}
	}
	return true
}

// IsPrime : test for primality
func IsPrime(n int) bool {
	if val, has := primes[n]; has {
		return val
	} else {
		for i := 3; i < int(math.Sqrt(float64(n))+1); i++ {
			if n%i == 0 {
				primes[n] = false
				return false
			}
		}
	}
	primes[n] = true
	return true
}

func main() {
	start := time.Now()
	fmt.Println("Starting")
	primes = make(map[int]bool)
	primes[0] = false
	primes[1] = false
	primes[2] = true

	// fmt.Println(IsGeneratingInteger(30))
	// // os.Exit(1)

	sum := 0
	for i := 1; i < 100000000; i++ {
		if IsGeneratingInteger(i) {
			sum += i
			// fmt.Println("New sum: ", sum)
		}
		if i%10000000 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(sum)

	fmt.Printf("Took %s\n", time.Since(start))

}
