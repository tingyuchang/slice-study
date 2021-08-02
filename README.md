
Array 定義了長度以及元素的類型
var a [4]int

Slice internals
A slice is a descriptor of an array segment.
It consists of a pointer to the array, the length of the segment, and
its capacity (the maximum length of the segment).


Reference:
- https://blog.golang.org/slices-intro
- https://github.com/golang/go/wiki/SliceTricks
