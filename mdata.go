package fingerTree

const ft_size_key string = "ft_size"

type mdataField struct {
	identity Any
	unit     Any
	compose  func(a, b Any) Any
}

var mdataTypes map[string]mdataField = map[string]mdataField{
	ft_size_key: mdataField{
		0,
		1,
		func(a, b Any) Any {
			return a.(int) + b.(int)
		},
	},
}

func mdataComposeFromSliceWithKey(a Slice, k string) Any {
	compose := mdataTypes[k].compose
	unit := mdataTypes[k].unit
	identity := mdataTypes[k].identity

	return a.Foldl(func(i Any, a Any) Any {
		an, succ := a.(mdata)
		if succ {
			return compose(i, an.mdataForKey(k))
		}
		return compose(i, unit)
	}, identity)
}

type mdata interface {
	mdataForKey(key string) Any
}
