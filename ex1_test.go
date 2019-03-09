package main

import (
	"fmt"
	"testing"
)

type myType int

func TestReverse(t *testing.T) {
	var slice = []myType{0, 1, 2, 3, 4}
	var expectedSlice1 = []myType{4, 3, 2, 1, 0}
	Reverse(slice)
	for i, v := range expectedSlice1 {
		if slice[i] != v {
			fmt.Print("Got: ")
			fmt.Println(slice)
			fmt.Print("Expected: ")
			fmt.Println(expectedSlice1)
			t.Fail()
			return
		}
	}
}
