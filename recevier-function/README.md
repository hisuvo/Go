# Go Receiver Function Complete Guide

## What is a Receiver Function in Go?

In Go, a receiver function is a function that belongs to a specific type.

It is similar to a method in Object-Oriented Programming languages like Java, C#, or Python.

Receiver functions allow us to attach behavior to structs or custom types.

---

# Basic Syntax

```go
func (receiver Type) methodName() {
    // code
}
```

Example:

```go
package main

import "fmt"

type Person struct {
    name string
    age  int
}

// Receiver Function
func (p Person) introduce() {
    fmt.Println("My name is", p.name)
    fmt.Println("My age is", p.age)
}

func main() {
    user := Person{
        name: "Suvo",
        age:  22,
    }

    user.introduce()
}
```

Output:

```bash
My name is Suvo
My age is 22
```

---

# Understanding the Receiver

```go
func (p Person) introduce()
```

Here:

| Part      | Meaning           |
| --------- | ----------------- |
| p         | Receiver variable |
| Person    | Receiver type     |
| introduce | Method name       |

This means:

> The introduce() method belongs to the Person type.

---

# Why Receiver Functions are Important

Receiver functions help developers:

- Organize code better
- Write clean and reusable code
- Create object-like behavior
- Attach functionality directly to structs
- Build scalable backend systems

---

# Real Life Analogy

Imagine a mobile phone.

The phone has:

- data → brand, model, battery
- behavior → call(), charge(), shutdown()

In Go:

```go
phone.call()
phone.charge()
```

Receiver functions make this possible.

---

# Example Without Receiver Function

```go
package main

import "fmt"

type Person struct {
    name string
}

func introduce(p Person) {
    fmt.Println("Hello", p.name)
}

func main() {
    user := Person{name: "Suvo"}

    introduce(user)
}
```

This works.

But receiver functions are cleaner:

```go
user.introduce()
```

instead of:

```go
introduce(user)
```

---

# Types of Receiver Functions

Go has two types of receivers:

1. Value Receiver
2. Pointer Receiver

---

# 1. Value Receiver

A value receiver works on a copy of the original value.

Syntax:

```go
func (p Person) methodName() {

}
```

Example:

```go
package main

import "fmt"

type Person struct {
    name string
}

func (p Person) changeName() {
    p.name = "Rahim"
}

func main() {
    user := Person{name: "Karim"}

    user.changeName()

    fmt.Println(user.name)
}
```

Output:

```bash
Karim
```

---

# Why Did the Name Not Change?

Because:

```go
func (p Person)
```

creates a copy of the original struct.

The original data remains unchanged.

Visual Representation:

```text
Original Struct → Copy Created → Method Uses Copy
```

---

# Value Receiver Memory Flow

```text
user ------> [Person{name:"Karim"}]
                  |
                  | copied
                  v
              p ------> [Person{name:"Karim"}]
```

Changes happen only inside the copy.

---

# When to Use Value Receiver

Use value receivers when:

- You do not need to modify the original data
- Struct is small
- Immutable behavior is preferred
- Simplicity is enough

Example:

```go
func (p Person) showName() {
    fmt.Println(p.name)
}
```

---

# 2. Pointer Receiver

Pointer receiver directly accesses the original memory address.

Syntax:

```go
func (p *Person) methodName() {

}
```

Example:

```go
package main

import "fmt"

type Person struct {
    name string
}

func (p *Person) changeName() {
    p.name = "Rahim"
}

func main() {
    user := Person{name: "Karim"}

    user.changeName()

    fmt.Println(user.name)
}
```

Output:

```bash
Rahim
```

---

# Why Did It Change?

Because pointer receivers modify the original value.

```go
func (p *Person)
```

The `*` means:

> work with the original memory address.

---

# Pointer Receiver Memory Flow

```text
user ------> [Person{name:"Karim"}]
                 ^
                 |
                 p
```

No copy is created.

---

# Value Receiver vs Pointer Receiver

| Feature                  | Value Receiver | Pointer Receiver |
| ------------------------ | -------------- | ---------------- |
| Uses Copy                | Yes            | No               |
| Changes Original Value   | No             | Yes              |
| Memory Efficient         | Less           | More             |
| Large Struct Recommended | No             | Yes              |
| Safer for Immutable Data | Yes            | No               |

---

# Important Go Behavior

Even if a method uses a pointer receiver:

```go
func (p *Person) update() {}
```

You can still call it like this:

```go
user.update()
```

Go automatically handles the pointer internally.

---

# Example with Multiple Methods

```go
package main

import "fmt"

type BankAccount struct {
    owner   string
    balance int
}

func (b *BankAccount) deposit(amount int) {
    b.balance += amount
}

func (b *BankAccount) withdraw(amount int) {
    b.balance -= amount
}

func (b BankAccount) showBalance() {
    fmt.Println("Owner:", b.owner)
    fmt.Println("Balance:", b.balance)
}

func main() {
    account := BankAccount{
        owner:   "Suvo",
        balance: 1000,
    }

    account.deposit(500)
    account.withdraw(200)

    account.showBalance()
}
```

Output:

```bash
Owner: Suvo
Balance: 1300
```

---

# Receiver Functions with Custom Types

Receivers are not limited to structs.

You can use them with custom types too.

Example:

```go
package main

import "fmt"

type Number int

func (n Number) square() Number {
    return n * n
}

func main() {
    var x Number = 5

    fmt.Println(x.square())
}
```

Output:

```bash
25
```

---

# Receiver Naming Convention

Developers usually use short receiver names.

Examples:

| Struct  | Receiver |
| ------- | -------- |
| Person  | p        |
| User    | u        |
| Account | a        |
| Server  | s        |
| Product | p        |

Good:

```go
func (u User) login() {}
```

Avoid:

```go
func (myVeryLongReceiverName User) login() {}
```

---

# Method Sets in Go

Go internally keeps track of which methods belong to which type.

Example:

```go
type User struct {}

func (u User) show() {}
func (u *User) update() {}
```

Method Set:

| Type   | Accessible Methods |
| ------ | ------------------ |
| User   | show()             |
| \*User | show(), update()   |

---

# Receiver Functions and Interfaces

Receiver methods are heavily used with interfaces.

Example:

```go
package main

import "fmt"

type Speaker interface {
    speak()
}

type Human struct {
    name string
}

func (h Human) speak() {
    fmt.Println("Hello, I am", h.name)
}

func main() {
    h := Human{name: "Suvo"}

    var s Speaker = h

    s.speak()
}
```

Output:

```bash
Hello, I am Suvo
```

---

# Real Backend Example

Example of a service in backend development:

```go
package main

import "fmt"

type UserService struct {}

func (u UserService) CreateUser(name string) {
    fmt.Println("Creating user:", name)
}

func (u UserService) DeleteUser(id int) {
    fmt.Println("Deleting user:", id)
}

func main() {
    service := UserService{}

    service.CreateUser("Suvo")
    service.DeleteUser(10)
}
```

---

# Common Beginner Mistakes

## Mistake 1

Trying to change value using value receiver.

Wrong:

```go
func (p Person) update() {
    p.name = "New"
}
```

Correct:

```go
func (p *Person) update() {
    p.name = "New"
}
```

---

## Mistake 2

Mixing receiver types unnecessarily.

Bad Practice:

```go
func (u User) show() {}
func (u *User) update() {}
```

Usually use one consistent style.

---

# Best Practices

## Use Pointer Receiver When:

- Modifying struct data
- Struct is large
- Better performance is needed
- Avoid copying data

## Use Value Receiver When:

- Data should not change
- Struct is small
- Immutable behavior is preferred

---

# Performance Consideration

Large structs copied repeatedly can reduce performance.

Example:

```go
type BigData struct {
    data [100000]int
}
```

Using value receivers here creates large copies.

Pointer receivers are better.

---

# Receiver Function vs Normal Function

## Normal Function

```go
add(2, 3)
```

## Receiver Function

```go
user.login()
```

Receiver functions make code more readable.

---

# Advanced Example

```go
package main

import "fmt"

type Counter struct {
    value int
}

func (c *Counter) increment() {
    c.value++
}

func (c *Counter) decrement() {
    c.value--
}

func (c Counter) show() {
    fmt.Println("Current Value:", c.value)
}

func main() {
    counter := Counter{}

    counter.increment()
    counter.increment()
    counter.increment()

    counter.decrement()

    counter.show()
}
```

Output:

```bash
Current Value: 2
```

---

# Important Interview Questions

## What is a receiver function in Go?

A receiver function is a method attached to a specific type.

---

## Difference between value receiver and pointer receiver?

| Value Receiver         | Pointer Receiver        |
| ---------------------- | ----------------------- |
| Works on copy          | Works on original value |
| Cannot modify original | Can modify original     |
| More copying           | Better performance      |

---

## Why use pointer receivers?

- Modify original data
- Better performance
- Avoid large memory copies

---

# Final Summary

Receiver functions are one of the most important concepts in Go.

They help developers:

- Build clean applications
- Organize business logic
- Write reusable code
- Create scalable backend systems
- Work efficiently with interfaces

Most real-world Go applications heavily use receiver functions.

Especially in:

- REST APIs
- Microservices
- Database services
- Authentication systems
- Web servers
- Clean Architecture
- Domain-driven applications

---

# Quick Revision

| Topic                         | Key Idea                    |
| ----------------------------- | --------------------------- |
| Receiver Function             | Function attached to a type |
| Value Receiver                | Uses copy                   |
| Pointer Receiver              | Uses original memory        |
| Best for Modification         | Pointer Receiver            |
| Best for Small Immutable Data | Value Receiver              |
| OOP Style                     | user.login()                |

---

# Practice Tasks

Try building:

1. Student Management System
2. Bank Account System
3. Shopping Cart
4. Todo Application
5. Inventory System

using receiver functions.

This will make your Go fundamentals much stronger.
