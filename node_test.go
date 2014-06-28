package fingerTree

import (
	"reflect"
	"testing"
)

func TestNode2ImplementsNode(test *testing.T) {
	stype := reflect.TypeOf(node2{})
	itype := reflect.TypeOf((*node)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestNode3ImplementsNode(test *testing.T) {
	stype := reflect.TypeOf(node3{})
	itype := reflect.TypeOf((*node)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestNode2Foldl(test *testing.T) {
	n := makeNode2(1, 2)
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 3 {
		test.Error("Expected n.Foldl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(int)))
	}
}

func TestNode3Foldl(test *testing.T) {
	n := makeNode3(1, 2, 3)
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 6 {
		test.Error("Expected n.Foldl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(uint)))
	}
}

func TestNode2Foldr(test *testing.T) {
	n := makeNode2(1, 2)
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 3 {
		test.Error("Expected n.Foldr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(int)))
	}
}

func TestNode3Foldr(test *testing.T) {
	n := makeNode3(1, 2, 3)
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 6 {
		test.Error("Expected n.Foldr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(uint)))
	}
}

func TestNode2Iterr(test *testing.T) {
	n := makeNode2(1, 2)
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 3 {
		test.Error("Expected n.Iterr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}

func TestNode3Iterr(test *testing.T) {
	n := makeNode3(1, 2, 3)
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 6 {
		test.Error("Expected n.Iterr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}

func TestNode2Iterl(test *testing.T) {
	n := makeNode2(1, 2)
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 3 {
		test.Error("Expected n.Iterl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}

func TestNode3Iterl(test *testing.T) {
	n := makeNode3(1, 2, 3)
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 6 {
		test.Error("Expected n.Iterl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}
