## 🔥 `log.Fatal` কী (Go তে)

👉 `log.Fatal` হলো Go এর logging function যা:

- message print করে
- তারপর program immediately stop (exit) করে

---

## 📌 Basic Example

```go id="lf9x2k"
package main

import (
	"log"
)

func main() {
	log.Println("Start app")

	log.Fatal("Something went wrong")

	log.Println("This will NOT run")
}
```

---

## ⚠️ Output

```bash id="x91kqp"
Start app
2026/01/01 Something went wrong
exit status 1
```

👉 শেষ line execute হবে না

---

# 🧠 ভিতরে কী হয়?

`log.Fatal()` basically 2 কাজ করে:

```go id="q2m8sd"
log.Print("message")
os.Exit(1)
```

---

# 🚨 গুরুত্বপূর্ণ পার্থক্য

## 1. `log.Fatal`

- print করে
- program exit করে
- defer কাজ করে না properly

---

## 2. `panic`

- panic throw করে
- defer কাজ করে
- recover করা যায়

---

## ⚖️ Panic vs Fatal

| Feature          | log.Fatal       | panic                  |
| ---------------- | --------------- | ---------------------- |
| Stop program     | Yes             | Yes                    |
| recover possible | ❌ No           | ✅ Yes                 |
| defer runs       | ❌ Not reliable | ✅ Yes                 |
| use case         | critical exit   | runtime crash handling |

---

# 💡 কখন use করবে?

👉 `log.Fatal` use করো যখন:

- config file missing
- database connection fail
- server start fail
- critical dependency missing

---

## 🔥 Real Example

```go id="p8kq2m"
func connectDB() {
	err := true

	if err {
		log.Fatal("Database connection failed")
	}
}
```

---

# ⚠️ Important Note

👉 `log.Fatal` program instantly terminate করে:

- cleanup code run নাও হতে পারে
- defer block skip হতে পারে (depends on situation)

---

# Simple understanding

- `log.Println` → শুধু print
- `log.Fatal` → print + stop program

---

চাও আমি তোমাকে **log.Panic vs log.Fatal vs panic vs os.Exit(1) comparison chart** দিয়ে দেই?
