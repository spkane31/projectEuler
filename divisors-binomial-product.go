package main

import (
	"fmt"
	"math"
	"time"
)

// BinomialProduct : Product of binomials
func BinomialProduct(n int) int {
	sum := 1
	for i := 1; i < n; i++ {
		sum *= B(n, i)
	}

	return sum
}

// B : binomial formula
func B(n, k int) int {
	return fact(n) / (fact(k) * fact(n-k))
}

func fact(n int) int {
	if n < 2 {
		return 1
	} else if n == 2 {
		return n
	} else {
		ret := 2
		for i := 3; i <= n; i++ {
			ret *= i
		}
		return ret
	}
}

// S : sum of divisors
func S(n int) int {
	var sum int
	for i := 1; i < n; i++ {
		sum += sumDivisors(BinomialProduct(i))

	}
	return sum
}

func sumDivisors(n int) int {
	var sum int
	for i := 1; i < int(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 {
			sum += i + n/i
		}
	}
	return sum
}

func main() {
	start := time.Now()

	MOD := 1000000007

	fmt.Println(BinomialProduct(5) % MOD)
	fmt.Println(S(5))
	fmt.Println(S(10))
	fmt.Println(S(100))

	fmt.Printf("Finished in %s\n", time.Since(start))
}
