package usecase

// Pop return latest item in slice
func Pop(x []interface{}) (interface{}, []interface{}) {
	return x[len(x)-1] , x[:len(x)-1]
}

// Push insert item at latest in slice
func Push(x []interface{}, item interface{}) []interface{} {
	return append(x, item)
}

// Unshift insert item at first in slice
func Unshift(x []interface{}, item interface{}) []interface{} {
	return append([]interface{}{item}, x...)
}

// Shift return first item in slice
func Shift(x []interface{}) (interface{}, []interface{}) {
	return x[0], x[1:]
}
