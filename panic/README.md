Go তে **panic** আর **recover** হলো runtime error handle করার mechanism—but এটা normal error handling (error interface) না। এটা use করা হয় _unexpected crash_ handle করতে।

---

# 🔥 1. Panic কী?

👉 `panic` মানে program suddenly stop হয়ে যায়।

যখন Go কোনো serious problem পায় তখন panic হয়।

### Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Start")
	panic("something went wrong")
	fmt.Println("End")
}
```

### Output:

```
Start
panic: something went wrong
```

👉 নিচের code আর execute হবে না

---

# ⚠️ Panic কখন হয়?

- Array index out of range
- Nil pointer access
- Manual panic()
- Critical runtime failure

---

# 🔥 2. Recover কী?

👉 `recover()` panic থেকে program কে আবার চালাতে সাহায্য করে।

👉 এটা শুধু `defer` এর ভিতরে কাজ করে।

---

# 💡 Basic Example (panic + recover)

```go
package main

import "fmt"

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start")
	panic("crash happened")
	fmt.Println("End")
}

func main() {
	test()
	fmt.Println("Program continues...")
}
```

---

# 🟢 Output:

```
Start
Recovered from panic: crash happened
Program continues...
```

---

# 🧠 কিভাবে কাজ করে?

## Flow:

```
panic happens
     ↓
defer runs
     ↓
recover catches panic
     ↓
program continues
```

---

# ⚡ Important Rules

## 1. recover শুধু defer এর ভিতরে কাজ করে

```go
recover() // ❌ কাজ করবে না
```

```go
defer func() {
	recover() // ✅ কাজ করবে
}()
```

---

## 2. recover না থাকলে program crash করবে

---

# 🔥 Real Use Case (Server)

ধরো API server:

```go
func handler() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("API crashed but recovered:", r)
		}
	}()

	// risky code
	panic("database connection failed")
}
```

👉 server পুরো crash হবে না

---

# ⚖️ Panic vs Error

| Feature  | error          | panic            |
| -------- | -------------- | ---------------- |
| Use case | expected issue | unexpected crash |
| handling | manual check   | recover          |
| flow     | normal         | stop execution   |
| usage    | recommended    | rare             |

---

# 🚀 Best Practice

👉 Go experts বলে:

- ❌ panic avoid করা উচিত (normal logic এ)
- ✅ error return করা উচিত
- ✅ panic only use for:
  - unrecoverable bug
  - initialization failure
  - truly fatal condition

---

# 🧠 Simple analogy

- `error` → “problem, but handleable”
- `panic` → “system crash situation”
- `recover` → “emergency rescue button”

---

চাও আমি তোমাকে **real API server middleware (panic recover layer)** দেখাই?
