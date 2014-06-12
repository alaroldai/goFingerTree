package fingerTree23

func Iterl(f IterFunc, s []Data, length int) {
	if length > 0 {
		f(s[0])
		Iterl(f, s[1:], length-1)
	}
}

func Iterr(f IterFunc, s []Data, length int) {
	if length > 0 {
		f(s[length-1])
		Iterr(f, s, length-1)
	}
}
