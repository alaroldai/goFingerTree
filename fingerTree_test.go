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
