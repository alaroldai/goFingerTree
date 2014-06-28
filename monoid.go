package fingerTree

type Monoid interface {
	Plus(right Monoid) Monoid
}

// Should never be instantiated except for here
type _mzero struct {}
func (m _mzero) Plus(foo Monoid) Monoid {
	return foo
}

var Zero = _mzero{}
