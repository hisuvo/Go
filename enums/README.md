# Enum in Go (Golang) — Full Details

Go language এ direct `enum` keyword নেই like C++, Java।
কিন্তু Go তে `const` + `iota` ব্যবহার করে enum-এর মতো behavior তৈরি করা হয়।

---

# Why Enum Use?

Enum use করা হয় fixed values define করার জন্য।

Example:

- Order Status
- Payment Status
- User Role
- HTTP Status
- File Permission
- Log Level

---

# Basic Enum Example

```go
package main

import "fmt"

type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)

func main() {
	fmt.Println(Pending)  // 0
	fmt.Println(Approved) // 1
	fmt.Println(Rejected) // 2
}
```

---

# Here What Happening?

```go
type Status int
```

এখানে custom type তৈরি হলো।

---

```go
const (
	Pending Status = iota
	Approved
	Rejected
)
```

`iota` auto increment করে।

Equivalent:

```go
Pending  = 0
Approved = 1
Rejected = 2
```

---

# Memory View

```go
Approved
```

Actually store হয়:

```go
1
```

কিন্তু type হলো:

```go
Status
```

So type safety পাওয়া যায়।

---

# Why Better Than Normal Int?

Without enum:

```go
var status int = 100
```

এখানে invalid value যেতে পারে।

But:

```go
var status Status
```

এখন logically fixed values use হবে।

---

# Real Industry Example — Order System

## order/status.go

```go
package order

type OrderStatus int

const (
	Pending OrderStatus = iota
	Paid
	Shipped
	Delivered
	Cancelled
)
```

---

# order/order.go

```go
package order

import "fmt"

type Order struct {
	ID     int
	Status OrderStatus
}

func UpdateOrderStatus(order *Order, status OrderStatus) {
	order.Status = status
}

func PrintOrder(order Order) {
	fmt.Println("Order ID:", order.ID)
	fmt.Println("Status:", order.Status)
}
```

---

# main.go

```go
package main

import (
	"myapp/order"
)

func main() {

	myOrder := order.Order{
		ID:     1,
		Status: order.Pending,
	}

	order.UpdateOrderStatus(&myOrder, order.Shipped)

	order.PrintOrder(myOrder)
}
```

---

# Problem With Default Enum Printing

Output:

```go
Status: 2
```

Human readable না।

---

# Solution — String Method

## status.go

```go
package order

type OrderStatus int

const (
	Pending OrderStatus = iota
	Paid
	Shipped
	Delivered
	Cancelled
)

func (s OrderStatus) String() string {

	switch s {

	case Pending:
		return "Pending"

	case Paid:
		return "Paid"

	case Shipped:
		return "Shipped"

	case Delivered:
		return "Delivered"

	case Cancelled:
		return "Cancelled"

	default:
		return "Unknown"
	}
}
```

---

# Now Output

```go
Status: Shipped
```

---

# Real Backend API Use Case

Imagine:

## Database

| id  | status |
| --- | ------ |
| 1   | 2      |

---

Backend:

```go
2 => Shipped
```

Frontend:

```json
{
  "status": "Shipped"
}
```

---

# Enum With Custom Values

```go
type Role int

const (
	Admin Role = 1
	Moderator Role = 5
	User Role = 10
)
```

---

# Enum With Skipping Values

```go
const (
	A = iota // 0
	B        // 1
	_
	C        // 3
)
```

---

# String Enum Style

Sometimes integer না use করে string use হয়।

```go
type PaymentStatus string

const (
	Success PaymentStatus = "SUCCESS"
	Failed  PaymentStatus = "FAILED"
	Pending PaymentStatus = "PENDING"
)
```

---

# Why String Enum Used?

Because:

- Database readable
- API readable
- Debug easy
- Logging easy

---

# Real JWT/Auth Example

```go
type UserRole string

const (
	Admin UserRole = "ADMIN"
	User  UserRole = "USER"
	Guest UserRole = "GUEST"
)
```

Middleware:

```go
if user.Role == Admin {
	fmt.Println("Access Granted")
}
```

---

# iota Advanced Example

```go
const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)
```

Equivalent:

```go
KB = 1024
MB = 1048576
GB = 1073741824
```

Real use:

- File system
- Storage calculation
- Cloud services

---

# Best Practice

## 1. Create Separate File

```plaintext
user/
 ├── role.go
 ├── user.go
```

---

## 2. Use String Method

Always implement:

```go
func (r Role) String() string
```

---

## 3. Validate Enum

```go
func IsValidStatus(s OrderStatus) bool {

	switch s {

	case Pending, Paid, Shipped:
		return true
	}

	return false
}
```

---

# Common Beginner Mistake

## Wrong

```go
const (
	Pending = iota
)
```

No custom type.

---

## Better

```go
type Status int

const (
	Pending Status = iota
)
```

---

# Enum vs Constant

| Feature     | Enum Style | Normal Const |
| ----------- | ---------- | ------------ |
| Type Safety | Yes        | No           |
| Grouping    | Yes        | No           |
| Readability | High       | Medium       |
| Backend Use | Excellent  | Limited      |

---

# Real-Life Example

Think traffic signal:

| Signal | Value |
| ------ | ----- |
| Red    | Stop  |
| Yellow | Wait  |
| Green  | Go    |

This fixed set = enum concept.

---

# Important Interview Questions

## 1. Go language এ enum আছে?

No direct enum keyword নেই।
`const + iota` দিয়ে enum behavior তৈরি করা হয়।

---

## 2. iota কী?

`iota` হলো Go constant generator।
It auto increments inside const block.

---

## 3. Why use custom type with enum?

Type safety and readability জন্য।

---

## 4. Difference between int enum and string enum?

| Int Enum     | String Enum      |
| ------------ | ---------------- |
| Faster       | More readable    |
| Small memory | Better API/debug |
| DB optimized | Human friendly   |

---

## 5. Why implement String() method?

Human readable output পাওয়ার জন্য।

---

## 6. Where enum used in real backend?

- Order status
- Payment status
- User roles
- Notification types
- File permissions
- HTTP methods
- Log levels

---

# Senior Level Interview Question

## Question:

Why Go does not provide real enum keyword?

## Answer:

Go keeps language simple.
`const + iota` already provides lightweight enum-like behavior without adding extra complexity.

---

# Real Industry Folder Structure

```plaintext
project/
│
├── user/
│   ├── role.go
│   ├── user.go
│
├── order/
│   ├── status.go
│   ├── order.go
│
├── payment/
│   ├── payment_status.go
│
└── main.go
```

---

# Final Summary

Go enum system mainly built using:

- `type`
- `const`
- `iota`

Most real backend applications use enum for:

- status management
- auth roles
- workflow states
- event types

This is heavily used in:

- microservices
- REST API
- gRPC
- event-driven systems
- distributed systems
