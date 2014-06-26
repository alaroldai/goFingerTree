package fingerTree

type single struct {
	metadata map[string]Any
	data     Any
}

func makeSingle(d Any, mdataTypes mdataTypeMap) *single {
	meta := make(map[string]Any)
	for k, v := range mdataTypes {
		sz := v.unit
		dn, succ := d.(mdataContainer)
		if succ {
			sz = dn.mdataForKey(k, mdataTypes)
		}
		meta[k] = sz
	}
	return &single{meta, d}
}

func (s *single) mdataForKey(key string, mdataTypes mdataTypeMap) Any {
	return s.metadata[key]
}

func (s *single) Foldl(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) Foldr(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) Pushl(d Any, mdataTypes mdataTypeMap) FingerTreeComponent {
	return makeFTreeTriple(
		[]Any{d},
		makeEmpty(),
		[]Any{s.data},
		mdataTypes,
	)
}

func (s *single) Popl(mdTypes mdataTypeMap) (FingerTreeComponent, Any) {
	return makeEmpty(), s.data
}

func (s *single) Popr(mdTypes mdataTypeMap) (FingerTreeComponent, Any) {
	return makeEmpty(), s.data
}

func (s *single) Pushr(d Any, mdataTypes mdataTypeMap) FingerTreeComponent {
	return makeFTreeTriple(
		[]Any{s.data},
		makeEmpty(),
		[]Any{d},
		mdataTypes,
	)
}

func (s *single) Iterl(f IterFunc) {
	f(s.data)
}

func (s *single) Iterr(f IterFunc) {
	f(s.data)
}

func (s *single) Headr() Any {
	return s.data
}

func (s *single) Headl() Any {
	return s.data
}

func (s *single) Tailr(mdTypes mdataTypeMap) FingerTreeComponent {
	return makeEmpty()
}

func (s *single) Taill(mdTypes mdataTypeMap) FingerTreeComponent {
	return makeEmpty()
}

func (s *single) IsEmpty() bool {
	return false
}

// Concat t to the right of the receiver
func (s *single) Concatr(t FingerTreeComponent, mdTypes mdataTypeMap) FingerTreeComponent {
	return t.Pushl(s.data, mdTypes)
}

func (s *single) Concatl(t FingerTreeComponent, mdTypes mdataTypeMap) FingerTreeComponent {
	return t.Pushr(s.data, mdTypes)
}

func (s *single) ToSlice() Slice {
	return ToSlice(s)
}
