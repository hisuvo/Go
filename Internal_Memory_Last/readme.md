# Golang Internal Memory Management (Interview + Deep Concept)

Go প্রোগ্রাম রান করার সময় Operating System RAM-এর মধ্যে কয়েকটি ভাগে মেমোরি সাজিয়ে রাখে।

```
+------------------+
| Code Segment     |
+------------------+
| Data Segment     |
+------------------+
| Stack            |
+------------------+
| Heap             |
+------------------+
```

---

# ১. Compilation Phase

ধরো তুমি এই কোড লিখেছ:

```go
package main

import "fmt"

const PI = 3.1416

func main() {
	name := "Suvo"
	fmt.Println(name)
}
```

### Compilation এর সময়

Go Compiler:

```bash
go build main.go
```

চালালে

```
main.go
   ↓
Compiler
   ↓
main.exe / binary
```

তৈরি হয়।

Compiler তখন:

- Syntax Check করে
- Type Check করে
- Machine Code তৈরি করে
- Constant optimize করে
- Binary File তৈরি করে

এখনও RAM ব্যবহার হচ্ছে না।

---

# ২. Execution Phase

যখন রান করবে:

```bash
go run main.go
```

অথবা

```bash
./main.exe
```

তখন OS:

1. Binary File RAM এ Load করে
2. Memory Segment তৈরি করে
3. main() Function Execute করে

---

# ৩. Code Segment

এখানে থাকে:

- Machine Instructions
- Function Code
- Program Logic

উদাহরণ:

```go
func add(a int, b int) int {
	return a + b
}
```

Compiler এটাকে Machine Instruction এ পরিণত করে।

সেগুলো Code Segment এ থাকে।

---

### বৈশিষ্ট্য

✅ Read Only

❌ Modify করা যায় না

কারণ:

```go
func add() {}
```

Runtime এ গিয়ে Function এর Code পরিবর্তন হলে Program Crash করতে পারে।

---

# ৪. Data Segment

এখানে থাকে:

- Global Variable
- Static Data
- Constant

উদাহরণ:

```go
package main

var age = 25

const PI = 3.1416
```

RAM এ:

```
Data Segment

PI  = 3.1416
age = 25
```

---

# Constant কেন Change করা যায় না?

```go
const PI = 3.14

PI = 5
```

Error:

```go
cannot assign to PI
```

কারণ:

Compiler জানে এটি Constant।

এটি Read Only Data হিসেবে Store হয়।

---

# ৫. Stack Memory

সবচেয়ে গুরুত্বপূর্ণ Interview Topic।

---

## Stack কী?

Function Call এর Local Data রাখে।

উদাহরণ:

```go
func main() {
	x := 10
	y := 20
}
```

Stack:

```
main frame

x = 10
y = 20
```

---

## Function Call হলে কী হয়?

```go
func sum(a int, b int) int {
	total := a + b
	return total
}
```

```go
func main() {
	result := sum(10, 20)
}
```

Stack:

```
main frame

result

----------------

sum frame

a = 10
b = 20
total = 30
```

---

Function শেষ:

```go
return total
```

হলে

```
sum frame
```

পুরো Delete হয়ে যায়।

---

### Stack এর সুবিধা

খুব Fast

কারণ:

```text
Push
Pop
```

Operation হয়।

---

# Stack Frame কী?

প্রতিটি Function Call এর জন্য Stack এ আলাদা Box তৈরি হয়।

```
main()

┌──────────┐
│ x        │
└──────────┘

call sum()

┌──────────┐
│ total    │
│ a        │
│ b        │
└──────────┘
```

এটাকে Stack Frame বলে।

---

# ৬. Heap Memory

Heap হলো Dynamic Memory Area।

যখন Data Function শেষ হওয়ার পরও বেঁচে থাকতে হবে তখন Heap ব্যবহার হয়।

---

উদাহরণ:

```go
func createUser() *string {
	name := "Suvo"
	return &name
}
```

এখানে:

```go
return &name
```

Local Variable এর Address Return হচ্ছে।

Function শেষ হলে Stack Destroy হয়ে যাবে।

তাই Go Compiler Escape Analysis করে।

---

Compiler বলে:

```text
name stack এ রাখা নিরাপদ নয়।
```

তাই

```
name → Heap
```

এ চলে যায়।

---

# Escape Analysis

Go Compiler সিদ্ধান্ত নেয়:

```text
Stack?
নাকি
Heap?
```

---

উদাহরণ ১

```go
func main() {
	x := 10
	fmt.Println(x)
}
```

Output:

```
Stack
```

---

উদাহরণ ২

```go
func getValue() *int {
	x := 10
	return &x
}
```

Output:

```
Heap
```

কারণ Address বাইরে যাচ্ছে।

---

Compiler Check করতে:

```bash
go build -gcflags="-m"
```

---

Output:

```text
moved to heap: x
```

---

# ৭. Function Expression Memory

ধরো:

```go
func main() {

	add := func(a int, b int) int {
		return a + b
	}

	fmt.Println(add(10,20))
}
```

অনেকে ভাবে পুরো Function Stack এ যায়।

আসলে:

Function Code →

```
Code Segment
```

এ থাকে।

---

Variable:

```go
add
```

শুধু Reference রাখে।

```
Stack

add -----> Function Code
```

---

# Closure Example

```go
func counter() func() int {

	count := 0

	return func() int {
		count++
		return count
	}
}
```

---

এখানে:

```go
count
```

Function শেষ হওয়ার পরও বেঁচে থাকে।

তাই

```
Heap
```

এ চলে যায়।

---

Memory:

```
Heap
┌─────────┐
│ count=0 │
└─────────┘

Closure
   ↓
count
```

---

# Garbage Collector (GC)

Heap Memory Clean করে।

উদাহরণ:

```go
user := &User{}
```

যদি পরে:

```go
user = nil
```

এবং আর কোথাও Reference না থাকে

তাহলে GC Memory Free করে দেয়।

---

# Interview Questions

### Q1: Code Segment কী?

**Answer:**
Program এর Machine Instructions এবং Function Code যেখানে থাকে তাকে Code Segment বলে।

---

### Q2: Stack কী?

**Answer:**
Function Call, Local Variable এবং Stack Frame সংরক্ষণের Memory Area।

---

### Q3: Heap কী?

**Answer:**
Dynamic Memory Allocation এর জন্য ব্যবহৃত Memory Area যেখানে Data Function Lifetime ছাড়িয়ে বেঁচে থাকতে পারে।

---

### Q4: Stack কেন Fast?

**Answer:**
Push এবং Pop Operation ব্যবহার করে Sequential Memory Access হয়।

---

### Q5: Escape Analysis কী?

**Answer:**
Go Compiler এর একটি Optimization Technique যা নির্ধারণ করে Variable Stack এ থাকবে নাকি Heap এ যাবে।

---

### Q6: Local Variable কি Heap এ যেতে পারে?

**Answer:**
হ্যাঁ।

```go
func getValue() *int {
	x := 10
	return &x
}
```

এখানে `x` Heap এ যাবে।

---

### Q7: Closure কী?

**Answer:**
যে Function তার Outer Scope এর Variable Access করতে পারে তাকে Closure বলে।

---

### Q8: Function Expression কোথায় থাকে?

**Answer:**
Function এর Code থাকে Code Segment এ, আর Function Variable এর Reference Stack/Heap এ থাকতে পারে।

---

### Q9: Garbage Collector কী?

**Answer:**
Unused Heap Memory Automatically Free করার System।

---

### Q10: Go তে Memory Leak হতে পারে?

**Answer:**
হ্যাঁ, যদি Unnecessary Reference ধরে রাখা হয় তাহলে Heap Memory Release হবে না এবং Memory Leak এর মতো সমস্যা তৈরি হতে পারে।

---

একজন Backend Go Developer হিসেবে Stack, Heap, Escape Analysis, Closure, Pointer এবং Garbage Collector এই ৫টি টপিক খুব ভালোভাবে বুঝতে পারলে Go-এর ইন্টারনাল মেমোরি ম্যানেজমেন্টের প্রায় ৮০% ধারণা পরিষ্কার হয়ে যাবে।

---

খুব ভালো প্রশ্ন SUVO। অনেকেই এখানে কনফিউজড হয়।

**Code Segment Read Only** আর **Constant Read Only** — দুটো একই জিনিস না।

---

## Code Segment কেন Read Only?

Code Segment এ থাকে:

```go
func add(a, b int) int {
	return a + b
}
```

কম্পাইল হওয়ার পর Machine Instructions:

```text
MOV
ADD
RET
```

এই Instructions গুলো Runtime এ পরিবর্তন করা যাবে না।

তাই Code Segment = Read Only।

---

## Constant কোথায় থাকে?

এটা একটু জটিল।

Go-তে Constant সবসময় Data Segment এ থাকবে এমন না।

Compiler Optimization এর উপর নির্ভর করে।

উদাহরণ:

```go
const PI = 3.14

func main() {
	fmt.Println(PI)
}
```

Compiler অনেক সময় `PI`-কে Binary এর মধ্যে সরাসরি বসিয়ে দেয়।

```text
fmt.Println(3.14)
```

তখন আলাদা Memory Location-ও নাও থাকতে পারে।

---

## তাহলে Data Segment এ Constant কেন বলা হয়?

Operating System Memory Layout শেখানোর সময় সাধারণভাবে বলা হয়:

```text
Code Segment
Data Segment
Stack
Heap
```

এবং

```text
Global Variables
Constants
```

Data Segment এর Read-Only অংশে রাখা হতে পারে।

বাস্তবে Modern Compiler:

- Constant inline করতে পারে
- Register এ রাখতে পারে
- Read-only data section (.rodata) এ রাখতে পারে

---

## Data Segment এর ভিতরও ভাগ আছে

সাধারণত:

```text
Data Segment
│
├── Initialized Data
├── Uninitialized Data (.bss)
└── Read Only Data (.rodata)
```

উদাহরণ:

```go
var age = 25
```

যাবে:

```text
Initialized Data
```

আর

```go
const PI = 3.14
```

যেতে পারে:

```text
.rodata
```

যা Read Only।

---

## Interview Answer

**Question:** Constant যদি Read Only হয় তাহলে Code Segment এ কেন থাকে না?

**Answer:**

Code Segment এ Program Instructions (Machine Code) থাকে। Constant কোনো Instruction নয়, এটি Data। তাই Constant সাধারণত Read-Only Data Section (`.rodata`) এ রাখা হয়, যা Data Segment-এর অংশ। Compiler Optimization এর কারণে অনেক Constant সরাসরি Code এর মধ্যে Inline-ও করে দিতে পারে।

---

সহজভাবে মনে রাখো:

```text
Code Segment
= কী কাজ করবে (Instructions)

Data Segment
= কী নিয়ে কাজ করবে (Data)

Stack
= Local Variables

Heap
= Dynamic Data
```

`const PI = 3.14` হলো **Data**, Instruction না। তাই এটি Code Segment এ নয়, Read-Only Data Area (.rodata) তে থাকে।
