package fingerTree23

func Foldl(f FoldFunc, initial Any, s []Any, length int) Any {
	if length > 0 {
		return f(Foldl(f, initial, s[:length-1], length-1), s[length-1])
	}
	return initial
}

func Foldr(f FoldFunc, initial Any, s []Any, length int) Any {
	if length > 0 {
		return f(Foldr(f, initial, s[1:], length-1), s[0])
	}
	return initial
}
