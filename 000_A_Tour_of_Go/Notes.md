# Switch

A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

```go
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

## Switch with no condition

Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

# Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
kjjo
upackage main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

## Stacking defers

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover).

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

/*
counting
done
9
8
7
6
5
4
3
2
1
0
*/
```

# Pointers
Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is nil.

The * operator denotes the pointer's underlying value.
The & operator generates a pointer to its operand.
```go
var p *int
i := 42
p = &i
```
This is known as "dereferencing" or "indirecting".
Unlike C, Go has no pointer arithmetic.

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(p)  // address of i : 0xc000012070
	fmt.Println(*p) // read i through the pointer : 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i : 21

	p = &j         // point to j
	fmt.Println(*p)  // see the value of j through the pointer : 2701
	*p = *p / 31   // divide j through the pointer
	fmt.Println(j) // see the new value of j : 87

}
```

## Pointers to Structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

# Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

## Slices are like reference to arrays
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) //[John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)  //[John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a, b)  //[John XXX] [XXX George]
	fmt.Println(names) //[John XXX George Ringo]
}
```

## Slices length and Capacity
The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

`s[:1]` -> `s[(start index):1]`   
`s[1:]` -> `s[1:(last index)]`
```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	// s[:1] reduces the length to 1 but keeps the full capacity,
	// as the slice still points to the original start index (index 0).
	s = s[:1]
	printSlice(s) // len=1 cap=6 [2]

	// Extend the slice length back to 4.
	// Since the start index hasn't changed, the slice can extend forward.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop the first two values by slicing forward.
	// s = s[2:] shifts the start index to index 2, making earlier elements inaccessible.
	// The capacity is reduced as well, as it now reflects the space available from index 2 onward.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]

	// Drop another two values by slicing forward again.
	// s = s[2:] shifts the start index further, now to index 4 in the original array.
	// The slice is empty, and capacity is reduced to 2 (the remaining elements in the original array).
	s = s[2:]
	printSlice(s) // len=0 cap=2 []

	// At this point, the original elements like [2, 3] are inaccessible through this slice.
	// This is because slicing forward changes the slice's start index, making previous data "lost".
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

## Range
Basically a lazy for-loop
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	/*
		output :
		2**0 = 1
		2**1 = 2
		2**2 = 4
		2**3 = 8
	*/

	//If you only want the index, you can omit the second variable.
	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
	//output:1 2 4 8

	//You can skip the index or value by assigning to _.
	for i := range pow {
		fmt.Printf("%d ", i)
	}
	for i, _ := range pow {
		fmt.Printf("%d ", i)
	}
	//They are the same, output:0 1 2 3
}
```

# Maps
Key-Value Map, Underlying data structure is hashmap , just like C++ unordered_map
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	/*
	   The zero value of a map is nil. A nil map has no keys, nor can keys be added.

	   The make function returns a map of the given type, initialized and ready for use.

	*/
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Google"] = Vertex{
		37.42202, -122.08408,
	}
	/*
	   If remove the make line, Compilation Error : panic: assignment to entry in nil map

	   use Map Literals to remove the make line

	*/
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	/*
	   If the top-level type is just a type name, you can omit it from the elements of the literal.
	*/
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m["Bell Labs"])
}
```
## Map CRUD
```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
//Insert or update an element in map m: m[key] = elem
	m["Answer"] = 42
	//Retrieve an element: elem = m[key]
	fmt.Println("The value:", m["Answer"]) //The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) //The value: 48
//Delete an element: delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) //The value: 0

	/*
use := instead of = when either elem or ok have not yet been declared
	*/
	//Test that a key is present with a two-value assignment: elem, ok = m[key]
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) //The value: 0 Present? false
	
	//If key is not in the map, then elem is the zero value for the map's element type. 
}
```
# Functions
## Function Value
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values. 

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
    //Return values
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) //13
    //Funtion arguement
	fmt.Println(compute(hypot)) //5
	fmt.Println(compute(math.Pow)) //81
}
```
## Function Closure
```go
package main

import "fmt"

// Function that returns a closure
func createCounter() func() int {
	count := 0 // This variable is initialized once and captured by the closure

	// Return an anonymous function (closure) that increments and returns the count
	return func() int {
		count++
		// Increment the captured variable
		return count // Return the updated count
	}
}

func main() {
	// Create a new counter closure
	counter := createCounter()
	/*
		now the counter is a function with
		1. count++
		2. return count
	*/
	// Use the closure multiple times
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
}
```
