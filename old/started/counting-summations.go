package main

import (
	"fmt"
	"sort"
)

func expand(sum []int) []list { //b[]int {
	var ret []list

	temp := []int{}

	for i := range sum {
		if sum[i] != 1 {
			temp = append(temp, sum[i]-1, 1)
			for j := i; j < len(sum); j++ {
				temp = append(temp, sum[j])
				ret = append(ret, temp)
				temp = []int{}
			}
		} else {
			temp = append(temp, sum[i])
		}
	}

	return ret
}

func isEqual(l1, l2 []int) bool {
	length := len(l1)
	if len(l2) < length {
		length = len(l2)
	}

	for i := 0; i < length; i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func filter(pairs [][]int) [][]int {
	for i := 0; i < len(pairs); i++ {
		if i >= len(pairs)-1 {
			break
		}
		if isEqual(pairs[i], pairs[i+1]) {
			copy(pairs[i:], pairs[i+1:])
			pairs[len(pairs)-1] = nil
			pairs = pairs[:len(pairs)-1]
		}
	}
	return pairs
}

func countSummations(n int) int {
	// sums := pairs //[][]int{}
	var sums pairs

	for i := n - 1; i >= n/2; i-- {
		sums = append(sums, []int{i, n - i})
	}
	if n%2 == 1 {
		sums = sums[:len(sums)-1]
	}
	// if n % 2 == 1 {
	//   sums = append(sums, []int{n/2, n/2, 1})
	// }

	fmt.Println(sums)

	for _, sum := range sums {
		for i := range sum {
			if sum[i] != 1 {
				fmt.Println(sum)
				new := expand(sum)
				for _, sum := range new {
					sort.Sort(sum)
				}
				fmt.Println(new)
				for _, sum := range new {
					sums = append(sums, sum)
				}
				// fmt.Println(expand(sum))
			}
		}
	}

	sort.Sort(sums)
	fmt.Println(sums)

	fmt.Println(filter(sums))

	return len(sums)
}

type list []int

func (l list) Len() int           { return len(l) }
func (l list) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l list) Less(i, j int) bool { return l[i] > l[j] }

type pairs [][]int

func (p pairs) Len() int { return len(p) }

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	less := len(p[i])
	if len(p[j]) < less {
		less = len(p[j])
	}

	for k := 0; k < less; k++ {
		if p[i][k] != p[j][k] {
			return p[i][k] < p[j][k]
		}
	}
	return false
}

func main() {
	// fmt.Println(amicablePairsUpTo(300))
	// fmt.Println(amicablePairsUpTo(3000))
	// pairs := amicablePairsUpTo(10000)
	// sum := 0
	// for _, pair := range pairs {sum += pair[0] + pair[1]}
	// fmt.Println(amicablePairsUpTo(20000))
	// fmt.Println(sum)
	// fmt.Println(check(1020304050647080900))
	// fmt.Println(concealedSquare())
	// fmt.Println(check(1389019170 * 1389019170))
	// fmt.Println(numReplacements(56003))
	fmt.Println(countSummations(5))
	// fmt.Println(countSummations(4))
	// fmt.Println(countSummations(3))
	// fmt.Println(countSummations(2))
}
