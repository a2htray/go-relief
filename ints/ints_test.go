package ints

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	Reverse(s)
	fmt.Println(s)

	s = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	Reverse(s)
	fmt.Println(s)
}
