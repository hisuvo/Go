Go-তে **empty interface** মানে এমন একটি interface যেটার ভিতরে কোনো method নেই।

```go
interface{}
```

অথবা Go 1.18+ এ:

```go
any
```

`any` আসলে `interface{}` এর alias।

---

# Empty Interface কী?

যেহেতু এতে কোনো method নেই, তাই Go-এর **সব type** এই interface implement করে।

কারণ Go-তে interface implement করতে method match লাগে।
আর empty interface-এ method-ই নাই 😄

তাই:

- int
- string
- struct
- slice
- function
- map

সব কিছু store করা যায়।

---

# Real Life Example

ধরো তুমি একটা **box** বানাইছো যেখানে যেকোনো জিনিস রাখা যাবে।

- বই
- মোবাইল
- টাকা
- কলম

সব রাখা যায়।

Empty interface ঠিক সেই universal box এর মতো।

---

# Basic Example

```go
package main

import "fmt"

func main() {

	var data interface{}

	data = "Suvo"
	fmt.Println(data)

	data = 100
	fmt.Println(data)

	data = true
	fmt.Println(data)
}
```

Output:

```go
Suvo
100
true
```

একই variable এ different type রাখা যাচ্ছে।

---

# any Keyword

Modern Go-তে সাধারণত `any` use করা হয়।

```go
var data any

data = "Hello"
data = 10
```

এটা internally:

```go
var data interface{}
```

---

# Function Example

## Without Empty Interface

```go
func printName(name string) {
	fmt.Println(name)
}
```

এখানে শুধু string যাবে।

---

## With Empty Interface

```go
func printAnything(value interface{}) {
	fmt.Println(value)
}
```

এখন:

```go
printAnything("Suvo")
printAnything(100)
printAnything(true)
```

সব কাজ করবে।

---

# Industry Use Case

## 1. JSON Handling

যখন API response আসে, তখন আগে জানা থাকে না data কী type হবে।

তখন use হয়:

```go
map[string]interface{}
```

Example:

```go
data := map[string]interface{}{
	"name": "Suvo",
	"age":  25,
}
```

---

## 2. fmt.Println()

তুমি প্রতিদিন use করো 😄

```go
fmt.Println("hello", 10, true)
```

এটা different type print করতে পারে কারণ ভিতরে empty interface use করা হয়।

---

# Type Assertion

Problem হলো:

Empty interface এর ভিতরের original type Go জানে না।

তাই type বের করতে হয়।

---

## Example

```go
var data interface{} = "Suvo"

name := data.(string)

fmt.Println(name)
```

---

# Wrong Assertion Panic

```go
var data interface{} = "Suvo"

num := data.(int)
```

এটা panic দিবে।

---

# Safe Type Assertion

```go
var data interface{} = "Suvo"

name, ok := data.(string)

fmt.Println(name, ok)
```

Output:

```go
Suvo true
```

---

# Type Switch

Multiple type check করতে:

```go
package main

import "fmt"

func checkType(value interface{}) {

	switch v := value.(type) {

	case string:
		fmt.Println("String:", v)

	case int:
		fmt.Println("Integer:", v)

	case bool:
		fmt.Println("Boolean:", v)

	default:
		fmt.Println("Unknown")
	}
}

func main() {
	checkType("Suvo")
	checkType(100)
	checkType(true)
}
```

---

# Memory Concept

যখন:

```go
var data interface{} = "Suvo"
```

তখন interface internally দুইটা জিনিস রাখে:

1. Type information
2. Actual value

Conceptually:

```text
interface
   |
   |---- Type: string
   |---- Value: "Suvo"
```

---

# Problem of Empty Interface

## 1. Type Safety কমে যায়

Compiler আগে check করতে পারে না।

```go
var data interface{}
```

এখানে যেকোনো type ঢুকতে পারে।

---

## 2. Runtime Error হতে পারে

Wrong type assertion panic দিবে।

---

## 3. Code বুঝা কঠিন হয়

Developer বুঝতে পারে না variable এ কী আছে।

---

# Best Practice

Empty interface কম use করা ভালো।

Use করো যখন:

- unknown data
- generic utility
- JSON parsing
- dynamic data

---

# Go Generics আসার আগে

আগে generic system ছিল না।

তখন সব generic কাজের জন্য:

```go
interface{}
```

use করা হতো।

---

# Modern Go (Generics)

এখন:

```go
func print[T any](value T) {
	fmt.Println(value)
}
```

এখানে:

```go
any
```

মানে empty interface।

---

# Empty Interface vs Interface

## Normal Interface

```go
type Animal interface {
	Sound()
}
```

এখানে method লাগবে।

---

## Empty Interface

```go
interface{}
```

কোনো method লাগে না।

সব type allowed।

---

# Summary

| Topic         | Meaning                 |
| ------------- | ----------------------- |
| `interface{}` | সব type hold করতে পারে  |
| `any`         | `interface{}` এর alias  |
| Use Case      | JSON, dynamic data      |
| Risk          | runtime panic           |
| Type Check    | assertion / type switch |

---

# Simple Line

> Empty interface হলো Go-এর universal container যেখানে যেকোনো type রাখা যায়।
