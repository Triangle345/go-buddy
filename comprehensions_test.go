package comprehensions

import (
	"testing"
)

func TestInts(t *testing.T) {

	mylist := []int{5, 1, 4, 100}
	val := Reduce(mylist, 10, func(acc int, num int) int { return acc + num })
	if val != 120 {
		t.Fatal("reduce failed:", val)
	}
}
