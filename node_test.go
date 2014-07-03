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

func TestNode2Fold(test *testing.T) {
	n := makeNode2(1, 2)
	Fold_Test(test, n, Slice{1, 2})
}

func TestNode3Fold(test *testing.T) {
	n := makeNode3(1, 2, 3)
	Fold_Test(test, n, Slice{1, 2, 3})
}

func TestNode2Iter(test *testing.T) {
	n := makeNode2(1, 2)
	Iter_Test(test, n, Slice{1, 2})
}

func TestNode3Iter(test *testing.T) {
	n := makeNode3(1, 2, 3)
	Iter_Test(test, n, Slice{1, 2, 3})
}
