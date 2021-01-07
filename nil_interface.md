# Answer
50 line will fail the compilation
```
invalid operation: InitType() == nil (mismatched types S and nil)
```

If we comment the 1st line:
```
true
false
false
false
false
true
```

# Explaination
> fmt.Println(InitType() == nil)

Failed compilation - we can't compare type and nil, because `type S` will never be nil(it will be struct with blank(default) fields). Only Pointer to the type can be nil.
i.e add some field to the struct and debug it with printf :) 
```
....
type S struct {
	t int
}
...
fmt.Printf("type %T, InitType is %+v\n", InitType(), InitType())
...
```
result:
```
type main.S, InitType is {t:0}
...
```

> fmt.Println(InitPointer() == nil)

true, because we are not initialized anything. So pointer == nil.

> fmt.Println(InitPointer() == nil) <br>
> fmt.Println(InitEfaceType() == nil) <br>
> fmt.Println(InitEfacePointer() == nil) <br>
> fmt.Println(InitIfaceType() == nil) <br>
> fmt.Println(InitIfacePointer() == nil) <br>

They are not nil, because of interface structure. It's consists of 2 tuple:

* Underlying Type
* Underlying Value

When we have something, even a nil value - interface will be non-nil.

>fmt.Println(InitIfacePointer2() == nil)

true, as was mentioned before - we have nothing in the interface - NO Type, NO Value. So it's truly nil.
