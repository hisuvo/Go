# Go `defer` কী?

Go language এ `defer` এমন একটি keyword যেটা কোনো function call-কে **বর্তমান function শেষ হওয়ার ঠিক আগে execute** করতে ব্যবহার হয়।

মানে:

- function এর মধ্যে `defer` লিখলে,
- ওই line সাথে সাথে execute হবে না,
- বরং function return হওয়ার আগে execute হবে।

---

# Basic Syntax

```go
defer functionName()
```

---

# Simple Example

```go
package main

import "fmt"

func main() {

	fmt.Println("Start")

	defer fmt.Println("Deferred Line")

	fmt.Println("End")
}
```

## Output

```txt
Start
End
Deferred Line
```

---

# Execution Flow

Go এ `defer` line execute না হয়ে stack এ জমা হয়।

Flow:

```txt
1. Start print
2. defer stack এ save
3. End print
4. function শেষ
5. defer execute
```

---

# Real Life Example

ধরো তুমি একটি file open করলে।

যদি file close করতে ভুলে যাও?

তাহলে:

- memory leak হতে পারে
- resource busy থাকতে পারে
- performance problem হতে পারে

এজন্য Go তে:

```go
defer file.Close()
```

ব্যবহার করা হয়।

---

# Real File Example (Industry Style)

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("File Open Error")
		return
	}

	defer file.Close()

	fmt.Println("File Successfully Opened")
}
```

---

# এখানে কী হচ্ছে?

## Step 1

```go
file, err := os.Open("test.txt")
```

File open হচ্ছে।

---

## Step 2

```go
defer file.Close()
```

main function শেষ হওয়ার আগে automatically file close হবে।

---

## Step 3

যদিও:

- error হয়
- return হয়
- panic হয়

তবুও defer run হবে।

এটাই defer-এর power।

---

# Database Real Example

```go
db, err := sql.Open("mysql", "root:1234@/test")

if err != nil {
	return
}

defer db.Close()
```

---

# API Server Real Example

```go
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Start")

	defer fmt.Println("Request End Log")

	fmt.Println("Processing Request")
}
```

---

# Multiple Defer

```go
package main

import "fmt"

func main() {

	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
}
```

## Output

```txt
Third
Second
First
```

---

# কেন এমন হয়?

কারণ defer works like:

# LIFO

```txt
Last In First Out
```

ঠিক stack data structure এর মতো।

---

# Memory Simulation

```go
defer A()
defer B()
defer C()
```

Stack:

```txt
TOP -> C
       B
       A
```

Execute হবে:

```txt
C
B
A
```

---

# Defer With Return

```go
package main

import "fmt"

func test() int {

	defer fmt.Println("Deferred")

	return 10
}

func main() {
	fmt.Println(test())
}
```

## Output

```txt
Deferred
10
```

---

# Important Concept

Return value আগে set হয়।

তারপর defer execute হয়।

তারপর function finally return করে।

---

# Named Return + Defer

```go
package main

import "fmt"

func test() (x int) {

	defer func() {
		x++
	}()

	return 10
}

func main() {
	fmt.Println(test())
}
```

## Output

```txt
11
```

---

# কেন 11?

Flow:

```txt
x = 10
defer run
x++
return x
```

---

# Panic Recovery Example

```go
package main

import "fmt"

func main() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}

	}()

	panic("Server Crashed")
}
```

---

# Output

```txt
Recovered: Server Crashed
```

---

# Industry Usage of defer

## 1. File Close

```go
defer file.Close()
```

---

## 2. Database Close

```go
defer db.Close()
```

---

## 3. Unlock Mutex

```go
mu.Lock()
defer mu.Unlock()
```

---

## 4. HTTP Body Close

```go
defer response.Body.Close()
```

---

## 5. Execution Time Measure

```go
start := time.Now()

defer func() {
	fmt.Println(time.Since(start))
}()
```

---

# Important Interview Questions

---

## 1. What is defer in Go?

`defer` delays execution of a function until the surrounding function returns.

---

## 2. Why use defer?

Resource cleanup safely handle করার জন্য।

Example:

- file close
- db close
- unlock mutex
- recover panic

---

## 3. Does defer execute immediately?

No.

Function return হওয়ার ঠিক আগে execute হয়।

---

## 4. How does multiple defer work?

LIFO order এ execute হয়।

```txt
Last In First Out
```

---

## 5. Does defer work after panic?

Yes.

panic হলেও defer execute হয়।

---

## 6. Can defer modify return value?

Yes.

Named return variable থাকলে modify করতে পারে।

---

## 7. Is defer expensive?

Slightly overhead আছে।

Loop এর ভিতরে unnecessary defer avoid করা ভালো।

---

# Bad Practice Example

```go
for i := 0; i < 100000; i++ {
	defer fmt.Println(i)
}
```

এতে huge memory overhead হতে পারে।

---

# Better Practice

Loop শেষে manual cleanup করো যদি huge iteration হয়।

---

# Senior Level Interview Question

## Question:

Why is defer heavily used in Go backend development?

## Answer:

Because Go backend applications frequently manage resources like:

- files
- database connections
- mutex locks
- network connections
- HTTP response bodies

`defer` ensures proper cleanup even during:

- errors
- early returns
- panic situations

This makes code safer, cleaner, and more maintainable.

---

# Short Definition

> `defer` হলো Go-এর cleanup mechanism যা function return হওয়ার আগে delayed function execute করে।
