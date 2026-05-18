# Variadic Function in Go

Variadic function মানে এমন function যেটা **multiple number of arguments নিতে পারে** (0, 1, 2, 100... যেকোনো সংখ্যা)।

---

# Basic Syntax

```go id="v6k8ab"
func name(param ...Type)
```

👉 এখানে `...Type` মানে অনেকগুলো value নিতে পারবে।

---

# Simple Example

```go id="2xq8lp"
package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
}

func main() {
	sum(1, 2, 3)
	sum(10, 20)
	sum()
}
```

---

# Output

```text id="p8h1kd"
[1 2 3]
[10 20]
[]
```

👉 এখানে `nums` আসলে একটা slice (`[]int`)

---

# Real Example: Sum Function

```go id="n3v8qz"
package main

import "fmt"

func sum(nums ...int) {

	total := 0

	for _, v := range nums {
		total += v
	}

	fmt.Println("Total:", total)
}

func main() {
	sum(1, 2, 3)
	sum(5, 10)
	sum(100)
}
```

---

# Output

```text id="c7qk2x"
Total: 6
Total: 15
Total: 100
```

---

# Important Rule

👉 Variadic parameter must be **last parameter**

❌ Wrong:

```go id="x9m2qp"
func test(a int, nums ...int, b string) {}
```

---

# Correct:

```go id="m2q8ls"
func test(a int, b string, nums ...int) {}
```

---

# Passing Slice to Variadic Function

```go id="p8v3ld"
package main

import "fmt"

func print(nums ...int) {
	fmt.Println(nums)
}

func main() {

	arr := []int{1, 2, 3}

	print(arr...) // important
}
```

---

# Output

```text id="l7k2nv"
[1 2 3]
```

👉 `arr...` মানে slice expand করা

---

# Real Life Example

## Logging system

```go id="z3v9qp"
func log(message string, values ...any) {
	fmt.Println(message, values)
}
```

---

```go id="b8k2mv"
log("User data:", "Suvo", 25, true)
```

---

# Output

```text id="r9n4px"
User data: [Suvo 25 true]
```

---

# fmt.Println is Variadic

তুমি প্রতিদিন use করো 😄

```go id="t4k9xp"
fmt.Println("A", "B", "C")
```

👉 ভিতরে এটা:

```go id="u7m2dq"
func Println(a ...interface{})
```

---

# Summary

| Feature        | Meaning                    |
| -------------- | -------------------------- |
| `...T`         | multiple values accept করে |
| becomes        | slice internally           |
| last parameter | must be variadic           |
| `slice...`     | unpack slice               |

---

# Simple Line

> Variadic function মানে এক function যেটা unlimited arguments নিতে পারে এবং internally slice হিসেবে handle করে।
