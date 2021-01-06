# Answer
```
[5 5 5 5 5]
[1 2 3 4]
```
# Explaination
> A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).[src](https://blog.golang.org/slices-intro)

> The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself [src](https://golang.org/pkg/builtin/#append)

Since we append `125` to slice which don't have enough capacity -> new underlying array is allocated so modifying this slice do not affect source slice.

If we want to change slice - we need initially create slice with bigger capacity:
```go
package main

import (
	"fmt"
)

func mod(a []int) {

	a = append(a, 125)

	for i := range a {
		a[i] = 5
	}

	fmt.Println(a)
}

func main() {
	sl := make([]int, 0, 5) // create slice with enough capacity
	sl = append(sl, []int{1, 2, 3, 4}...) // fill slice with initial data
	mod(sl)
	fmt.Println(sl)
}

```
```
[5 5 5 5 5]
[5 5 5 5]
```
