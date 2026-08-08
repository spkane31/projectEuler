package main

import "fmt"

func factorialChain(n int) int {
	chainLength := 1
	m := make(map[int]bool)

	m[n] = true

	next := sumFactorial(n)
	contained := false
	for !contained {
		// fmt.Println(next)
		if _, has := m[next]; has {
			return chainLength
		}
		m[next] = true
		next = sumFactorial(next)
		chainLength++

	}

	return chainLength
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

func sumFactorial(n int) int {
	ret := 0
	for n != 0 {
		ret += fact(n % 10)
		n /= 10
	}
	return ret
}

func main() {
	// fmt.Println(fact(3) + fact(6) + fact(3) + fact(6) + fact(1))
	// fmt.Println(fact(1) + fact(4) + fact(5) + fact(4))
	// fmt.Println(sumFactorial(145))
	fmt.Println(factorialChain(169))

	var count int
	for i := 1; i < 1000000; i++ {
		if factorialChain(i) == 60 {
			count++
			fmt.Println(count)
		}
	}
}
