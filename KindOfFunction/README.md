# Golang Functional Programming Concepts (Interview Preparation)

---

# Function / Functional Paradigm কি?

Functional programming paradigm is a style where computation is done using functions without changing state and by treating functions as values.

# ১. Parameter vs Argument

অনেক ইন্টারভিউতেই প্রথমে এই প্রশ্ন করা হয়।

## Parameter কী?

ফাংশন ডিফাইন করার সময় যে ভেরিয়েবলগুলো ইনপুট হিসেবে লেখা হয়, সেগুলো Parameter।

```go
func Add(a int, b int) {
	fmt.Println(a + b)
}
```

এখানে:

```go
a int
b int
```

দুটোই Parameter।

---

## Argument কী?

ফাংশন কল করার সময় যে আসল ভ্যালুগুলো পাঠানো হয় সেগুলো Argument।

```go
Add(10, 20)
```

এখানে:

```go
10
20
```

দুটোই Argument।

---

## Visualization

```text
Function Definition
        ↓
func Add(a,b int)

      Parameter
        ↓

Function Call
        ↓
Add(10,20)

      Argument
        ↓
```

---

## Interview Question

### Q: Parameter এবং Argument এর মধ্যে পার্থক্য কী?

**Answer**

| Parameter                  | Argument             |
| -------------------------- | -------------------- |
| Function definition এ থাকে | Function call এ থাকে |
| Value receive করে          | Value send করে       |
| Placeholder                | Actual value         |

---

# ২. First-Order Function

যে Function অন্য Function কে Parameter হিসেবে নেয় না এবং Function Return করে না, তাকে First-Order Function বলা হয়।

## Example

```go
package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(10, 20)
	fmt.Println(result)
}
```

Output

```text
30
```

এখানে:

- Function Parameter হিসেবে int নিয়েছে
- Function Return করেছে int

তাই এটি First-Order Function।

---

## Interview Question

### Q: First-Order Function কী?

**Answer**

যে Function শুধুমাত্র সাধারণ Data Type নিয়ে কাজ করে এবং Function কে Parameter বা Return হিসেবে ব্যবহার করে না তাকে First-Order Function বলে।

---

# ৩. First-Class Function (First-Class Citizen)

Go-তে Function-কে Variable এর মতো ব্যবহার করা যায়।

এ কারণেই Go-তে Function হলো First-Class Citizen।

---

## Function Variable এ Store

```go
package main

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func main() {

	greet := SayHello

	greet()
}
```

Output

```text
Hello
```

---

## Function Variable Memory View

```text
SayHello()
     ↑
     │
greet variable
```

---

## Function Pass করা যায়

```go
func Print(fn func()) {
	fn()
}
```

---

## Function Return করা যায়

```go
func GetFunction() func() {

	return func() {
		fmt.Println("Hello")
	}
}
```

---

## Interview Question

### Q: First-Class Function বলতে কী বোঝায়?

**Answer**

যখন Function কে Variable এর মতো Assign, Pass এবং Return করা যায় তখন Function কে First-Class Function বলা হয়।

---

# ৪. Higher-Order Function

যে Function

- Function কে Parameter হিসেবে নেয়
- অথবা Function Return করে

তাকে Higher-Order Function বলে।

---

## Example 1: Function as Parameter

```go
package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Execute(fn func(int, int) int) {

	result := fn(10, 20)

	fmt.Println(result)
}

func main() {

	Execute(Add)
}
```

Output

```text
30
```

---

## Visualization

```text
Add()
  ↓
Execute(Add)

fn → Add
```

---

## Example 2: Function Return

```go
package main

import "fmt"

func Multiplier() func(int) int {

	return func(x int) int {
		return x * 2
	}
}

func main() {

	double := Multiplier()

	fmt.Println(double(10))
}
```

Output

```text
20
```

---

## Interview Question

### Q: Higher-Order Function কী?

**Answer**

যে Function অন্য Function কে Parameter হিসেবে নেয় অথবা Function Return করে তাকে Higher-Order Function বলে।

---

# ৫. Callback Function

যখন একটি Function কে অন্য Function এর মধ্যে Argument হিসেবে পাঠানো হয় তখন সেটি Callback Function।

---

## Example

```go
package main

import "fmt"

func Process(callback func()) {

	fmt.Println("Processing...")

	callback()
}

func Finished() {
	fmt.Println("Completed")
}

func main() {

	Process(Finished)
}
```

Output

```text
Processing...
Completed
```

---

## এখানে Callback কে?

```go
Finished
```

কারণ এটি:

```go
Process(Finished)
```

এর মাধ্যমে পাঠানো হয়েছে।

---

## Real Life Example

ধরো:

```text
Food Order
    ↓
Food Ready হলে
    ↓
Customer কে Call
```

এখানে Customer Notification Function হচ্ছে Callback।

---

## Interview Question

### Q: Callback Function কী?

**Answer**

যে Function অন্য Function এর Argument হিসেবে Pass করা হয় এবং পরে Execute করা হয় তাকে Callback Function বলে।

---

# Higher-Order Function vs Callback

অনেক Interview তে এই প্রশ্ন আসে।

```go
func Execute(fn func()) {
	fn()
}
```

এখানে:

```go
Execute
```

Higher-Order Function

কারণ Function Parameter নিয়েছে।

---

```go
Execute(SayHello)
```

এখানে:

```go
SayHello
```

Callback Function

কারণ এটি Argument হিসেবে Pass হয়েছে।

---

# Complete Example

```go
package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Calculator(
	fn func(int, int) int,
	a int,
	b int,
) {

	fmt.Println(fn(a, b))
}

func main() {

	Calculator(Add, 10, 20)
}
```

এখানে:

| Item          | Type                  |
| ------------- | --------------------- |
| Calculator    | Higher-Order Function |
| Add           | Callback Function     |
| Function Type | First-Class Function  |
| a,b           | Parameter             |
| 10,20         | Argument              |

---

# Most Important Interview Questions

### 1. Parameter এবং Argument এর মধ্যে পার্থক্য কী?

### 2. First-Class Function কী?

### 3. Go-তে Function কি First-Class Citizen?

**Answer:** Yes.

---

### 4. Higher-Order Function কী?

### 5. Callback Function কী?

### 6. Higher-Order Function এবং Callback Function এর মধ্যে পার্থক্য কী?

**Answer:**

| Higher-Order Function       | Callback Function    |
| --------------------------- | -------------------- |
| Function receive/return করে | Function pass করা হয় |
| Receiver                    | Sender               |

---

### 7. Function Variable এ Store করা যায়?

```go
greet := SayHello
```

**Answer:** Yes.

---

### 8. Function Return করা যায়?

```go
func GetFunc() func()
```

**Answer:** Yes.

---

### 9. First-Class Function না থাকলে Higher-Order Function সম্ভব?

**Answer:** না।

কারণ Higher-Order Function কাজ করার জন্য Function-কে Variable এর মতো Treat করতে হয়।

---

### 10. Go Functional Programming Support করে?

**Answer:**

হ্যাঁ, Go Partial Functional Programming Support করে।

Features:

- First-Class Functions
- Anonymous Functions
- Closures
- Higher-Order Functions
- Callback Functions

তবে Go পুরোপুরি Functional Language নয়, যেমন: Haskell বা Elixir। এটি মূলত একটি Procedural + Concurrent + Multi-Paradigm Language।
