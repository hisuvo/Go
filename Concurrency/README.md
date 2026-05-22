# Go Language Concurrency — Beginner → Advanced

## Concurrency কী?

Concurrency মানে একসাথে অনেক কাজ handle করা।

যেমন:

- Food delivery app
- Chat application
- Video streaming
- Banking system
- Ride sharing app

এগুলোতে হাজার হাজার request একসাথে আসে।
Go language concurrency খুব efficient ভাবে handle করতে পারে।

Go এর concurrency model mainly based on:

- Goroutine
- Channel
- Select
- WaitGroup
- Mutex
- Context
- Worker Pool
- Fan In / Fan Out
- Pipeline
- Atomic Operation

---

# 1. Goroutine

## Goroutine কী?

Lightweight thread।

Java/Python thread এর তুলনায় অনেক কম memory লাগে।

একটা normal function কে concurrent বানাতে `go` keyword ব্যবহার করা হয়।

---

## Example

```go
package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello Goroutine")
}

func main() {
	go printHello()

	time.Sleep(time.Second)
}
```

---

## Real Life Example

ধরো:

- Facebook user upload করছে image
- একই সাথে notification যাচ্ছে
- image resize হচ্ছে
- database save হচ্ছে

সব কাজ parallel/concurrent ভাবে করতে goroutine ব্যবহার হয়।

---

# 2. Multiple Goroutine

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Println("Worker", id, "started")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker", id, "finished")
}

func main() {
	for i := 1; i <= 5; i++ {
		go worker(i)
	}

	time.Sleep(3 * time.Second)
}
```

---

# Problem: Main Function আগে শেষ হয়ে যায়

এটা production এ dangerous।

Solution:

- WaitGroup

---

# 3. WaitGroup

## কেন লাগে?

সব goroutine finish হওয়া পর্যন্ত wait করতে।

---

## Example

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All Done")
}
```

---

# Real Life Example

E-commerce order system:

- inventory update
- payment verify
- email send
- invoice generate

সব finish না হওয়া পর্যন্ত response complete হবে না।

---

# 4. Channel

## Channel কী?

Goroutine এর মধ্যে data communication system।

Go philosophy:

> “Do not communicate by sharing memory; instead, share memory by communicating.”

---

## Channel Create

```go
ch := make(chan int)
```

---

## Send & Receive

```go
ch <- 10
value := <- ch
```

---

## Example

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello"
	}()

	msg := <-ch

	fmt.Println(msg)
}
```

---

# Real Life Example

Food delivery:

- kitchen prepares food
- delivery boy receives order

Kitchen → Channel → Delivery

---

# 5. Buffered Channel

Normal channel blocking করে।

Buffered channel কিছু data temporarily store করতে পারে।

---

## Example

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
}
```

---

# Real Life Example

Message Queue:

- RabbitMQ
- Kafka

Temporary data hold করে।

---

# 6. Channel Direction

## Receive Only

```go
func receive(ch <-chan int)
```

## Send Only

```go
func send(ch chan<- int)
```

---

## কেন important?

Production code safe হয়।

---

# 7. Select Statement

Multiple channel handle করতে।

---

## Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "API Response"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Database Response"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)

	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
```

---

# Real Life Example

Backend service:

- API response
- Cache response
- Database response

যেটা আগে আসবে সেটা return করবে।

---

# 8. Timeout Pattern

Production app এ খুব important।

---

## Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- "Data"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)

	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}
```

---

# Real Life Example

Payment gateway:

- 2 second এর মধ্যে response না এলে timeout।

---

# 9. Worker Pool

## Most Important Interview Topic

---

## Concept

Fixed number worker থাকবে।

Jobs queue থেকে কাজ নিবে।

---

## Example

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 10)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)

		go worker(w, jobs, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()
}
```

---

# Real Life Example

Image Processing Server

10000 image upload হলে:

- 3 worker image resize করছে
- memory overload হচ্ছে না

---

# 10. Mutex

## Problem: Race Condition

---

## Example Problem

```go
counter++
```

একই সময়ে multiple goroutine modify করলে data corrupted হতে পারে।

---

# Mutex Solution

```go
package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println(counter)
}
```

---

# Real Life Example

Bank balance update।

একই account এ দুইজন একই সময়ে টাকা তুললে race condition হতে পারে।

---

# 11. RWMutex

Read বেশি হলে use করা হয়।

- Multiple reader allowed
- One writer only

---

## Real Life Example

Blog website:

- লাখ লাখ user read করছে
- খুব কম write হচ্ছে

---

# 12. Atomic Operation

Mutex থেকেও faster।

---

## Example

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt64(&counter, 1)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go increment(&wg)
	}

	wg.Wait()

	fmt.Println(counter)
}
```

---

# Real Life Example

Page view counter।

---

# 13. Context Package

## Production Level Important Topic

Request cancel করতে।

---

## Example

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

---

# Timeout Example

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	2*time.Second,
)

defer cancel()
```

---

# Real Life Example

User browser close করলে:

- DB query stop
- API call stop

resource save হয়।

---

# 14. Fan Out / Fan In

## Fan Out

একটা job multiple worker এ distribute।

## Fan In

সব result এক channel এ collect।

---

# Real Life Example

Search engine:

- Google multiple server এ search পাঠায়
- result collect করে

---

# 15. Pipeline Pattern

এক stage এর output আরেক stage এ যায়।

---

## Example

```go
func generate(nums ...int) <-chan int
func square(in <-chan int) <-chan int
```

---

# Real Life Example

Video processing pipeline:

- upload
- compress
- watermark
- save CDN

---

# 16. Deadlock

## Deadlock কী?

সব goroutine wait করছে।

কেউ progress করতে পারছে না।

---

## Example

```go
ch := make(chan int)

<-ch
```

---

# Error

```bash
fatal error: all goroutines are asleep - deadlock!
```

---

# 17. Race Condition Detect

Go built-in tool:

```bash
go run -race main.go
```

---

# 18. Scheduler

Go runtime scheduler:

- Goroutine manage করে
- OS thread এর সাথে mapping করে

Model:

- G → Goroutine
- M → Machine(Thread)
- P → Processor

---

# 19. CSP Model

Go inspired by:

Tony Hoare

Communicating Sequential Processes।

---

# 20. Goroutine Leak

## Dangerous Production Issue

Channel wait এ forever আটকে থাকা।

---

## Bad Example

```go
func worker(ch chan int) {
	<-ch
}
```

---

# Solution

- context
- timeout
- close channel

---

# Production Architecture Example

## Ride Sharing App

### Concurrent Tasks

1. Driver location tracking
2. Payment processing
3. Notification sending
4. Route calculation
5. Chat system
6. Trip history save

সব goroutine based architecture এ হতে পারে।

---

# Advanced Interview Questions

## Beginner

### 1. Goroutine কী?

Lightweight concurrent function।

---

### 2. Channel কেন use করা হয়?

Goroutine communication এর জন্য।

---

### 3. Buffered vs Unbuffered channel difference?

Buffered temporary data hold করতে পারে।

---

### 4. WaitGroup কী?

Multiple goroutine finish হওয়া পর্যন্ত wait করে।

---

# Intermediate

### 5. Mutex vs Channel?

| Mutex                 | Channel              |
| --------------------- | -------------------- |
| Shared memory protect | Communication        |
| Faster sometimes      | Cleaner architecture |

---

### 6. Deadlock কী?

সব goroutine blocking state এ থাকা।

---

### 7. Race Condition কী?

Multiple goroutine unsafe shared data access।

---

### 8. Select statement কী?

Multiple channel operation handle।

---

# Advanced Interview Questions

### 9. Go scheduler explain করুন।

G-M-P model।

---

### 10. Context package কেন important?

Cancellation & timeout management।

---

### 11. Worker Pool কেন use করা হয়?

Memory control + scalable processing।

---

### 12. Goroutine leak কী?

Unused blocked goroutine memory consume করা।

---

### 13. Fan In / Fan Out explain করুন।

Task distribution and aggregation pattern।

---

### 14. Mutex vs RWMutex?

RWMutex multiple reader allow করে।

---

### 15. Atomic vs Mutex?

Atomic lightweight and faster for simple operation।

---

# Senior Level System Design Question

## Design YouTube Video Processing System Using Go Concurrency

Expected Answer:

- Worker pool
- Queue
- Pipeline
- Context cancellation
- Retry mechanism
- Fan out processing
- CDN upload concurrent handling

---

# Best Practices

## DO

✅ Use context
✅ Close channels properly
✅ Use worker pool
✅ Use race detector
✅ Limit goroutine creation

---

## DON'T

❌ Infinite goroutine
❌ Shared memory without lock
❌ Forgetting cancel()
❌ Blocking channel forever

---

# Most Important Packages

| Package | Usage            |
| ------- | ---------------- |
| sync    | WaitGroup, Mutex |
| context | cancellation     |
| atomic  | atomic operation |
| time    | timeout          |
| runtime | goroutine info   |

---

# Real Interview Coding Task

## Task

Create:

- 5 worker
- process 100 jobs
- timeout support
- graceful shutdown

এটা খুব common interview problem।

---

# Learning Roadmap

## Step 1

Learn:

- goroutine
- channel
- select

## Step 2

Learn:

- mutex
- waitgroup
- worker pool

## Step 3

Learn:

- context
- pipeline
- fan in/out

## Step 4

Learn:

- scheduler
- atomic
- distributed systems

---

# Best Resources

- [Go Official Website](https://go.dev?utm_source=chatgpt.com)
- [Go By Example](https://gobyexample.com?utm_source=chatgpt.com)
- [Effective Go](https://go.dev/doc/effective_go?utm_source=chatgpt.com)
- [Go Concurrency Patterns Talk by Google](https://go.dev/talks/2012/concurrency.slide?utm_source=chatgpt.com)

---

# Final Summary

Go concurrency powerful কারণ:

- Lightweight goroutine
- Fast scheduler
- Easy communication via channel
- Production-ready concurrency primitives

এই কারণেই:

- Docker
- Kubernetes
- Terraform
- Prometheus

এর মতো systems Go দিয়ে build করা হয়েছে।

Examples:

- Docker
- Kubernetes
- Terraform
- Prometheus
