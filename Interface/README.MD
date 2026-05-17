# Go Interface কী?

Go-তে `interface` হলো এমন একটি টাইপ যা বলে দেয়:

> “যে struct এই method গুলো implement করবে, সে এই interface ব্যবহার করতে পারবে।”

সহজভাবে:

- Interface = Contract / Rules
- Struct = বাস্তব object
- Method = কাজ

---

# Real Life Example 🚗

ধরো, একটি অ্যাপ আছে যেখানে বিভিন্ন vehicle আছে।

- Car
- Bike
- Truck

সব vehicle-এর একটি common কাজ আছে:

```go
Start()
```

এখন সবগুলোর জন্য আলাদা আলাদা function না লিখে আমরা interface ব্যবহার করবো।

---

# Basic Industry Style Example

```go
package main

import "fmt"

// Interface
type Vehicle interface {
	Start()
}

// Car struct
type Car struct {
	Name string
}

// Car method
func (c Car) Start() {
	fmt.Println(c.Name, "car started")
}

// Bike struct
type Bike struct {
	Name string
}

// Bike method
func (b Bike) Start() {
	fmt.Println(b.Name, "bike started")
}

// Common function
func EngineStart(v Vehicle) {
	v.Start()
}

func main() {
	car := Car{Name: "BMW"}
	bike := Bike{Name: "Yamaha"}

	EngineStart(car)
	EngineStart(bike)
}
```

---

# Output

```bash
BMW car started
Yamaha bike started
```

---

# এখানে কী হলো?

`Vehicle` interface বলেছে:

```go
Start()
```

যে struct এর মধ্যে `Start()` method থাকবে, সে automatically `Vehicle` interface implement করবে।

Go-তে:

❌ `implements` keyword লাগে না
✅ Automatically implement হয়

---

# Memory Simulation 🧠

যখন:

```go
EngineStart(car)
```

call হয়,

তখন Go internally দেখে:

```text
Car struct এর মধ্যে কি Start() method আছে?
```

যদি থাকে:

✅ তাহলে `Vehicle` interface satisfy করবে।

---

# Industry Real Example — Payment Gateway 💳

বাস্তব software company-তে interface অনেক বেশি use হয় payment system-এ।

যেমন:

- Stripe
- PayPal
- SSLCommerz
- bKash

সব payment provider-এর কাজ একই:

```text
Pay()
Refund()
```

---

# Industry Example

```go
package main

import "fmt"

type PaymentGateway interface {
	Pay(amount float64)
	Refund(amount float64)
}

// bKash
type BKash struct{}

func (b BKash) Pay(amount float64) {
	fmt.Println("Paid using bKash:", amount)
}

func (b BKash) Refund(amount float64) {
	fmt.Println("Refund from bKash:", amount)
}

// Stripe
type Stripe struct{}

func (s Stripe) Pay(amount float64) {
	fmt.Println("Paid using Stripe:", amount)
}

func (s Stripe) Refund(amount float64) {
	fmt.Println("Refund from Stripe:", amount)
}

// Common payment processor
func ProcessPayment(p PaymentGateway) {
	p.Pay(500)
}

func main() {
	bkash := BKash{}
	stripe := Stripe{}

	ProcessPayment(bkash)
	ProcessPayment(stripe)
}
```

---

# কেন Industry-তে Interface Important?

## 1. Loose Coupling

তুমি code change না করেই নতুন payment system add করতে পারো।

যেমন:

```text
Rocket
Nagad
Payoneer
```

শুধু method implement করলেই হবে।

---

## 2. Scalable Architecture

Big project-এ easily feature add করা যায়।

---

## 3. Testing Easy হয়

Mock object বানানো যায়।

Example:

```go
type FakePayment struct{}
```

Testing-এ real payment ছাড়াই test করা যায়।

---

# Interface ছাড়া Problem 😵

যদি interface না ব্যবহার করো:

```go
func PayWithBkash()
func PayWithStripe()
func PayWithPaypal()
```

তাহলে:

- Code duplicate হবে
- Maintain করা hard হবে
- New provider add করা difficult হবে

---

# Empty Interface

```go
interface{}
```

এটা সব ধরনের value রাখতে পারে।

```go
var data interface{}

data = "suvo"
data = 10
data = true
```

Go 1.18+ এ সাধারণত `any` ব্যবহার করা হয়:

```go
var data any
```

---

# Type Assertion

Interface-এর ভিতরের real value বের করতে:

```go
var data interface{} = "Hello"

value := data.(string)

fmt.Println(value)
```

---

# Type Switch

```go
func checkType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown")
	}
}
```

---

# Pointer Receiver with Interface

```go
type User struct {
	Name string
}

func (u *User) Print() {
	fmt.Println(u.Name)
}
```

এখন:

```go
var p Printer

u := User{Name: "Suvo"}

p = &u
```

কারণ method pointer receiver দিয়ে বানানো।

---

# Interface Composition

একাধিক interface combine করা যায়।

```go
type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}
```

---

# Real Industry Uses 🔥

Go interface use হয়:

- Database Driver
- Payment Gateway
- Logger System
- Authentication System
- Cloud SDK
- File Storage
- Microservices
- Queue System
- Cache Layer
- Notification System

---

# Example — Notification System

```go
type Notifier interface {
	Send(message string)
}
```

Implementations:

- Email
- SMS
- Push Notification
- WhatsApp

সব একই interface follow করবে।

---

# Interface vs Struct

| Struct             | Interface               |
| ------------------ | ----------------------- |
| Real object        | Behavior define করে     |
| Data রাখে          | Rules রাখে              |
| Memory allocate হয় | Contract হিসেবে কাজ করে |

---

# Important Thing ⚠️

Go-তে:

> Interface should be small

Good:

```go
type Reader interface {
	Read()
}
```

Bad:

```go
type MegaInterface interface {
	Read()
	Write()
	Delete()
	Update()
	Upload()
	Download()
}
```

---

# Famous Go Standard Library Interfaces

Go standard library-তেও interface heavily use হয়।

Example:

## io.Reader

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

এটা implement করে:

- File
- Network Connection
- HTTP Body
- Buffer

সব একইভাবে read করা যায়।

---

# Final Industry Mental Model 🧠

```text
Interface = Capability

Car can Start()
Bike can Start()

bKash can Pay()
Stripe can Pay()

Email can Send()
SMS can Send()
```

মানে:

> “What object CAN DO”
> না যে
> “What object IS”
