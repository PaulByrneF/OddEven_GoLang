package main

import "fmt"

func main() {

	n := []int{}
	oe := []string{}

	for i := range [11]int{} {
		n = append(n, i)
	}

	for i := range n {
		if i%2 == 0 {
			oe = append(oe, "Even")
		}
		if i%2 == 1 {
			oe = append(oe, "Odd")
		}
	}

	for i := range oe {
		fmt.Println(i, "is", oe[i])
	}
}
