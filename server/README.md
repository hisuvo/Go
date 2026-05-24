Go language এ server তৈরি করার সবচেয়ে common way হলো `net/http` package ব্যবহার করা।

## Basic HTTP Server Example

```go
package main

import (
	"fmt"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello SUVO! Go Server Running...")
}

func main() {

	// route create
	http.HandleFunc("/", home)

	fmt.Println("Server running on port 5000")

	// server start
	http.ListenAndServe(":5000", nil)
}
```

---

# Code Explanation

## 1. net/http package

```go
import "net/http"
```

এটা Go এর built-in package
যা দিয়ে HTTP server তৈরি করা হয়।

---

## 2. Handler Function

```go
func home(w http.ResponseWriter, r *http.Request)
```

এখানে:

- `w http.ResponseWriter`
  → client কে response পাঠায়

- `r *http.Request`
  → client request receive করে

---

## 3. Response পাঠানো

```go
fmt.Fprintln(w, "Hello SUVO! Go Server Running...")
```

Browser এ এই text show করবে।

---

## 4. Route Create

```go
http.HandleFunc("/", home)
```

মানে:

| Route | Function |
| ----- | -------- |
| `/`   | `home`   |

---

## 5. Server Start

```go
http.ListenAndServe(":5000", nil)
```

মানে:

- server run হবে
- port `5000` এ

---

# Run How?

## Step 1

File create:

```bash
main.go
```

## Step 2

Run:

```bash
go run main.go
```

---

# Browser Open

```txt
http://localhost:5000
```

Output:

```txt
Hello SUVO! Go Server Running...
```

---

# Multiple Route Example

```go
package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contact Page")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("Server running on port 5000")

	http.ListenAndServe(":5000", nil)
}
```

---

# Real Life Server Flow

```txt
Browser Request
       ↓
Go Server
       ↓
Route Match
       ↓
Handler Function
       ↓
Response Send
```

---

# Important Interview Question

## Question 1:

What is `http.HandleFunc`?

Answer:
Route এর সাথে handler function connect করে।

---

## Question 2:

Why `ResponseWriter` use হয়?

Answer:
Client কে response পাঠানোর জন্য।

---

## Question 3:

Why request parameter pointer?

```go
r *http.Request
```

কারণ request object বড় structure।
Pointer দিলে memory efficient হয়।

---

# Production Way

Real project এ সাধারণত use হয়:

- Gin
- Fiber
- Echo

কারণ এগুলো fast এবং easy routing দেয়।

---

# Gin Example

```go
package main

import "github.com/gin-gonic/gin"

func main() {

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello SUVO",
		})
	})

	server.Run(":5000")
}
```

---

# http.NewServerMux

---

`http.NewServeMux()` হলো Go এর built-in HTTP request router।

এটা URL path অনুযায়ী request handle করে।

```go
mux := http.NewServeMux()
```

এখানে:

- `http` → Go এর `net/http` package
- `NewServeMux()` → নতুন router/multiplexer তৈরি করে
- `mux` → routes handle করবে

---

## কেন ব্যবহার হয়?

একাধিক route handle করার জন্য।

যেমন:

- `/`
- `/about`
- `/login`

---

## Example

```go
package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)

	http.ListenAndServe(":8080", mux)
}
```

---

## এখানে কী হচ্ছে?

### 1. ServeMux তৈরি

```go
mux := http.NewServeMux()
```

এটা route manager তৈরি করছে।

---

### 2. Route register

```go
mux.HandleFunc("/", home)
```

মানে:

যদি user `/` এ যায় → `home()` function run হবে।

---

### 3. Server start

```go
http.ListenAndServe(":8080", mux)
```

`mux` কে server এর handler হিসেবে দেওয়া হয়েছে।

---

# ServeMux internally কী করে?

Request আসলে:

```txt
/about
```

ServeMux check করে:

```txt
"/about" => about handler
```

তারপর সেই function execute করে।

---

# Without ServeMux

তুমি direct handlerও দিতে পারো:

```go
http.ListenAndServe(":8080", nil)
```

`nil` দিলে Go default global mux ব্যবহার করে:

```go
http.HandleFunc("/", home)
```

কিন্তু বড় project এ custom mux ভালো practice।

---

# Interview Question

## What is ServeMux in Go?

ServeMux is an HTTP request multiplexer that matches incoming requests against registered routes and dispatches them to the correct handler.

---

## Why use custom ServeMux?

- Better route management
- Avoid global state
- Cleaner architecture
- Easier testing

---

# Real Project Example

```go
mux := http.NewServeMux()

mux.HandleFunc("/api/users", getUsers)
mux.HandleFunc("/api/login", login)
mux.HandleFunc("/api/products", products)

server := &http.Server{
	Addr:    ":8080",
	Handler: mux,
}

server.ListenAndServe()
```

এভাবে production app এ routes organize করা হয়।
