package usecase

// Filter keeps valid items in slice
func Filter(x []interface{}, keep func(interface{}) bool) []interface{} {
	n := 0
	for _, v := range x {
		if keep(v) {
			x[n] = v
			n++
		}
	}
	// For elements which must be garbage collected
	for i := n; i < len(x); i++ {
		x[i] = nil
	}
	return x[:n]
}

// Filter2 using append to keep valid items
func Filter2(x []interface{}, keep func(interface{}) bool) []interface{} {
	b := x[:0]

	for _, v := range x {
		if keep(v) {
			b = append(b, v)
		}
	}
	// For elements which must be garbage collected
	for i := len(b); i < len(x); i++ {
		x[i] = nil
	}

	return b
}
