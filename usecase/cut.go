package usecase

// Cut using append to cut items
func Cut(from, count int, x []interface{}) []interface{} {
	return append(x[:from], x[from+count:]...)
}

// CutByCopy using copy to cut items
func CutByCopy(from, count int, x []interface{}) []interface{} {
	copy(x[from:], x[from+count:])
	for k, n := len(x)-count, len(x); k < n; k++ {
		x[k] = nil // or the zero value of T
	}
	return x[:len(x)-count]
}