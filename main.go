package main

import "fmt"

func main() {

	l1 := []int{9, 0, 1} ///// Numbers can be changed
	l2 := []int{5, 6, 4} ///// Numbers can be changed

	run, erro := AddTwoNumbers(l1, l2)
	if erro == "" {
		fmt.Println(run)
	} else {
		fmt.Println(erro)
	}
}

func AddTwoNumbers(l1 []int, l2 []int) ([]int, string) {
	var errors string = ""
	var result []int

	ln1 := len(l1)
	ln2 := len(l2)

	if ln1 == 0 || ln2 == 0 {
		errors := "The nodes are empty"
		return result, errors
	} else if ln1 > 100 || ln2 > 100 {
		errors := "Nodes should not exceed 100"
		return result, errors
	} else if (ln1 > 1 && l1[ln1-1] == 0) || (ln2 > 1 && l2[ln2-1] == 0) {
		errors := "The lead node must not be zero"
		return result, errors
	} else {
		l1min, l1max := MinMax(l1)
		l2min, l2max := MinMax(l2)
		if l1min < 0 || l1max > 9 || l2min < 0 || l2max > 9 {
			errors := "The node must be 0 <= node <= 9"
			return result, errors
		}
	}

	var st int = 0
	var care int = 0
	var i int

	if ln1 == ln2 {
		st = ln1
	} else if ln1 > ln2 {
		st = ln1
		lns := ln1 - ln2
		for i = 0; i < lns; i++ {
			l2 = append(l2, 0)
		}
	} else {
		st = ln2
		lns := ln2 - ln1
		for i = 0; i < lns; i++ {
			l1 = append(l1, 0)
		}
	}

	for i = 0; i < st; i++ {
		se := l1[i] + l2[i]
		if care > 0 {
			se += care
			care = 0
		}
		if se >= 10 {
			se -= 10
			care = 1
		}
		result = append(result, se)
	}
	if care == 1 {
		result = append(result, 1)
	}
	return result, errors
}

func MinMax(numbers []int) (int, int) {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	min := max
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min, max
}
