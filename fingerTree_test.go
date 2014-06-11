package fingerTree23

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func MethodsMissingFromType(inter, typ reflect.Type) []string {
	missingMethods := make([]string, 0)
	for n := 0; n < inter.NumMethod(); n++ {
		_, present := typ.MethodByName(inter.Method(n).Name)
		if !present {
			fmt.Println(inter.Method(n).Name)
			missingMethods = append(missingMethods, inter.Method(n).Name)
		}
	}
	return missingMethods
}

func TypeConformityTest(test *testing.T, stype, itype reflect.Type) {
	if !stype.Implements(itype) {
		missingMethods := MethodsMissingFromType(itype, stype)
		test.Error("struct '" + stype.Name() + "' does not implement interface '" + itype.Name() + "' (missing methods: " + strings.Join(missingMethods, ", ") + ")")
	}
}

func TestFTreeImplementsFingerTree(test *testing.T) {
	stype := reflect.TypeOf(ftree{})
	itype := reflect.TypeOf((*FingerTree)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSingleImplementsFingerTree(test *testing.T) {
	stype := reflect.TypeOf(single{})
	itype := reflect.TypeOf((*FingerTree)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

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
	n := node2{[2]Data{1, 2}}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 3 {
		test.Error("Expected n.Foldl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(int)))
	}
}

func TestNode3Foldl(test *testing.T) {
	n := node3{[3]Data{1, 2, 3}}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 6 {
		test.Error("Expected n.Foldl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(uint)))
	}
}

func TestSingleFoldl(test *testing.T) {
	n := &single{1}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 1 {
		test.Error("Expected n.Foldl to return 1, got " + string(r.(int)))
	}
}

func cmpslices(a, b []Data) bool {
	if len(a) != len(b) {
		fmt.Println("Lengths differ")
		return false
	}
	for i, v := range a {
		if v != b[i] {
			fmt.Println("Item at index ", i, " differs")
			return false
		}
	}
	return true
}

func TestSinglePushf(test *testing.T) {
	n := &single{1}
	r := n.Pushf(2)
	if cmpslices(ToSlice(r), []Data{2, 1}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushf(2) to result in sequence [2 1], got %v", ToSlice(r)))
	}
}

func TestFTreePushf(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushf(i)
	}

	if cmpslices(ToSlice(n), []Data{7, 6, 5, 4, 3, 2, 1, 0}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [7 6 5 4 3 2 1 0], got %v", ToSlice(n)))
	}
}

func TestFTreePushb(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushb(i)
	}

	if cmpslices(ToSlice(n), []Data{0, 1, 2, 3, 4, 5, 6, 7}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [0 1 2 3 4 5 6 7], got %v", ToSlice(n)))
	}
}
