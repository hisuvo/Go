# Buffered Channel in Go (Golang)

Go language এ channel দুই ধরনের হয়:

1. **Unbuffered Channel**
2. **Buffered Channel**

আজ আমরা **Buffered Channel** deep dive করবো।

---

# Buffered Channel কী?

Buffered channel এমন channel যেটা কিছু data temporary memory তে store করতে পারে receiver ready হওয়ার আগেই।

মানে:

- sender immediately block হবে না
- channel এর ভিতরে buffer/full না হওয়া পর্যন্ত data জমা হবে

---

# Channel Create করার Syntax

```go
ch := make(chan int, 3)
```

এখানে:

- `chan int` → int type data যাবে
- `3` → buffer size

মানে channel সর্বোচ্চ ৩টা value store করতে পারবে।

---

# Real Life Analogy

ধরো:

- Unbuffered channel = ফোন call
  - দুইজন একসাথে ready থাকতে হবে

- Buffered channel = WhatsApp message
  - message send করে রাখতে পারো
  - receiver পরে পড়তে পারবে

---

# How Buffered Channel Work

## Example 1

```go
package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("All data sent")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

---

# Memory Visualization

```text
Buffer Size = 3

+----+----+----+
| 10 | 20 | 30 |
+----+----+----+
```

এরপর receiver data বের করে।

---

# Important Point

এই code এ sender block হয়নি কারণ:

- buffer size = 3
- আমরা 3টা data পাঠিয়েছি

---

# Buffer Full হলে কী হয়?

## Example

```go
package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	fmt.Println("Buffer Full")

	ch <- 30 // DEADLOCK
}
```

---

# Why Deadlock?

কারণ:

```text
Buffer Size = 2
```

Already:

```text
10
20
```

stored আছে।

এখন:

```go
ch <- 30
```

send করতে গেলে জায়গা নেই।

তাই sender wait করবে।

কিন্তু receiver নাই।

তাই:

```text
fatal error: all goroutines are asleep - deadlock!
```

---

# Buffered Channel Flow

```text
Sender ---> Buffer ---> Receiver
```

Unbuffered এ:

```text
Sender ---> Receiver
```

---

# len() and cap()

Buffered channel এ খুব useful।

---

## cap()

buffer capacity বলে।

```go
fmt.Println(cap(ch))
```

---

## len()

বর্তমানে কয়টা data আছে।

```go
fmt.Println(len(ch))
```

---

# Example

```go
package main

import "fmt"

func main() {

	ch := make(chan int, 5)

	ch <- 10
	ch <- 20

	fmt.Println("Length:", len(ch))
	fmt.Println("Capacity:", cap(ch))
}
```

---

# Output

```text
Length: 2
Capacity: 5
```

---

# Goroutine Example

```go
package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {

	ch <- "Job 1"
	ch <- "Job 2"

	fmt.Println("Worker Finished")
}

func main() {

	ch := make(chan string, 2)

	go worker(ch)

	time.Sleep(time.Second)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

---

# Why Buffered Channel Useful?

## 1. Reduce Blocking

sender বারবার wait করে না।

---

## 2. Producer-Consumer Pattern

```text
Producer -> Buffer -> Consumer
```

---

## 3. Task Queue

jobs queue হিসেবে use হয়।

---

## 4. Goroutine Communication Faster

temporary async communication দেয়।

---

# Buffered vs Unbuffered

| Feature                | Buffered        | Unbuffered |
| ---------------------- | --------------- | ---------- |
| Storage আছে            | ✅              | ❌         |
| Sender wait করে        | Buffer full হলে | Always     |
| Async communication    | ✅              | ❌         |
| Synchronization strong | কম              | বেশি       |
| Speed                  | অনেক সময় fast   | slower     |

---

# Important Interview Concept

# Buffered Channel does NOT mean fully asynchronous

অনেকে ভুল ভাবে buffered channel মানে completely async।

না।

Buffer full হলে sender আবার block হবে।

---

# Example

```go
ch := make(chan int, 1)

ch <- 10 // OK

ch <- 20 // BLOCK
```

---

# close() with Buffered Channel

```go
package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20

	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
```

---

# Important

channel close করার পরও buffered data read করা যায়।

---

# Internal Concept

Buffered channel internally:

- queue maintain করে
- FIFO order follow করে

```text
First In First Out
```

---

# FIFO Example

```go
ch <- 1
ch <- 2
ch <- 3
```

receive:

```text
1
2
3
```

---

# Common Mistakes

# 1. Over Buffering

```go
make(chan int, 1000000)
```

Huge memory waste হতে পারে।

---

# 2. Thinking Buffer Fixes Concurrency

না।

Race condition still হতে পারে।

---

# 3. Forgetting Receiver

buffer eventually full হবে।

---

# Real Life Use Case

# Worker Pool

```text
Jobs -> Buffered Channel -> Workers
```

---

# Example

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int) {

	for job := range jobs {
		fmt.Println("Worker", id, "processing", job)
		time.Sleep(time.Second)
	}
}

func main() {

	jobs := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	time.Sleep(6 * time.Second)
}
```

---

# Interview Questions

## Basic Questions

### 1. Buffered channel কী?

Receiver ready না থাকলেও limited data store করতে পারে এমন channel।

---

### 2. Buffer full হলে কী হয়?

Sender block করে।

---

### 3. Buffered channel FIFO follow করে?

হ্যাঁ।

---

### 4. len(ch) কী return করে?

বর্তমানে queued item count।

---

### 5. cap(ch) কী return করে?

maximum buffer size।

---

# Intermediate Questions

## 6. Buffered vs Unbuffered difference?

Buffered temporary storage দেয়।
Unbuffered direct synchronization দেয়।

---

## 7. Buffered channel কি race condition prevent করে?

না।

---

## 8. close() করার পর কি buffered data পাওয়া যায়?

হ্যাঁ।

---

## 9. Buffered channel কখন use করবেন?

- worker pool
- job queue
- producer-consumer
- rate limiting

---

# Advanced Interview Questions

## 10. Buffered channel কি internally mutex use করে?

Go runtime internally synchronization mechanism use করে safe concurrent access এর জন্য।

---

## 11. Buffered channel কি memory consume করে?

হ্যাঁ।
Buffer size অনুযায়ী memory allocate হয়।

---

## 12. Large buffer সবসময় better?

না।

কারণ:

- memory waste
- hidden bugs
- latency issue

---

# Best Practice

✅ Small meaningful buffer use করো

✅ Worker pool এ useful

✅ Backpressure বুঝে use করো

✅ Always receiver maintain করো

---

# Important Interview Line

> “Buffered channels improve throughput but reduce synchronization guarantees.”

---

# Most Important Concept

Buffered channel mainly use হয়:

```text
Decoupling sender and receiver speed
```

মানে sender আর receiver একই speed এ না হলেও communication চলবে।
