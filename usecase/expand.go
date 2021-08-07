package usecase

// Expand using append to expand length space in index of slice
func Expand(index, length int, x []interface{}) []interface{} {
	return append(x[:index], append(make([]interface{}, length), x[index:]...)...)
}

// Extend to the last of slice
func Extend(length int, x []interface{}) []interface{} {
	return append(x, make([]interface{}, length)...)
}

// Insert insert item into index of slice
func Insert(x []interface{}, index int, items ...interface{}) []interface{} {
	return append(x[:index], append(items, x[index:]...)...)
}

func InsertByCopy(x []interface{}, index int, item interface{}) []interface{} {
	s := append(x, 0)
	copy(s[index+1:], s[index:])
	s[index] = item
	return s
}

//
// Insert2 is verbose way only copies elements
// in a[i:] once and allocates at most once.
// But, as of Go toolchain 1.16, due to lacking of
// optimizations to avoid elements clearing in the
// "make" call, the verbose way is not always faster.
//
// Future compiler optimizations might implement
// both in the most efficient ways.
//
// Assume element type is int.
func Insert2(x []interface{}, index int, items ...interface{}) []interface{} {
	if n := len(x) + len(items); n <= cap(x) {
		s2 := x[:n]
		copy(s2[index+len(items):], x[index:])
		copy(s2[index:], items)
		return s2
	}
	s2 := make([]interface{}, len(x)+len(items))
	copy(s2, x[:index])
	copy(s2[index:], items)
	copy(s2[index+len(items):], x[index:])
	return s2
}
