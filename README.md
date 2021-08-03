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

雖然 b 的 cap 是 7 ，但是 len 只有 3 因此如果直接 assign [3] 會造成 panic error。那麼我們要怎麼擴充 b 呢？有三個方法： append, copy and reslice

```go
// append
b = append(b, item)

//copy

b2 := make([]int, len(b)+1)
copy(b2. b)

//re-slice
b = [:len(b)+1]
```

append 比較簡單也常用，如果要在 slice 尾端增加元素，可以直接使用 append:

但是要注意到 b 的 underlyig array 也會同時被修改這一點

```go
b = append(b, 20)

// b = [3,4,5,20]

// c = [3,4,5,20,7,8,9]

// a = [0,1,2,3,4,5,20.7,8,9]
```

在講解為什麼之前，想請大家再看一段程式碼：

```go
c = append(c, 20)
b = b[:cap(b)]
```

請問 a, b, c 分別最後是多少呢？這邊先不公佈答

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

### Reference

- https://blog.golang.org/slices-intro
- https://github.com/golang/go/wiki/SliceTricks