এটা হলো Go-এর **Generics** example।

```go id="epzgw9"
func print[T any](value T) {
	fmt.Println(value)
}
```

এখানে function যেকোনো type নিতে পারে কিন্তু type-safe way তে।

---

# Full Example

```go id="mtag2q"
package main

import "fmt"

// Generic Function
func print[T any](value T) {
	fmt.Println(value)
}

func main() {

	print("Suvo")

	print(100)

	print(true)

	print(99.5)
}
```

---

# Output

```text id="ukkg9z"
Suvo
100
true
99.5
```

---

# Breakdown

## `[T any]`

এটা generic type parameter।

মানে:

```go id="m72tpg"
T
```

একটা temporary type variable।

---

## `any`

```go id="uzd1mg"
any
```

মানে:

```go id="d4g6ho"
interface{}
```

অর্থাৎ যেকোনো type allowed।

---

# Execution Flow

যখন:

```go id="r0qqi0"
print("Suvo")
```

তখন Go internally বানায়:

```go id="0f0mvr"
func print(value string)
```

---

আবার:

```go id="a09e42"
print(100)
```

তখন internally:

```go id="i9xh5s"
func print(value int)
```

---

# Real Memory Concept

```text id="a1m8u0"
print("suvo")
        |
        ↓
T = string


print(100)
        |
        ↓
T = int
```

প্রতিবার call অনুযায়ী type replace হয়।

---

# Without Generics

আগে করতে হতো:

```go id="0zjwqc"
func print(value interface{}) {
	fmt.Println(value)
}
```

কিন্তু এখানে type safety কম।

---

# Generic Advantage

## Type Safe

```go id="4fe8jp"
func get[T any](value T) T {
	return value
}
```

---

```go id="x8wz5f"
name := get("suvo")
```

এখানে compiler জানে:

```go id="r4nhjo"
name is string
```

---

# Another Real Example

## Generic Add Function

```go id="dlf6bj"
package main

import "fmt"

func show[T any](data T) {
	fmt.Println(data)
}

func main() {

	show("Go")

	show(500)

	show(false)
}
```

---

# Multiple Type Parameter

```go id="a8m4mw"
func userInfo[T any, U any](name T, age U) {
	fmt.Println(name, age)
}
```

---

```go id="zy43xt"
userInfo("Suvo", 25)
```

---

# Real Industry Use

Generics use হয়:

- reusable libraries
- database packages
- caching systems
- utility functions
- slice/map operations

---

# Example: Generic Slice Printer

```go id="r8c9cs"
package main

import "fmt"

func printSlice[T any](items []T) {

	for _, value := range items {
		fmt.Println(value)
	}
}

func main() {

	nums := []int{1, 2, 3}

	names := []string{"suvo", "datta"}

	printSlice(nums)

	printSlice(names)
}
```

---

# Summary

| Part      | Meaning              |
| --------- | -------------------- |
| `T`       | custom type variable |
| `any`     | any type allowed     |
| `[T any]` | generic declaration  |
| `value T` | parameter type is T  |

---

# Simple Line

> Generics মানে একবার code লিখে multiple type এর জন্য safely ব্যবহার করা।
