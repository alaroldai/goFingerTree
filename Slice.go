package fingerTree23

func SliceEqual(a, b []Any) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ToSlice(t Foldable) []Any {
	app := func(a Any, b Any) Any {
		return append(a.([]Any), b)
	}
	return t.Foldl(app, make([]Any, 0)).([]Any)
}
