package set

// OfInts creates set using map of ints.
func OfInts(i ...int) map[int]bool {
	out := map[int]bool{}
	for _, v := range i {
		out[v] = true
	}
	return out
}

// OfInts64 creates set using map of int64.
func OfInts64(i ...int64) map[int64]bool {
	out := map[int64]bool{}
	for _, v := range i {
		out[v] = true
	}
	return out
}

// OfUints64 creates set using map of uint64.
func OfUints64(i ...uint64) map[uint64]bool {
	out := map[uint64]bool{}
	for _, v := range i {
		out[v] = true
	}
	return out
}
