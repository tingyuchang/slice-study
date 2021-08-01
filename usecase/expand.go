package usecase

// Expand using append to expand length space in index of slice
func Expand(index, length int, x []interface{}) []interface{} {
	return append(x[:index], append(make([]interface{}, length), x[index:]...)...)
}

// Extend to the last of slice
func Extend(length int, x []interface{}) []interface{} {
	return append(x, make([]interface{}, length)...)
}