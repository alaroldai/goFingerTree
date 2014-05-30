package fingerTree23

func Foldl(f FoldFunc, initial Data, s []Data, length int) Data {
	if length > 0 {
		return f(Foldl(f, initial, s[:length-1], length-1), s[length-1])
	}
	return initial
}
