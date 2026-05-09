# Go Function Notes (Bangla)

## Function কী?

Function হলো এমন একটি কোড ব্লক যাকে বারবার ব্যবহার করা যায়।

---

# 1. Function without Input & Output

```go
package main

import "fmt"

func printSomething() {
	fmt.Println("Education must be free")
}

func main() {
	printSomething()
}
```

## ব্যাখ্যা

- `printSomething()` নামে একটি function বানানো হয়েছে।
- এটার:
  - কোনো input নেই
  - কোনো output নেই

- শুধু `"Education must be free"` প্রিন্ট করে।

## Output

```txt
Education must be free
```

---

# 2. Function with String Input

```go
package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Welcome to the Go course,", name)
}

func main() {
	sayHello("Habib")
}
```

---

# Function Parameter

```go
name string
```

এখানে:

- `name` = variable
- `string` = data type

মানে function একটি string input নিবে।

---

# Function Call

```go
sayHello("Habib")
```

এখানে:

- `"Habib"` function এর ভিতরে `name` variable এ জমা হবে।

মানে:

```go
name = "Habib"
```

---

# Output

```txt
Welcome to the Go course, Habib
```

---

# fmt.Println এ Multiple Value Print

```go
fmt.Println("Welcome", name)
```

এখানে:

- `"Welcome"` একটি string
- `name` আরেকটি string variable
- comma `,` দিয়ে multiple value print করা যায়

---

# Important Things

## 1. Function Input না দিলে Error হবে

যদি function input চায়:

```go
func sayHello(name string)
```

তাহলে call করার সময় অবশ্যই value দিতে হবে।

ভুল:

```go
sayHello()
```

সঠিক:

```go
sayHello("Habib")
```

---

# 2. String কী?

ডাবল কোটেশনের `" "` ভিতরের লেখা হলো string।

উদাহরণ:

```go
"Habib"
"Hello"
"Go Language"
```

---

# Summary

- Function reusable code block
- Function input নিতে পারে
- Function output দিতে পারে
- `fmt.Println()` দিয়ে print করা হয়
- Multiple value print করতে comma ব্যবহার হয়
- String সবসময় `" "` এর ভিতরে থাকে
