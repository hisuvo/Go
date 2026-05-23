# Unbuffered Channel in Go (Golang)

Unbuffered channel হলো Go concurrency এর সবচেয়ে important concept।

এটা goroutine এর মধ্যে:

- communication
- synchronization
- signaling

এর জন্য use হয়।

---

# Unbuffered Channel কী?

Unbuffered channel এ:

```text
Sender and Receiver must be ready at the same time
```

মানে:

- sender value send করবে
- receiver same time এ receive করবে

না হলে block হবে।

---

# Syntax

```go
ch := make(chan int)
```

এখানে:

- কোনো buffer size নাই
- capacity = 0

---

# Visualization

```text
Sender -----> Receiver
```

মাঝে কোনো storage নাই।

---

# Real Life Example

ধরো phone call:

- তুমি কথা বললে
- অন্যজন একই সময় শুনতে হবে

store করে পরে শোনার option নাই।

এটাই unbuffered channel।

---

# First Example

```go
package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	value := <-ch

	fmt.Println(value)
}
```

---

# Step By Step Flow

## Step 1

```go
ch := make(chan int)
```

unbuffered channel তৈরি হলো।

---

## Step 2

```go
go func() {
	ch <- 10
}()
```

goroutine send করার চেষ্টা করছে।

---

## Step 3

```go
value := <-ch
```

receiver receive করছে।

---

# Important

Send এবং receive একই সময় match হওয়ায় communication complete হয়।

---

# Synchronization Concept

Unbuffered channel শুধু data pass না।

এটা synchronization ও করে।

---

# Example

```go
package main

import (
	"fmt"
)

func worker(done chan bool) {

	fmt.Println("Working...")

	done <- true
}

func main() {

	done := make(chan bool)

	go worker(done)

	<-done

	fmt.Println("Finished")
}
```

---

# Why Useful?

```go
<-done
```

main goroutine wait করছে।

মানে worker finish না হওয়া পর্যন্ত main এগোবে না।

---

# Deadlock Example

```go
package main

func main() {

	ch := make(chan int)

	ch <- 10
}
```

---

# What Happens?

Sender আছে:

```go
ch <- 10
```

কিন্তু receiver নাই।

তাই:

```text
fatal error: all goroutines are asleep - deadlock!
```

---

# Another Deadlock

```go
package main

func main() {

	ch := make(chan int)

	<-ch
}
```

---

# Why?

Receiver wait করছে।

কিন্তু sender নাই।

---

# Important Rule

# Unbuffered channel এ:

## Send blocks until receive happens

and

## Receive blocks until send happens

---

# Internal Working

Unbuffered channel internally:

- data copy করে directly
- sender ↔ receiver handshake হয়

---

# Memory Visualization

Buffered channel:

```text
Sender -> BUFFER -> Receiver
```

Unbuffered:

```text
Sender <-> Receiver
```

Direct exchange।

---

# Capacity

```go
ch := make(chan int)

fmt.Println(cap(ch))
```

Output:

```text
0
```

কারণ buffer নাই।

---

# Goroutine Coordination Example

```go
package main

import (
	"fmt"
	"time"
)

func task(ch chan string) {

	time.Sleep(2 * time.Second)

	ch <- "Task Done"
}

func main() {

	ch := make(chan string)

	go task(ch)

	fmt.Println("Waiting...")

	msg := <-ch

	fmt.Println(msg)
}
```

---

# Output

```text
Waiting...
Task Done
```

---

# Why Unbuffered Channel Powerful?

## 1. Strong Synchronization

goroutine coordination এ best।

---

## 2. Prevent Timing Issues

exact communication point create করে।

---

## 3. Safe Goroutine Signaling

```text
done
stop
ready
finished
```

এসব signal এ useful।

---

# Buffered vs Unbuffered

| Feature         | Unbuffered                  | Buffered        |
| --------------- | --------------------------- | --------------- |
| Buffer আছে      | ❌                          | ✅              |
| Capacity        | 0                           | >0              |
| Sender blocks   | Always until receiver ready | Buffer full হলে |
| Synchronization | Strong                      | Less            |
| Async           | ❌                          | Partial         |
| Memory usage    | Low                         | More            |

---

# Important Concept

# Unbuffered channel = synchronization first

# Buffered channel = throughput first

---

# Real Use Cases

# 1. Worker Completion Signal

```go
done <- true
```

---

# 2. Goroutine Coordination

---

# 3. Pipeline Communication

---

# 4. Request-Response Flow

---

# Example: Ping Pong

```go
package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		ch <- "Ping"
	}()

	fmt.Println(<-ch)
}
```

---

# Channel Direction

## Send Only

```go
func send(ch chan<- int)
```

---

## Receive Only

```go
func receive(ch <-chan int)
```

---

# Example

```go
package main

import "fmt"

func send(ch chan<- int) {
	ch <- 100
}

func receive(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {

	ch := make(chan int)

	go send(ch)

	receive(ch)
}
```

---

# Common Mistakes

# 1. No Goroutine

```go
ch <- 10
```

receiver না থাকলে deadlock।

---

# 2. Forgetting Receiver

send forever block করবে।

---

# 3. Infinite Wait

```go
<-ch
```

sender না থাকলে forever wait।

---

# Important Interview Questions

## 1. Unbuffered channel কী?

Sender এবং receiver simultaneously ready থাকতে হয় এমন channel।

---

## 2. Why does unbuffered channel block?

কারণ buffer নাই।

Direct synchronization লাগে।

---

## 3. Capacity কত?

```text
0
```

---

## 4. Unbuffered channel কি asynchronous?

না।

এটা synchronous communication।

---

## 5. Buffered vs unbuffered main difference?

Buffered storage দেয়।
Unbuffered direct synchronization দেয়।

---

## 6. Why unbuffered channels are useful?

goroutine coordination এবং synchronization এর জন্য।

---

## 7. Can unbuffered channels prevent race condition?

Directly না।
কিন্তু proper synchronization provide করতে পারে।

---

## 8. What happens if no receiver exists?

Sender block করবে।

---

## 9. What happens if no sender exists?

Receiver block করবে।

---

# Advanced Interview Question

## 10. Why unbuffered channels are slower sometimes?

কারণ sender-receiver coordination wait লাগে।

---

## 11. When should you prefer unbuffered channel?

যখন:

- exact synchronization দরকার
- completion signal দরকার
- strict communication দরকার

---

# Important Interview Line

> “Unbuffered channels guarantee synchronization between goroutines.”

---

# Most Important Concept

Unbuffered channel এর মূল idea:

```text
Communication happens only when both sides are ready
```

এটাই Go concurrency এর core philosophy:

```text
Do not communicate by sharing memory;
instead, share memory by communicating.
```
