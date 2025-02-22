package comprehensions

import (
	"testing"
)

var intlist []int = []int{5, 1, 4, 100}

func TestReducePrimitive(t *testing.T) {

	val := Reduce(intlist, 10, func(acc int, num int) int { return acc + num })
	if val != 120 {
		t.Fatal("reduce failed:", val)
	}
}

func TestFilterPrimitive(t *testing.T) {

	val := Filter(intlist, func(val int) bool { return val == 4 })
	if len(val) != 1 {
		t.Fatal("filter failed, incorrect size:", val)
	}

	if val[0] != 4 {
		t.Fatal("filter failed, incorrect value:", val)
	}
}

func TestMapPrimitive(t *testing.T) {

	val := Map(intlist, func(val int) int { return val * 2 })
	if len(val) != len(intlist) {
		t.Fatal("filter failed, incorrect size:", val)
	}

	if val[0] != 10 {
		t.Fatal("filter failed, incorrect value:", val)
	}

	if val[len(intlist)-1] != 200 {
		t.Fatal("filter failed, incorrect value:", val)
	}
}

func TestFindPrimitive(t *testing.T) {

	val, found := Find(intlist, func(val int) bool { return val == 100 })
	if found != true {
		t.Fatal("filter failed, incorrect size:", val)
	}

	if val != 100 {
		t.Fatal("filter failed, incorrect value:", val)
	}

}
