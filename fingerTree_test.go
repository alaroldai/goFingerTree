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

func TestEmptyFoldl(test *testing.T) {
	n := &empty{}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 0 {
		test.Error("Expected n.Foldl to return 0, got " + string(r.(int)))
	}
}

func TestNode2Foldr(test *testing.T) {
	n := node2{[2]Data{1, 2}}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 3 {
		test.Error("Expected n.Foldr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(int)))
	}
}

func TestNode3Foldr(test *testing.T) {
	n := node3{[3]Data{1, 2, 3}}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 6 {
		test.Error("Expected n.Foldr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(r.(uint)))
	}
}

func TestSingleFoldr(test *testing.T) {
	n := &single{1}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 1 {
		test.Error("Expected n.Foldr to return 1, got " + string(r.(int)))
	}
}

func TestEmptyFoldr(test *testing.T) {
	n := &empty{}
	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 0 {
		test.Error("Expected n.Foldr to return 0, got " + string(r.(int)))
	}
}

func TestNode2Iterr(test *testing.T) {
	n := node2{[2]Data{1, 2}}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 3 {
		test.Error("Expected n.Iterr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}

func TestNode3Iterr(test *testing.T) {
	n := node3{[3]Data{1, 2, 3}}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 6 {
		test.Error("Expected n.Iterr(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}

func TestSingleIterr(test *testing.T) {
	n := &single{1}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 1 {
		test.Error("Expected n.Iterr to return 1, got " + string(sum))
	}
}

func TestEmptyIterr(test *testing.T) {
	n := &empty{}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 0 {
		test.Error("Expected n.Iterr to return 0, got " + string(sum))
	}
}

func TestNode2Iterl(test *testing.T) {
	n := node2{[2]Data{1, 2}}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 3 {
		test.Error("Expected n.Iterl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}

func TestNode3Iterl(test *testing.T) {
	n := node3{[3]Data{1, 2, 3}}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 6 {
		test.Error("Expected n.Iterl(func (a, b uint) { return a + b }, 0) to return 3, got " + string(sum))
	}
}

func TestSingleIterl(test *testing.T) {
	n := &single{1}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 1 {
		test.Error("Expected n.Iterl to return 1, got " + string(sum))
	}
}

func TestEmptyIterl(test *testing.T) {
	n := &empty{}
	sum := 0
	add := func(b Data) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 0 {
		test.Error("Expected n.Iterl to return 0, got " + string(sum))
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

func TestEmptyPushl(test *testing.T) {
	v := empty{}.Pushl(1)
	if cmpslices(ToSlice(v), []Data{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushl(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptyPopl(test *testing.T) {
	n := empty{}
	r, e := n.Popl()

	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error(fmt.Sprintf("Expected n.Popl() result to be an empty node, got %v", r))
	}
	if e != nil {
		test.Error("Expected n.Popl() result to be nil")
	}
}

func TestEmptyPushr(test *testing.T) {
	v := empty{}.Pushr(1)
	if cmpslices(ToSlice(v), []Data{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushr(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestSinglePushl(test *testing.T) {
	n := &single{1}
	r := n.Pushl(2)
	if cmpslices(ToSlice(r), []Data{2, 1}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushl(2) to result in sequence [2 1], got %v", ToSlice(r)))
	}
}

func TestSinglePopl(test *testing.T) {
	n := &single{1}
	r, e := n.Popl()
	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error("Expected n.Popl() result to be an empty node")
	}
	if e != 1 {
		test.Error("Expected n.Popl() result to be 1")
	}
}

func TestSinglePushr(test *testing.T) {
	n := &single{1}
	r := n.Pushr(2)
	if cmpslices(ToSlice(r), []Data{1, 2}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushr(2) to result in sequence [1 2], got %v", ToSlice(r)))
	}
}

func TestFTreeFoldl(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 45 {
		test.Error("Expected n.Foldl to return 1, got " + string(r.(int)))
	}
}

func TestFTreeFoldr(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushr(i)
	}

	add := func(a interface{}, b Data) interface{} {
		return interface{}(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 45 {
		test.Error("Expected n.Foldl to return 1, got " + string(r.(int)))
	}
}

func TestFTreePushl(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushl(i)
	}

	if cmpslices(ToSlice(n), []Data{7, 6, 5, 4, 3, 2, 1, 0}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [7 6 5 4 3 2 1 0], got %v", ToSlice(n)))
	}
}

func TestFTreePopl(test *testing.T) {
	var n FingerTree = &empty{}

	for i := 0; i < 20; i++ {
		n = n.Pushl(i)
	}

	var e Data
	for i := 19; i >= 0; i-- {
		n, e = n.Popl()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.Pushr(i)
	}
	for i := 0; i < 22; i++ {
		n, e = n.Popl()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreePushr(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushr(i)
	}

	if cmpslices(ToSlice(n), []Data{0, 1, 2, 3, 4, 5, 6, 7}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [0 1 2 3 4 5 6 7], got %v", ToSlice(n)))
	}
}

func TestFTreeIterl(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	sum := 0
	add := func(d Data) {
		sum += d.(int)
	}
	n.Iterl(add)
	if sum != 45 {
		test.Error(fmt.Sprintf("Expected n.Iterl to result in sum 110, got %v", sum))
	}
}

func TestFTreeIterr(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	sum := 0
	add := func(d Data) {
		sum += d.(int)
	}
	n.Iterr(add)
	if sum != 45 {
		test.Error(fmt.Sprintf("Expected n.Iterl to result in sum 110, got %v", sum))
	}
}
