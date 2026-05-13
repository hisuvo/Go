# Go Pointer Receiver (Bangla)

আগে আমরা Value Receiver দেখেছি।
এখন শিখবো **Pointer Receiver**।

Go language এ যখন method এর receiver হিসেবে struct এর **memory address** পাঠানো হয়, তখন তাকে Pointer Receiver বলে।

---

# Pointer Receiver কী?

```go
func (u *User) changeName() {
	u.name = "Rahim"
}
```

এখানে:

```go
*u
```

মানে method এর ভিতরে original object এর **pointer/address** গেছে।

এটা copy না।

---

# Real Life Example

ধরো তোমার বন্ধুকে তুমি বইয়ের photocopy না দিয়ে original বইটাই দিলে।

এখন সে বইতে লিখলে original বই change হয়ে যাবে।

ঠিক একইভাবে:

- Pointer Receiver original data modify করতে পারে
- কারণ এটা original memory কে access করে

---

# Basic Example

```go
package main

import "fmt"

type User struct {
	name string
}

func (u *User) changeName() {
	u.name = "Rahim"
}

func main() {
	user := User{name: "Suvo"}

	user.changeName()

	fmt.Println(user.name)
}
```

---

# Output

```txt
Rahim
```

---

# কেন original change হলো?

কারণ এখানে:

```go
(u *User)
```

মানে:

- struct এর copy যায়নি
- original object এর address গেছে

তাই method directly original memory modify করেছে।

---

# Memory Simulation

## Original Object

```txt
user
----
name = Suvo
address = 0x100
```

---

# Method Call

```go
user.changeName()
```

Go internally প্রায় এমন করে:

```go
changeName(&user)
```

মানে address পাঠায়।

---

# Receiver এ যা যায়

```txt
u ---> 0x100
```

মানে `u` original object কে point করছে।

---

# যখন change করি

```go
u.name = "Rahim"
```

তখন আসলে original memory change হয়।

---

# Visualization

```txt
main()
  |
  |---- user
         name = Suvo
         address = 0x100
                ↑
                |
             pointer
                |
           changeName()
```

---

# Pointer Receiver এর আসল Power

## Original Modify করতে পারে

```go
package main

import "fmt"

type Bank struct {
	balance int
}

func (b *Bank) deposit(amount int) {
	b.balance += amount
}

func main() {
	account := Bank{balance: 100}

	account.deposit(500)

	fmt.Println(account.balance)
}
```

---

# Output

```txt
600
```

---

# Real Life Example

ধরো:

- তোমার bank account এ 100 টাকা আছে
- deposit করলে original balance change হবে

তাই Pointer Receiver use করা হয়েছে।

---

# Value Receiver হলে কী হতো?

```go
func (b Bank) deposit(amount int) {
	b.balance += amount
}
```

এখানে copy change হতো।

Original balance same থাকতো।

---

# Large Struct এ Pointer Receiver Important

ধরো huge struct:

```go
type BigData struct {
	data [1000000]int
}
```

Value Receiver হলে পুরো data copy হবে।

এটা:

- slow
- memory expensive

কিন্তু Pointer Receiver:

```go
func (b *BigData) process() {
}
```

এখানে শুধু address pass হয়।

তাই:

- faster
- memory efficient

---

# Pointer Receiver Syntax

```go
func (variable *StructName) methodName() {

}
```

Example:

```go
func (u *User) update() {

}
```

---

# Go এর Interesting Feature

তুমি এভাবে call করতে পারো:

```go
user.changeName()
```

যদিও method হলো:

```go
func (u *User) changeName()
```

Go automatically address নেয়।

Internally:

```go
(&user).changeName()
```

---

# Pointer Receiver vs Value Receiver

| Feature            | Value Receiver | Pointer Receiver |
| ------------------ | -------------- | ---------------- |
| Copy হয়            | Yes            | No               |
| Original change হয় | No             | Yes              |
| Memory efficient   | কম             | বেশি             |
| Large struct       | Bad            | Good             |
| Read-only          | Good           | Possible         |
| Modify data        | No             | Yes              |

---

# কখন Pointer Receiver ব্যবহার করবো?

## Use when:

- Original data modify করতে হবে
- Large struct
- Performance important
- Memory save করতে হবে

---

# Developer Real Life Example

## User Profile Update

```go
package main

import "fmt"

type Profile struct {
	name string
	email string
}

func (p *Profile) updateEmail(newEmail string) {
	p.email = newEmail
}

func main() {
	user := Profile{
		name:  "Suvo",
		email: "old@gmail.com",
	}

	user.updateEmail("new@gmail.com")

	fmt.Println(user.email)
}
```

---

# Output

```txt
new@gmail.com
```

---

# Important Interview Question

## Question:

Why use pointer receiver in Go?

## Answer:

- To modify original struct
- To avoid copying large data
- Better memory efficiency

---

# Internal Execution Simulation

এই code:

```go
user.updateEmail("new@gmail.com")
```

Internally প্রায়:

```go
updateEmail(&user, "new@gmail.com")
```

---

# Very Important Rule

## যদি কিছু method pointer receiver হয়,

তাহলে usually সব method pointer receiver রাখা ভালো।

কারণ consistency maintain হয়।

---

# Summary

## Pointer Receiver

- Original object এর address নিয়ে কাজ করে
- Original data modify করতে পারে
- Large struct এ efficient
- Memory save করে
- Performance ভালো

---

# Short Formula

```txt
Pointer Receiver
=
Address of Original Object
```

---

# Final Visualization

## Value Receiver

```txt
Original ---> Copy ---> Method
```

## Pointer Receiver

```txt
Original <--- Method Direct Access
```
