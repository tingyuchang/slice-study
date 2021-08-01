package usecase

// Delete using append to delete index item
func Delete(index int, x []interface{}) interface{} {
	return append(x[:index], x[index+1:]...)
}

// DeleteByCopy using copy to delete index item
func DeleteByCopy(index int, x []interface{}) interface{} {
	copy(x[:index], x[index+1:])
	x[len(x)-1] = nil
	return x[:len(x)-1]
}