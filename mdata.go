package fingerTree

type MetaMap struct {
	items map[string]Monoid
}

func (m MetaMap) Plus(rm Monoid) Monoid {
	if rm == Zero {
		return m
	}

	r := rm.(MetaMap)

	// merge the two maps with Plus
	merged := make(map[string]Monoid)
	for k, v1 := range m.items {
		merged[k] = v1
	}
	for k, v2 := range r.items {
		if merged[k] != nil {
			merged[k] = merged[k].Plus(v2)
		} else {
			merged[k] = v2
		}
	}

	return MetaMap{merged}
}
