package fingerTree

const ft_size_key string = "ft_size"

type mdataField struct {
	identity Any
	unit     Any

	compose func(a, b Any) Any
	// compare  func(a, b Any) int // Should return -1 if a < b, +1 if a > b, and 0 if a == b
}

type mdataTypeMap map[string]mdataField

// var mdataTypes map[string]mdataField = map[string]mdataField{
// 	ft_size_key: mdataField{
// 		0,
// 		1,
// 		func(a, b Any) Any {
// 			return a.(int) + b.(int)
// 		},
// 	},
// }

func mdataComposeFromSliceWithKey(a Slice, k string, mdataTypes mdataTypeMap) Any {
	compose := mdataTypes[k].compose
	unit := mdataTypes[k].unit
	identity := mdataTypes[k].identity

	return a.Foldl(func(i Any, a Any) Any {
		an, succ := a.(mdataContainer)
		if succ {
			return compose(i, an.mdataForKey(k, mdataTypes))
		}
		return compose(i, unit)
	}, identity)
}

type mdataContainer interface {
	mdataForKey(key string, mdataTypes mdataTypeMap) Any
}
