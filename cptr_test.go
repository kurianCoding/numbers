package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	ttl := Sum([]int{1, 2, 3, 4, 5})
	if ttl != 15 {
		t.Error(fmt.Printf("got %d, correct %d", ttl, 15))
	}
}
