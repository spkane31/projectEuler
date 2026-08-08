package main

import (
	"fmt"
)

// HCF :
func HCF(n, d int) int {
	for d != 0 {
		n, d = d, n%d
	}
	return n
}

func Between(n, d int) bool {
	if float64(n)/float64(d) >= .5 {
		return false
	}
	if float64(n)/float64(d) <= 1.0/3.0 {
		return false
	}
	return true
}

func countFracs(d int) int {
	count := 0
	for d > 2 {
		for n := d / 3; n < d/2+1; n++ {
			if HCF(n, d) == 1 && Between(n, d) {
				// fmt.Println(n, d, HCF(n, d))
				count++
			}
		}
		d--
	}

	return count
}

func main() {
	fmt.Println(countFracs(12000))
}
