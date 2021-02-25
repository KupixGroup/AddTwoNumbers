package main

import (
	"reflect"
	"testing"
)

func TestMinMax(t *testing.T) {
	testTable := []struct {
		numbers []int
		min     int
		max     int
	}{
		{
			numbers: []int{3, -5, 9, 2},
			min:     -5,
			max:     9,
		}, {
			numbers: []int{3, 5, 8, 2, 1},
			min:     1,
			max:     8,
		},
	}

	for _, utest := range testTable {
		min, max := MinMax(utest.numbers)
		t.Logf("Error TestMinMax() function")
		if min != utest.min || max != utest.max {
			t.Errorf("array in %d, (min: %d, max: %d) error", utest.numbers, utest.min, utest.max)
		}
	}

}

func TestAddTwoNumbers(t *testing.T) {

	testTable := []struct {
		num1 []int
		num2 []int
		out  []int
	}{
		{
			num1: []int{2, 4, 3},
			num2: []int{5, 6, 4},
			out:  []int{7, 0, 8},
		}, {
			num1: []int{9, 9, 9, 9, 9, 9, 9},
			num2: []int{9, 9, 9, 9},
			out:  []int{8, 9, 9, 9, 0, 0, 0, 1},
		}, {
			num1: []int{0},
			num2: []int{0},
			out:  []int{0},
		},
	}
	for _, utest := range testTable {
		outs, erro := AddTwoNumbers(utest.num1, utest.num2)
		if erro == "" {
			t.Logf("Error TestAddTwoNumbers() function")
			if reflect.DeepEqual(outs, utest.out) == false {
				t.Errorf("ERROR: %d + %d, out: %d", utest.num1, utest.num2, utest.out)
			}
		} else {
			t.Logf("Error TestAddTwoNumbers() function")
			t.Errorf("Error writing node. (%d + %d, out: %d). ERROR: "+erro, utest.num1, utest.num2, utest.out)
		}
	}

}
