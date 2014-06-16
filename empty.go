package fingerTree

type empty struct{}

func (tree empty) Foldl(f FoldFunc, initial Any) Any {
	return initial
}

func (e empty) Foldr(f FoldFunc, initial Any) Any {
	return initial
}

func (tree empty) Pushl(d Any) FingerTree {
	return &single{d}
}

func (e empty) Popl() (FingerTree, Any) {
	return &empty{}, nil
}

func (e empty) Popr() (FingerTree, Any) {
	return &empty{}, nil
}

func (tree empty) Pushr(d Any) FingerTree {
	return &single{d}
}

func (e empty) Iterl(f IterFunc) {
	return
}

func (e empty) Iterr(f IterFunc) {
	return
}

func (e empty) Headr() Any {
	return nil
}

func (e empty) Headl() Any {
	return nil
}

func (e empty) Tailr() FingerTree {
	// Not sure if this makes sense
	return nil
}

func (e empty) Taill() FingerTree {
	return nil
}

func (e empty) IsEmpty() bool {
	return true
}
