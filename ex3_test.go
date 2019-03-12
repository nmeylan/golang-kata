package main

import (
	"testing"
)

func TestGame(t *testing.T) {
	for i := 0; i < 1000; i++ {
		win := Game()
		for _, w := range win {
			if !w {
				t.Fail()
				return
			}
		}
	}
}
