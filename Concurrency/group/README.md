Go concurrency মানে হলো একই সময়ে একাধিক কাজ efficiently handle করা।
এটা mainly goroutine + channel দিয়ে করা হয়।

## কেন Concurrency দরকার?

যখন application এ অনেক task একসাথে আসে:

- API request handle
- File upload
- Database query
- Email send
- Notification
- Background job
- Multiple user request

Sequentially করলে slow হয়ে যায়।
Concurrency use করলে একই সাথে কাজ চলতে পারে।

---

# Real Problem Without Concurrency

ধরো ৩টা API call করতে হবে।

প্রতিটা call নিতে ২ second।

## Without Concurrency

```go
package main

import (
	"fmt"
	"time"
)

func fetchData(name string) {
	fmt.Println(name, "started")

	time.Sleep(2 * time.Second)

	fmt.Println(name, "finished")
}

func main() {

	start := time.Now()

	fetchData("API 1")
	fetchData("API 2")
	fetchData("API 3")

	fmt.Println("Total Time:", time.Since(start))
}
```

## Output

```go
API 1 started
API 1 finished

API 2 started
API 2 finished

API 3 started
API 3 finished

Total Time: 6s
```

এখানে problem:

- সব কাজ একটার পর একটা হচ্ছে
- সময় বেশি লাগছে
- CPU idle থাকতে পারে

---

# With Concurrency (Goroutine)

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(name string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(name, "started")

	time.Sleep(2 * time.Second)

	fmt.Println(name, "finished")
}

func main() {

	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(3)

	go fetchData("API 1", &wg)
	go fetchData("API 2", &wg)
	go fetchData("API 3", &wg)

	wg.Wait()

	fmt.Println("Total Time:", time.Since(start))
}
```

## Output

```go
API 1 started
API 2 started
API 3 started

API 2 finished
API 1 finished
API 3 finished

Total Time: 2s
```

---

# এখানে কী হলো?

আগে:

```text
API1 -> API2 -> API3
```

এখন:

```text
API1
API2   -> same time run
API3
```

তাই ৬ second এর বদলে প্রায় ২ second লাগছে।

---

# Goroutine কী?

`go` keyword দিয়ে lightweight thread তৈরি হয়।

```go
go myFunction()
```

এটাকে goroutine বলে।

---

# Concurrency কোন Problem Solve করে?

## 1. Multiple Request Handle

Web server এ হাজার user request।

Without concurrency:

```text
1 user finish -> next user
```

Very slow.

Go concurrency:

```text
1000 request same time handle
```

এই কারণে Go backend খুব fast।

---

## 2. I/O Waiting Problem

যখন:

- DB query
- API call
- File read

হয়, CPU wait করে।

Concurrency অন্য কাজ চালিয়ে যেতে পারে।

---

## 3. Background Task

Example:

- Email send
- OTP send
- Notification

Main response block না করে background এ run করা যায়।

```go
go sendEmail()
```

---

# Real Backend Example

## Without Concurrency

```go
func handler() {

	saveUser()

	sendEmail()

	generateReport()

}
```

সব sequential।

---

## With Concurrency

```go
func handler() {

	go sendEmail()

	go generateReport()

	saveUser()
}
```

User fast response পাবে।

---

# sync.WaitGroup কেন লাগে?

Main function wait না করলে program শেষ হয়ে যাবে।

```go
var wg sync.WaitGroup
```

## Add

কয়টা goroutine আছে।

```go
wg.Add(3)
```

## Done

একটা কাজ শেষ।

```go
defer wg.Done()
```

## Wait

সব finish হওয়া পর্যন্ত wait।

```go
wg.Wait()
```

---

# Concurrency vs Parallelism

## Concurrency

একই সময়ে multiple task manage করা।

```text
Task switching
```

---

## Parallelism

একই সময়ে multiple CPU core এ সত্যি সত্যি run।

```text
True simultaneous execution
```

Go দুটোই support করে।

---

# Interview Questions

## 1. Goroutine কী?

Lightweight thread managed by Go runtime।

---

## 2. Goroutine আর Thread difference?

- Thread heavy
- Goroutine lightweight
- হাজার হাজার goroutine possible

---

## 3. WaitGroup কেন use হয়?

Multiple goroutine finish হওয়া wait করার জন্য।

---

## 4. Concurrency problem কী কী?

- Race condition
- Deadlock
- Data inconsistency

---

# Race Condition Example

```go
count++
```

Multiple goroutine একই variable change করলে problem হয়।

Solution:

- Mutex
- Channel
- Atomic package

---

# Industry Use Case

## E-commerce

এক user order করলে:

- Save order
- Payment process
- Email send
- SMS send
- Inventory update

সব concurrency দিয়ে handle করা হয়।

---

# Go কেন Concurrency তে famous?

- Built-in goroutine
- Lightweight
- Fast scheduler
- Simple syntax
- High performance server

এই কারণে Google Go তৈরি করেছিল high concurrent system এর জন্য।
