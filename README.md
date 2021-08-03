# Slice 的研究

在 Golang 中，有 Array and Slice 這兩種 type。

Array 類似其他語言的定義，但是在 Golang 中 Array type 必須定義長度以及元素的類型

```go
var a [4]int
a[0] = 1
i := a[0] // i == 1
```

Array 的長度是固定的，因此，不同長度的 Array 是不一樣的類型

```go
var a [4]int
var b [5]int
reflect.TypeOf(a) == reflect.TypeOf(b) // false
```

Array 的定義其實不是這麼好使用，像是在呼叫一個 method 的時候，必須傳入符合其定義的 type ，否則 complier 是不會讓它通過的

```go
func sum(arr [10]int) int {}

sum([]int{1,2,3}) 
// cannot use []int{...} (type []int) as type [10]int in argument to sum
```

從以上的範例可以想像，如果我們要計算一個 array 的總和，但限制一定要傳入 10 個元素的陣列，這在使用上是非常不方便的。

因此 Golang 提供了另一個在使用上比較方便的類型：**Slice**

Slice 是一個 struct ，可以看到在 source code 中的定義：

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

array unsafe.Pointer 是一個指向 underlying array 內元素的指標，len 是 slice 的長度，cap 是從目前 underlying array index 到 array 底的長度，有點難以理解，所以用程式碼說明：

```go
a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
b := a[3:6] // [3,4,5] len: 3 cap: 7
c := a[3:] // [3,4,5,6,7,8,9] len: 7 cap: 7
```

b 跟 c 都指向同一個 array，b = [3,4,5] c = [3,4,5,6,7,8,9]，len 分別是 3 跟 7 應該沒有問題，而兩者的 cap 都是 7 ，這是因為兩個 slice 都指向 underlying array a 的 index 3，而 a 的長度是 10，所以 10 - 3 = 7。

---

再來談談一些 slice 應用上的可能會有的疑惑

```go
a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
b := a[3:6] // [3,4,5] len: 3 cap: 7
c := a[3:] // [3,4,5,6,7,8,9] len: 7 cap: 7

b[3] = 20 // panic: runtime error: index out of range [3] with length 3
```

雖然 b 的 cap 是 7 ，但是 len 只有 3 因此如果直接 assign [3] 會造成 panic error。那麼我們要怎麼擴充 b 呢？有三個方法： append, copy and re-slice

```go
// append
b = append(b, item)

//copy

b2 := make([]int, len(b)+1)
copy(b2. b)
b2[len(b)] = item

//re-slice
b = [:len(b)+1]
b[len(b)] = item
```

append 比較簡單也常用。copy 要注意到 underlying array 是不同的，在某些情況下，使用 copy 來 insert 會比 append 來得有效率 ([slice tricks](https://github.com/golang/go/wiki/SliceTricks) 中有提到，之後會再說明)。re-slice 也很直覺易懂，不過要注意到 capacity 是否足夠，否則是會產生 runtime error 的。

值得一提的是使用 append or re-slice 要注意到 b 的 underlying array 也會被修改這一點

```go
b = append(b, 20)

// b = [3,4,5,20]

// c = [3,4,5,20,7,8,9]

// a = [0,1,2,3,4,5,20.7,8,9]
```

在 underlying array 有足夠的 capacity 下，會做一次的 re-slice 並將新的元素放置進去，因此 underlying array 中的元素就會被置換，連帶影響到其他指向這個 array 的 slices。

但如果 underlying array 沒有足夠的 capacity 呢？ 請大家再看一段程式碼：

```go
c = append(c, 20)
// [3,4,5,6,7,8,9,20]
b = b[:cap(b)]
// [3,4,5,6,7,8,9]
```

想要知道其中運作的原理，要先從 [growslice](https://github.com/golang/go/blob/4bb0847b088eb3eb6122a18a87e1ca7756281dcc/src/runtime/slice.go#L162) 這一段 source code 來下手。

---

```go
// growslice handles slice growth during append.
// It is passed the slice element type, the old slice, and the desired new minimum capacity,
// and it returns a new slice with at least that capacity, with the old data
// copied into it.
```

當 slice 的 cap 不足需要擴大的時候，就會呼叫 growslice 這一個 func 處理，會先根據 old. cap (old.cap) 以及 exp.cap 來計算出 new.cap。處理的邏輯如下: (old.cap ⇒ 舊的 exp.cap ⇒ 期望的 new.cap ⇒ 最後計算出來新的)

1. 如果 exp.cap 大於兩倍的 old.cap，new.cap = exp.cap
2. 如果 exp.cap 小於兩倍的 old.cap，而且 old.cap < 1024，那麼 new.cap 就等於兩倍的 old.cap
3. 如果 exp.cap 小於兩倍的 old.cap，但是 old.cap ≥ 1024，那麼就會跑一個 for loop 讓 new.cap = 1.25 old.cap ，直到 new.cap ≥ exp.cap
4. 最後，如果 old.cap ≤ 0, new.cap = exp.cap

計算完新的 cap 之後，就會把目前 array 內的元素複製到新的 array 中

所以當 c 要擴充的時候，因為 underlying array 的 capacity 已經不夠了，因此重新產生了一組新的 underlying array 給他，這個時候 b, c 兩者的 underlying array 已經不同了。

## Call by Value

```go
func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[3:6]
	c := a[3:]
	double(b)
	fmt.Println(b) // [6 8 10]
	fmt.Println(c) // [6 8 10 6 7 8 9]
}
func double(x []int) {
	for i, v := range x {
		x[i] = v * 2
	}
}
```

在 Golang 的世界裡面都是 call by value ，slice 也不例外，但是為什麼上面的程式碼 double 卻會影響到 c 呢？這是因為在傳遞 b 給 double 的時候，的確是複製了一份 b 的值，但是 b 的 slice (struct) 只有 ptr, len 以及 cap ，並沒有真正的持有元素，而複製出來的 slice 也指向了同樣的 underlying array ，所以在 double 裡面修改了元素，就會影響到 c。 

不過要利用這個特性必須要注意到改變 slice 的長度時 (append, re-slice or copy) 都會讓新的 slice 的 underlying array 變成新的，因此可能就會發生在新的 slice 中修改，但是其他地方的 slice 因為兩者的 underlying array 不一樣了，造成修改是無效的，要怎麼避免這個情況發生呢？有一個作法是 Slice of Pointers，也就是只傳遞指標，而不是數據本身，如此一來即使 underlying array 的元素被複製了，也還是指向相同的數據，但是使用 slice of pointers 有什麼好處與壞處呢？下一段我們將會來討論其優缺點。

## Slice of pointers vs Slice of structs

### Reference

- https://blog.golang.org/slices-intro
- https://github.com/golang/go/wiki/SliceTricks
- [https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b](https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b)
- [https://medium.com/@opto_ej/there-are-other-nuances-one-should-consider-c798f12be15c](https://medium.com/@opto_ej/there-are-other-nuances-one-should-consider-c798f12be15c)
- [https://philpearl.github.io/post/bad_go_slice_of_pointers/](https://philpearl.github.io/post/bad_go_slice_of_pointers/)
-