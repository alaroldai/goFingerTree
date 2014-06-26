package fingerTree

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var mdataStandardTypes mdataTypeMap = mdataTypeMap{
	ft_size_key: mdataField{
		0,
		1,
		func(a, b Any) Any {
			return a.(int) + b.(int)
		},
	},
}

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

func cmpslices(a, b []Any) bool {
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
