package numbers

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

func TestDot(t *testing.T) {
	dtl := Dot([]int{1, 2, 3, 4}, []int{4, 3, 2, 1})
	if dtl != 20 {
		t.Error(fmt.Printf("got %d, correct %d", dtl, 20))
	}
}
