package usecase

import (
	"math/rand"
	"sort"
)

// Reverse is using reverse loop to switch items (but not efficient)
func Reverse(x []interface{}) []interface{} {
	s := make([]interface{}, len(x))
	for i := len(x) - 1; i >= 0; i-- {
		s[len(x)-1-i] = x[i]
	}
	return s
}

// Reverse2 switch item from left(start) and right(last)
func Reverse2(x []interface{}) []interface{} {
	left, right := 0, len(x)-1
	for left < right {
		x[left], x[right] = x[right], x[left]
		left++
		right--
	}
	return x
}

// Shuffling assign random new index for each item in slice
func Shuffling(x []interface{}) []interface{} {
	for i := len(x) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	return x
}

// Batch is useful when we must to handle huge amount slice
func Batch(x []interface{}, batchSize int) [][]interface{} {
	s := make([][]interface{}, (len(x)+batchSize-1)/batchSize)
	for i, v := range x {
		quotient := i / batchSize
		s[quotient] = append(s[quotient], v)
	}
	return s
}

// Batch2 is more efficient than Batch
func Batch2(x []interface{}, batchSize int) [][]interface{} {
	s := make([][]interface{}, 0, (len(x)+batchSize-1)/batchSize)
	for batchSize < len(x) {
		x, s = x[batchSize:], append(s, x[0:batchSize])
	}
	s = append(s, x)
	return s
}

// Dedupalicate returns unique items in slice
func Dedupalicate(x []int) []int {
	if len(x) == 0 {
		return nil
	}

	if len(x) == 1 {
		return x
	}

	sort.Ints(x)
	j := 0

	for i := 1; i < len(x); i++ {
		if x[i] == x[j] {
			continue
		}
		j++
		x[j], x[i] = x[i], x[j]
	}

	return x[:j+1]
}

// MoveToFront picks select item to first one
func MoveToFront(needle string, haystack []string) []string {
	if len(haystack) != 0 && haystack[0] == needle {
		return haystack
	}

	prev := needle

	for i, elem := range haystack {
		haystack[i] = prev
		if elem == needle {
			return haystack
		} else {
			prev = elem
		}
	}
	return append(haystack, prev)
}

func SlidingWindow(size int, input []int) [][]int {
	if len(input) < size {
		return [][]int{input}
	}
	r := make([][]int, 0, len(input)-size+1)

	for i, j := 0, size; j < len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}
	return r
}
