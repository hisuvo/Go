# Race Condition in Go (Detailed Bangla Explanation)

Go language এ concurrency খুব powerful feature।
কিন্তু multiple goroutine যদি একই data একসাথে read/write করে তাহলে **Race Condition** হয়।

---

# Simple Example

```go
package main

import (
	"fmt"
)

var count int = 0

func increment() {
	for i := 0; i < 1000; i++ {
		count++
	}
}

func main() {

	go increment()
	go increment()

	fmt.Scanln()

	fmt.Println("Final Count:", count)
}
```

---

# Expected Output কত হওয়া উচিত?

2 goroutine × 1000 increment

Expected:

```txt
2000
```

কিন্তু actual output হতে পারে:

```txt
1543
1788
1932
```

প্রতি run এ different result আসতে পারে।

এটাই Race Condition।

---

# কেন হয়?

কারণ:

```go
count++
```

এটা single operation না।

এটা internally 3 step এ কাজ করে:

1. Read current value
2. Add 1
3. Write back

মানে:

```txt
Read count = 5
Add 1 => 6
Write 6
```

---

# Problem Visualization

ধরি:

```txt
count = 5
```

এখন 2 goroutine একই সময়ে execute করলো।

---

## Goroutine A

```txt
Read count => 5
```

---

## Goroutine B

```txt
Read count => 5
```

---

## Goroutine A

```txt
Add 1 => 6
Write => 6
```

---

## Goroutine B

```txt
Add 1 => 6
Write => 6
```

---

# Final Result

Expected ছিল:

```txt
7
```

কিন্তু হলো:

```txt
6
```

একটা increment lost হয়ে গেছে।

এটাই race condition।

---

না, শুধু `count++` এ না।
যেকোনো shared variable multiple goroutine একসাথে access করলে race condition হতে পারে।

Example:

```go
name = "Suvo"
```

```go
users = append(users, user)
```

```go
mapData["id"] = 10
```

```go
balance = balance - 500
```

এগুলাও race condition create করতে পারে যদি multiple goroutine same data modify/read করে।

`count++` শুধু সবচেয়ে common example কারণ এটা internally:

```txt
Read → Modify → Write
```

এই 3 step এ কাজ করে।

---

# Race Condition Detect করার উপায়

Go built-in race detector দেয়।

Run:

```bash
go run -race main.go
```

বা:

```bash
go test -race
```

---

# Output Example

```txt
WARNING: DATA RACE
Read at 0x...
Write at 0x...
```

---

# Solution 1: Mutex

Mutex মানে:

> একসময় শুধু একজন access করতে পারবে।

---

## Example

```go
package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {

		mu.Lock()

		count++

		mu.Unlock()
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	fmt.Println("Final Count:", count)
}
```

---

# Mutex কীভাবে কাজ করে?

```txt
Lock()
```

দেওয়ার পর অন্য goroutine wait করবে।

যখন:

```txt
Unlock()
```

হবে তখন next goroutine ঢুকতে পারবে।

---

# Mutex এর সুবিধা

- Simple
- Shared memory safe করে
- Easy to understand

---

# Mutex এর Problem

বেশি lock ব্যবহার করলে:

- Performance কমে
- Deadlock হতে পারে
- Goroutine wait করে

---

# Deadlock Example

```go
mu.Lock()
mu.Lock()
```

একই goroutine দুইবার lock নিলে unlock না করলে program আটকে যাবে।

---

# Solution 2: Channel

Go এর philosophy:

> "Do not communicate by sharing memory; instead, share memory by communicating."

মানে shared variable না use করে data channel দিয়ে pass করো।

---

# Example

```go
package main

import (
	"fmt"
)

func worker(ch chan int) {

	for i := 0; i < 1000; i++ {
		ch <- 1
	}
}

func main() {

	ch := make(chan int)

	go worker(ch)
	go worker(ch)

	total := 0

	for i := 0; i < 2000; i++ {
		total += <-ch
	}

	fmt.Println(total)
}
```

---

# এখানে কী হলো?

worker shared variable touch করেনি।

বরং:

```txt
channel এ data send করেছে
```

তাই race condition হয়নি।

---

# Channel এর সুবিধা

- Safe communication
- Go style concurrency
- Cleaner architecture

---

# Channel এর Problem

- Complex হতে পারে
- Large system এ debugging hard
- Slow হতে পারে কিছু ক্ষেত্রে

---

# Solution 3: Atomic Package

যখন শুধু simple counter update দরকার হয় তখন atomic fastest solution।

Go package:

```go
sync/atomic
```

---

# Example

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {

		atomic.AddInt64(&count, 1)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	fmt.Println(count)
}
```

---

# Atomic কী করে?

CPU level এ safe operation করে।

এটা internally lock ছাড়া thread-safe update দেয়।

---

# Atomic এর সুবিধা

- Very fast
- Lightweight
- Counter এর জন্য best

---

# Atomic এর limitation

Complex logic handle করতে পারে না।

এটা mainly:

- counter
- flag
- simple state

এর জন্য ভালো।

---

# Real Backend Example

ধরি website এ visitor count:

```txt
10,000 request/sec
```

সব request যদি:

```go
visitorCount++
```

করে তাহলে race condition হবে।

তখন:

```go
atomic.AddInt64()
```

ব্যবহার করা হয়।

---

# Mutex vs Channel vs Atomic

| Feature          | Mutex  | Channel | Atomic |
| ---------------- | ------ | ------- | ------ |
| Shared Data Safe | Yes    | Yes     | Yes    |
| Speed            | Medium | Slow    | Fast   |
| Complex Logic    | Best   | Good    | Bad    |
| Simple Counter   | Okay   | Bad     | Best   |
| Go Style         | Medium | Best    | Low    |

---

# কখন কোনটা ব্যবহার করবে?

## Use Mutex

যখন:

- shared struct modify করো
- multiple field update
- critical section আছে

---

## Use Channel

যখন:

- goroutine communication
- worker pool
- pipeline
- message passing

---

## Use Atomic

যখন:

- counter
- flag
- metrics
- request count

---

# Interview Questions

## Q1: Race Condition কী?

Multiple goroutine একই data একই সময়ে modify করলে unpredictable result হওয়াকে race condition বলে।

---

## Q2: `count++` কেন unsafe?

কারণ এটা atomic operation না।
এটা read-modify-write process।

---

## Q3: Race detect কীভাবে করো?

```bash
go run -race main.go
```

---

## Q4: Mutex vs Atomic difference?

Mutex lock ব্যবহার করে।
Atomic CPU level thread-safe operation।

---

## Q5: Channel কি race condition prevent করে?

হ্যাঁ, কারণ shared memory না use করে communication করা হয়।

---

# Important Concept

Concurrency ≠ Parallelism

Concurrency মানে:

```txt
many tasks managed together
```

Parallelism মানে:

```txt
many tasks running literally same time
```

Race condition mainly shared memory problem।
