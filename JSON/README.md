Go-তে JSON খুব গুরুত্বপূর্ণ, কারণ backend API, database communication, microservice, frontend communication — সব জায়গায় JSON use হয়।
Go-তে JSON handle করার জন্য built-in `encoding/json` package ব্যবহার করা হয়।

---

# JSON কী?

JSON = JavaScript Object Notation

এটি data exchange format।

Example JSON:

```json
{
  "name": "Suvo",
  "age": 24,
  "skills": ["Go", "React", "MongoDB"]
}
```

---

# Go-তে JSON এর Main কাজ

Go-তে সাধারণত ২টা কাজ হয়:

| কাজ       | Meaning          |
| --------- | ---------------- |
| Marshal   | Go Struct → JSON |
| Unmarshal | JSON → Go Struct |

---

# 1. Marshal (Struct → JSON)

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {

	user := User{
		Name: "Suvo",
		Age:  24,
	}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
```

---

# Output

```json
{ "Name": "Suvo", "Age": 24 }
```

---

# এখানে কী হচ্ছে?

```go
json.Marshal(user)
```

এটি Go struct কে JSON bytes এ convert করছে।

Return দেয়:

```go
[]byte
```

তাই:

```go
string(jsonData)
```

দিয়ে print করা হয়েছে।

---

# 2. JSON Tag ব্যবহার

Real project এ আমরা চাই:

```json
{
  "name": "Suvo"
}
```

না যে:

```json
{
  "Name": "Suvo"
}
```

তাই JSON tag ব্যবহার করা হয়।

---

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	user := User{
		Name: "Suvo",
		Age:  24,
	}

	jsonData, _ := json.Marshal(user)

	fmt.Println(string(jsonData))
}
```

---

# Output

```json
{ "name": "Suvo", "age": 24 }
```

---

# 3. Unmarshal (JSON → Struct)

এটি সবচেয়ে বেশি backend API তে use হয়।

---

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	jsonData := `{
		"name": "Suvo",
		"age": 24
	}`

	var user User

	err := json.Unmarshal([]byte(jsonData), &user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
```

---

# Output

```go
Suvo
24
```

---

# কেন `&user` দিলাম?

কারণ:

```go
json.Unmarshal()
```

data modify করে।

তাই memory address লাগে।

এই কারণে pointer ব্যবহার হয়।

---

# Real Backend Example

ধরো frontend থেকে request আসছে:

```json
{
  "email": "suvo@gmail.com",
  "password": "123456"
}
```

Go backend এ:

```go
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
```

তারপর:

```go
json.NewDecoder(r.Body).Decode(&login)
```

এভাবে request body parse করা হয়।

---

# API Response Example

```go
response := map[string]string{
	"message": "Login Success",
}

json.NewEncoder(w).Encode(response)
```

Frontend এ JSON চলে যাবে।

---

# Struct Tag Important

```go
type User struct {
	Name string `json:"name"`
}
```

এখানে:

| Part   | Meaning         |
| ------ | --------------- |
| json   | package name    |
| "name" | JSON field name |

---

# `omitempty`

যদি value empty হয় তাহলে JSON এ যাবে না।

---

## Example

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}
```

যদি:

```go
Age = 0
```

তাহলে output:

```json
{
  "name": "Suvo"
}
```

---

# Nested JSON Example

## JSON

```json
{
  "name": "Suvo",
  "address": {
    "city": "Dhaka",
    "country": "Bangladesh"
  }
}
```

---

## Go Struct

```go
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}
```

---

# JSON Array Example

## JSON

```json
[
  {
    "name": "A"
  },
  {
    "name": "B"
  }
]
```

---

## Go

```go
type User struct {
	Name string `json:"name"`
}

var users []User
```

---

# Real Industry Use Case

Backend API:

1. Frontend JSON পাঠায়
2. Go backend JSON parse করে
3. Database এ save করে
4. আবার JSON response পাঠায়

---

# Interview Questions

## 1. Marshal vs Unmarshal difference?

| Marshal       | Unmarshal     |
| ------------- | ------------- |
| Struct → JSON | JSON → Struct |

---

## 2. কেন struct field Capital letter হতে হয়?

```go
type User struct {
	Name string
}
```

কারণ Go only exported fields encode/decode করতে পারে।

এটি কাজ করবে না:

```go
type User struct {
	name string
}
```

---

## 3. কেন pointer লাগে Unmarshal এ?

কারণ function data modify করে।

---

# Most Important Functions

| Function          | কাজ            |
| ----------------- | -------------- |
| json.Marshal()    | Struct → JSON  |
| json.Unmarshal()  | JSON → Struct  |
| json.NewEncoder() | Response write |
| json.NewDecoder() | Request read   |

---

# Real Production Example

```go
package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	response := map[string]any{
		"success": true,
		"user":    user,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/user", createUser)

	http.ListenAndServe(":8080", nil)
}
```

---

# Request

```json
{
  "name": "Suvo",
  "age": 24
}
```

---

# Response

```json
{
  "success": true,
  "user": {
    "name": "Suvo",
    "age": 24
  }
}
```
