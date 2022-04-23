package set

// OfStrings creates set using map of strings.
func OfStrings(s ...string) map[string]bool {
	out := map[string]bool{}
	for _, v := range s {
		out[v] = true
	}
	return out
}
