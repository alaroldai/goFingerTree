package fingerTree23

func Foldl(f FoldFunc, initial Any, s []Any) Any {
	if len(s) > 0 {
		return f(Foldl(f, initial, s[:len(s)-1]), s[len(s)-1])
	}
	return initial
}
