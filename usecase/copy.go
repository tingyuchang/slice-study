package usecase

// CopyByCopy using origin provided copy method
func CopyByCopy(x []interface{}) []interface{} {
	result := make([]interface{}, len(x))
	copy(result, x)
	return result
}

// CopyByAppend using create nil slice and append source data into it
func CopyByAppend(x []interface{}) [] interface{} {
	return append([]interface{}{nil}, x...)
}
// CopyByAppendReSlice using re-slice [i:j:k]
func CopyByAppendReSlice(x []interface{}) [] interface{} {
	return append(x[:0:0], x...)
}