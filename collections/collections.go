package collections

// Contains ... TODO do it in a generic way
func Contains(xs []string, x string) bool {
	for _, y := range xs {
		if x == y {
			return true
		}
	}
	return false
}
