package usecase


func Filter(x []interface{}, keep func(interface{})bool) []interface{} {
	n := 0
	for _, v := range x {
		if keep(v) {
			x[n] = v
			n++
		}
	}
	return x[:n]
}