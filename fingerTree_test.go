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

func TestEmptyImplementsFingerTree(test *testing.T) {
	stype := reflect.TypeOf(empty{})
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

func TestEmptyPushf(test *testing.T) {
	v := empty{}.Pushf(1)
	if SliceEqual(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushf(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptyPushb(test *testing.T) {
	v := empty{}.Pushb(1)
	if SliceEqual(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushb(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestSinglePushf(test *testing.T) {
	n := &single{1}
	r := n.Pushf(2)
	if SliceEqual(ToSlice(r), []Any{2, 1}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushf(2) to result in sequence [2 1], got %v", ToSlice(r)))
	}
}

func TestFTreePushf(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushf(i)
	}

	if SliceEqual(ToSlice(n), []Any{7, 6, 5, 4, 3, 2, 1, 0}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [7 6 5 4 3 2 1 0], got %v", ToSlice(n)))
	}
}

func TestFTreePushb(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushb(i)
	}

	if SliceEqual(ToSlice(n), []Any{0, 1, 2, 3, 4, 5, 6, 7}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [0 1 2 3 4 5 6 7], got %v", ToSlice(n)))
	}
}
