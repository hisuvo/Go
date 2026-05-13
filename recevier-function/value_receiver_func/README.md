# Go Value Receiver Function (Bangla)

Go language এ যখন কোনো `struct` এর সাথে function attach করা হয়, তখন সেটাকে **method** বলে।
এই method এর receiver দুই ধরনের হতে পারে:

1. **Value Receiver**
2. **Pointer Receiver**

আজ আমরা শুধু **Value Receiver** নিয়ে বিস্তারিত জানবো।

---

# Value Receiver কী?

যখন method এ struct এর **copy** যায়, তখন তাকে Value Receiver বলে।

```go
func (u User) showName() {
	fmt.Println(u.name)
}
```

এখানে:

```go
(u User)
```

এটাই receiver।

মানে `User` struct এর একটা **copy** method এর ভিতরে গেছে।

---

# Real Life Example

ধরো তোমার কাছে একটা বই আছে।

তুমি বন্ধুকে বইয়ের **ফটোকপি** দিলে।

এখন সে ফটোকপিতে লিখলে বা কাটাকাটি করলে তোমার original বইয়ের কিছু হবে না।

ঠিক একইভাবে:

- Value Receiver original data change করে না
- শুধু copy নিয়ে কাজ করে

---

# Basic Example

```go
package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s Student) showInfo() {
	fmt.Println("Name:", s.name)
	fmt.Println("Age:", s.age)
}

func main() {
	st := Student{
		name: "Suvo",
		age:  23,
	}

	st.showInfo()
}
```

---

# Memory Simulation

ধরো:

```go
st := Student{name:"Suvo", age:23}
```

Memory:

```txt
Original Object
---------------
name = Suvo
age  = 23
```

যখন:

```go
st.showInfo()
```

তখন Go internally copy তৈরি করে:

```txt
Copied Object
-------------
name = Suvo
age  = 23
```

Method এই copied object নিয়ে কাজ করে।

---

# Important Thing

## Value Receiver এ change করলে original change হয় না

```go
package main

import "fmt"

type User struct {
	name string
}

func (u User) changeName() {
	u.name = "Rahim"
	fmt.Println("Inside Method:", u.name)
}

func main() {
	user := User{name: "Suvo"}

	user.changeName()

	fmt.Println("Main Function:", user.name)
}
```

---

# Output

```txt
Inside Method: Rahim
Main Function: Suvo
```

---

# কেন এমন হলো?

কারণ:

```go
(u User)
```

এখানে পুরো struct এর copy গেছে।

তাই:

```go
u.name = "Rahim"
```

শুধু copy change করেছে।

Original `user` change হয়নি।

---

# Memory Visualization

## Original

```txt
user
----
name = Suvo
```

## Method এ Copy

```txt
u
-
name = Suvo
```

তারপর:

```go
u.name = "Rahim"
```

হলো:

```txt
u
-
name = Rahim
```

কিন্তু original এখনও:

```txt
user
----
name = Suvo
```

---

# Value Receiver কবে ব্যবহার করবো?

যখন:

- Original data change করতে চাই না
- Small struct
- Read-only operation
- শুধু data show/calculation করতে চাই

---

# Real Life Developer Example

## Product Price দেখানো

```go
package main

import "fmt"

type Product struct {
	name  string
	price int
}

func (p Product) display() {
	fmt.Println("Product:", p.name)
	fmt.Println("Price:", p.price)
}

func main() {
	product := Product{
		name:  "Laptop",
		price: 50000,
	}

	product.display()
}
```

এখানে শুধু data দেখানো হচ্ছে।

Original product modify দরকার নেই।

তাই Value Receiver perfect।

---

# Value Receiver এর সুবিধা

## 1. Original Safe থাকে

Accidentally original data change হয় না।

---

## 2. Readability ভালো

বোঝা যায় method data modify করবে না।

---

## 3. Small Data এর জন্য Fast

ছোট struct হলে copy করা সমস্যা না।

---

# অসুবিধা

## Large Struct হলে copy cost বেশি

ধরো huge struct:

```go
type BigData struct {
	data [100000]int
}
```

এখন Value Receiver use করলে পুরো data copy হবে।

এটা memory expensive।

---

# Value Receiver vs Pointer Receiver

| Feature            | Value Receiver      | Pointer Receiver |
| ------------------ | ------------------- | ---------------- |
| Copy হয়            | Yes                 | No               |
| Original change হয় | No                  | Yes              |
| Memory বেশি লাগে   | Large struct এ বেশি | কম               |
| Safe               | বেশি                | কম               |
| Use Case           | Read only           | Modify           |

---

# Important Interview Question

## Question:

Why original value does not change in value receiver?

## Answer:

Because Go passes a copy of the struct to the method when using value receiver.

---

# Internal Simulation

এই code:

```go
user.changeName()
```

Go internally প্রায় এমন করে:

```go
temp := user
temp.changeName()
```

মানে copy বানিয়ে method call করে।

---

# Summary

## Value Receiver

- Struct এর copy নিয়ে কাজ করে
- Original data change করে না
- Small struct এর জন্য ভালো
- Read-only method এ বেশি use হয়

---

# Short Formula

```txt
Value Receiver
=
Copy of Original Object
```

---

# Syntax

```go
func (variable StructName) methodName() {

}
```

Example:

```go
func (u User) show() {

}
```
