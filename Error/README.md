In Go, handling errors isn't about catching exceptions; it’s about treating errors as values that your code inspects, wraps, and passes around. When standard errors aren't enough, **custom errors** and **type assertions** allow you to build robust, production-grade error handling.

Here is an in-depth breakdown of how they work, how they pair together, and the exact concepts you need to know to ace your interview.

---

## 1. Custom Errors in Go

At its core, the built-in `error` type is a simple single-method interface:

```go
type error interface {
    Error() string
}

```

Any type that implements the `Error() string` method satisfies this interface. While `errors.New()` and `fmt.Errorf()` are great for simple text errors, you need **custom error structs** when you want to attach extra context (like HTTP status codes, error codes, or timestamps) that your code can programmatically act upon.

### How to Define and Implement a Custom Error

```go
package main

import (
	"fmt"
	"time"
)

// QueryError is a custom error struct carrying rich context
type QueryError struct {
	Query     string
	ErrCode   int
	Timestamp time.Time
	Err       error // Underlying error
}

// Implementing the error interface
func (e *QueryError) Error() string {
	return fmt.Sprintf("code %d failed at %s for query %q: %v", e.ErrCode, e.Timestamp.Format(time.RFC3339), e.Query, e.Err)
}

```

---

## 2. Type Assertion & Error Inspections

Because custom errors are passed up the call stack as the generic `error` interface, you lose direct access to the underlying struct fields. To get those fields back, you must extract the concrete type from the interface.

In Go, this is done using **Type Assertion** or the modern `errors.As()` function.

### Method A: The Classic Type Assertion

Type assertion uses the syntax `interfaceVariable.(ConcreteType)`. It extracts the concrete value stored inside an interface.

```go
func executeQuery(q string) error {
	// Simulating an error return
	return &QueryError{
		Query:     q,
		ErrCode:   5001,
		Timestamp: time.Now(),
		Err:       fmt.Errorf("connection timeout"),
	}
}

func main() {
	err := executeQuery("SELECT * FROM users")
	if err != nil {
		// Type assertion: extracting the *QueryError pointer
		if qErr, ok := err.(*QueryError); ok {
			fmt.Printf("Programmatic handling: Found Error Code %d\n", qErr.ErrCode)
			fmt.Printf("Failed Query: %s\n", qErr.Query)
		} else {
			fmt.Println("Generic error encountered:", err)
		}
	}
}

```

### Method B: The Modern Way (`errors.As`)

If your error gets wrapped (e.g., using `fmt.Errorf("context: %w", err)`), classic type assertion will **fail** because the outer error is no longer of type `*QueryError`.

To fix this, Go 1.13 introduced `errors.As()`. It unwraps the error chain recursively until it finds a match. **In production and interviews, always favor `errors.As`.**

```go
import "errors"

func handle() {
	err := executeQuery("SELECT * FROM users")
	// Wrapped error
	wrappedErr := fmt.Errorf("database layer: %w", err)

	var qErr *QueryError
	// errors.As checks the entire chain and populates qErr if matched
	if errors.As(wrappedErr, &qErr) {
		fmt.Printf("Successfully unpacked wrapped error! Code: %d\n", qErr.ErrCode)
	}
}

```

---

## Must-Know Interview Questions & Answers

### Q1: What is the difference between `errors.Is()` and `errors.As()`?

**Answer:**

- `errors.Is()` is used to compare an error against a **specific value** (like a sentinel error like `sql.ErrNoRows` or `io.EOF`). It behaves like `if err == target`.
- `errors.As()` is used to check if an error matches a **specific type** (like our `*QueryError` struct). It behaves like a type assertion and populates a target variable with the concrete error value if a match is found in the error chain.
- Both functions are superior to direct `==` or type assertions because they automatically unwrap errors that have been bundled via `%w`.

### Q2: Why do we almost always implement the `Error()` method on a pointer receiver (`func (e *MyError) Error()`) rather than a value receiver?

**Answer:**
There are two critical reasons:

1. **Consistency:** It allows us to modify internal state if needed, though errors are usually immutable.
2. **Avoiding `nil` interface traps:** If you assign a `nil` pointer of a custom type to an `error` interface, the interface itself is **not nil** (it contains a concrete type descriptor but a `nil` value). If you use a pointer receiver, it makes it distinct and easier to explicitly type-assert against `(*MyError)(nil)`.

### Q3: Explain this common bug: "An error is not nil even when the returned custom pointer is nil."

**Answer:**
An interface value in Go consists of two components under the hood: a **concrete type** and a **value**. An interface is only truly `nil` if both components are empty.

If a function returns a custom error pointer that happens to be `nil`, but the function return signature is the generic `error` interface, Go wraps that `nil` pointer inside the interface structure.

```go
func dummy() error {
    var p *QueryError = nil
    return p // Returns an interface containing (Type: *QueryError, Value: nil)
}

func main() {
    err := dummy()
    if err != nil {
        fmt.Println("This WILL print!") // Bug! err is not nil because the type descriptor exists.
    }
}

```

**Fix:** Always explicitly return `nil` instead of a typed pointer when no error occurs.

---

## Production Blueprint

Here is a practical, interview-ready pattern showing custom errors, dynamic context, wrapping, and extraction:

```go
package main

import (
	"errors"
	"fmt"
)

// Domain-specific error codes
const (
	ErrNotFound = 404
	ErrTimeout  = 408
)

type AppError struct {
	Code    int
	Message string
	Op      string // Operation name (e.g., "UserService.Fetch")
	Err     error  // Underlying root cause
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] code=%d message=%s: %v", e.Op, e.Code, e.Message, e.Err)
}

func (e *AppError) Unwrap() error {
	return e.Err // Allows errors.Is and errors.As to drill down
}

// Simulating a database fetch
func DBGetRecord() error {
	return fmt.Errorf("postgres driver connection dropped") // Root cause
}

// Application Layer wrapping the error
func FetchUser() error {
	err := DBGetRecord()
	if err != nil {
		return &AppError{
			Code:    ErrTimeout,
			Message: "unable to load user profile",
			Op:      "UserService.FetchUser",
			Err:     err,
		}
	}
	return nil
}

func main() {
	err := FetchUser()
	if err != nil {
		// Extracting custom attributes at the top level (e.g., API Gateway)
		var appErr *AppError
		if errors.As(err, &appErr) {
			fmt.Printf("HTTP Status to return: %d\n", appErr.Code)
			fmt.Printf("Logging internal op details: %s\n", appErr.Op)
			fmt.Printf("Root Cause: %v\n", errors.Unwrap(err))
		}
	}
}

```

---

গো-ল্যাং-এ (Go language) এরর হ্যান্ডলিং কোনো `try-catch` বা এক্সেপশন (Exception) দিয়ে হয় না। গো-তে এররকে সাধারণ ডেটা বা ভ্যালু (Value) হিসেবে দেখা হয়। যখন সাধারণ এরর দিয়ে কাজ চলে না, তখন বড় বা প্রোডাকশন লেভেলের অ্যাপ্লিকেশনে **কাস্টম এরর (Custom Errors)** এবং **টাইপ অ্যাসারশন (Type Assertions)** ব্যবহার করতে হয়।

নিচে ইন্টারভিউ ক্র্যাক করার মতো গভীর আলোচনা এবং গুরুত্বপূর্ণ প্রশ্নগুলো বাংলায় সহজভাবে বুঝিয়ে দেওয়া হলো।

---

## ১. কাস্টম এরর (Custom Errors)

গো-তে বিল্ট-ইন `error` টাইপটি আসলে একটি ইন্টারফেস (Interface), যার ভেতরে কেবল একটি মেথড থাকে:

```go
type error interface {
    Error() string
}

```

যে কোনো স্ট্রাক্ট (Struct) বা টাইপ যদি এই `Error() string` মেথডটি ইমপ্লিমেন্ট করে, তবেই সেটি একটি এরর হিসেবে গণ্য হয়। যখন আমাদের এররের সাথে অতিরিক্ত তথ্য (যেমন: HTTP Status Code, Error Code, বা Timestamps) ট্র্যাক করতে হয়, তখন আমরা কাস্টম এরর স্ট্রাক্ট তৈরি করি।

### কাস্টম এরর তৈরি করার নিয়ম:

```go
package main

import (
	"fmt"
	"time"
)

// QueryError একটি কাস্টম এরর স্ট্রাক্ট যা অতিরিক্ত তথ্য ধারণ করে
type QueryError struct {
	Query     string
	ErrCode   int
	Timestamp time.Time
	Err       error // আসল বা মূল এরর (Underlying error)
}

// error ইন্টারফেস ইমপ্লিমেন্ট করা হচ্ছে
func (e *QueryError) Error() string {
	return fmt.Sprintf("কোড %d ব্যর্থ হয়েছে %s সময়ে, কুয়েরি: %q: %v", e.ErrCode, e.Timestamp.Format(time.RFC3339), e.Query, e.Err)
}

```

---

## ২. টাইপ অ্যাসারশন এবং এরর ইন্সপেকশন (Type Assertion)

ফাংশন থেকে যখন এরর রিটার্ন করা হয়, তখন সেটি জেনেরিক `error` ইন্টারফেস হিসেবে ওপরে আসে। ফলে কাস্টম স্ট্রাক্টের ভেতরের নির্দিষ্ট ফিল্ডগুলো (যেমন `ErrCode`) সরাসরি অ্যাক্সেস করা যায় না। ইন্টারফেসের ভেতর থেকে আসল কাস্টম টাইপটিকে বের করে আনার প্রক্রিয়াকেই **টাইপ অ্যাসারশন** বলে।

### পদ্ধতি ক: ক্লাসিক টাইপ অ্যাসারশন (Classic Type Assertion)

এটি `interfaceVariable.(ConcreteType)` সিনট্যাক্স ব্যবহার করে কাজ করে।

```go
func executeQuery(q string) error {
	return &QueryError{
		Query:     q,
		ErrCode:   5001,
		Timestamp: time.Now(),
		Err:       fmt.Errorf("connection timeout"),
	}
}

func main() {
	err := executeQuery("SELECT * FROM users")
	if err != nil {
		// টাইপ অ্যাসারশন: চেক করা হচ্ছে এররটি *QueryError কিনা
		if qErr, ok := err.(*QueryError); ok {
			fmt.Println("সফলভাবে টাইপ অ্যাসারশন হয়েছে!")
			fmt.Printf("এরর কোড: %d\n", qErr.ErrCode)
			fmt.Printf("ব্যর্থ কুয়েরি: %s\n", qErr.Query)
		} else {
			fmt.Println("এটি একটি সাধারণ এরর:", err)
		}
	}
}

```

### পদ্ধতি খ: আধুনিক নিয়ম (`errors.As`)

ইন্টারভিউ এবং প্রোডাকশন কোডের জন্য এটি **সবচেয়ে গুরুত্বপূর্ণ**। আপনার এররটি যদি অন্য কোনো কনটেক্সট দিয়ে র‍্যাপ (Wrap) করা থাকে (যেমন: `fmt.Errorf("ভুল হয়েছে: %w", err)`), তবে ওপরের ক্লাসিক টাইপ অ্যাসারশন **ফেইল (Fail) করবে**।

এই সমস্যার সমাধানের জন্য Go 1.13 সংস্করণে `errors.As()` আনা হয়েছে। এটি পুরো এরর চেইনটি খুঁজে ভেতর থেকে আসল কাস্টম এররটিকে বের করে আনে।

```go
import "errors"

func handle() {
	err := executeQuery("SELECT * FROM users")
	// এরর র‍্যাপ করা হলো
	wrappedErr := fmt.Errorf("ডাটাবেজ লেয়ার এরর: %w", err)

	var qErr *QueryError
	// errors.As পুরো চেইন চেক করে qErr-এর ভেতর ডেটা পোপুলেট করবে
	if errors.As(wrappedErr, &qErr) {
		fmt.Printf("র‍্যাপ করা এরর থেকে কোড পাওয়া গেছে: %d\n", qErr.ErrCode)
	}
}

```

---

## ইন্টারভিউতে আসার মতো গুরুত্বপূর্ণ প্রশ্ন ও উত্তর (Crucial Interview Q&A)

### প্রশ্ন ১: `errors.Is()` এবং `errors.As()` এর মধ্যে পার্থক্য কী?

**উত্তর:**

- **`errors.Is()`:** এটি ব্যবহার করা হয় কোনো এরর একটি **নির্দিষ্ট ভ্যালু (Value)** এর সমান কিনা তা পরীক্ষা করতে (যেমন sentinel error: `sql.ErrNoRows` বা `io.EOF`)। এটি মূলত `if err == target` এর মতো কাজ করে, তবে র‍্যাপ করা এররকেও আনর‍্যাপ করে চেক করতে পারে।
- **`errors.As()`:** এটি ব্যবহার করা হয় কোনো এরর একটি **নির্দিষ্ট টাইপ (Type)** এর কিনা তা পরীক্ষা করতে (যেমন আমাদের তৈরি করা `*QueryError` স্ট্রাক্ট)। এটি টাইপ অ্যাসারশনের মতো কাজ করে এবং টাইপ মিলে গেলে সেই কাস্টম এররের ডেটা টার্গেট ভ্যারিয়েবলে অ্যাসাইন করে।

### প্রশ্ন ২: কাস্টম এররের মেথড লেখার সময় আমরা কেন ভ্যালু রিসিভারের চেয়ে পয়েন্টার রিসিভার (`func (e *MyError) Error()`) বেশি ব্যবহার করি?

**উত্তর:**
এর প্রধান কারণ ২টি:
১. **Consistency (সঙ্গতি):** গো-তে কাস্টম এররগুলোকে সাধারণত পয়েন্টার হিসেবে পাস করা হয়। তাই মেথডটি পয়েন্টার রিসিভারে রাখলে সব জায়গায় সামঞ্জস্য থাকে।
২. **Nil Interface Trap এড়ানো:** যদি কোনো কারণে মেথডটি ভ্যালু রিসিভারে থাকে এবং আমরা ভুল করে একটি `nil` পয়েন্টার ইন্টারফেসে অ্যাসাইন করি, তবে ইন্টারফেসটি রানটাইমে প্যানিক (Panic) করতে পারে। পয়েন্টার রিসিভার ব্যবহার করলে সেফটি বজায় রাখা সহজ হয়।

### প্রশ্ন ৩: এই সাধারণ বাগটি (Bug) বুঝিয়ে বলুন: "ফাংশন থেকে রিটার্ন করা কাস্টম পয়েন্টারটি `nil` হওয়া সত্ত্বেও এরর ইন্টারফেসটি `nil` হয় না।" (Very Important)

**উত্তর:**
গো-তে একটি ইন্টারফেসের ব্যাকএন্ডে দুটি জিনিস থাকে: **টাইপ (TypeDescriptor)** এবং **ভ্যালু (Value)**। একটি ইন্টারফেস তখনই পুরোপুরি `nil` হয়, যখন তার টাইপ এবং ভ্যালু দুইটাই খালি থাকে।

যদি কোনো ফাংশন একটি কাস্টম এরর পয়েন্টার রিটার্ন করে যা বর্তমানে `nil` (যেমন: `var p *QueryError = nil`), এবং ফাংশনের রিটার্ন টাইপ যদি থাকে জেনেরিক `error` ইন্টারফেস, তবে গো ইন্টারফেসের টাইপের জায়গায় `*QueryError` বসিয়ে দেয় আর ভ্যালুর জায়গায় `nil` বসায়। যেহেতু টাইপের ঘর খালি নয়, তাই ইন্টারফেসটি `err != nil` চেকে পাস হয়ে যায় (অর্থাৎ গো মনে করে এরর আছে!)।

```go
func dummy() error {
    var p *QueryError = nil
    return p // ইন্টারফেসে জমা হবে: (Type: *QueryError, Value: nil)
}

func main() {
    err := dummy()
    if err != nil {
        fmt.Println("এই লাইনটি প্রিন্ট হবে! (যা একটি বাগ)")
    }
}

```

**সমাধান:** কোনো এরর না থাকলে সবসময় সরাসরি `return nil` লিখতে হবে, কোনো কাস্টম টাইপের `nil` ভ্যারিয়েবল রিটার্ন করা যাবে না।

---

## প্রোডাকশন লেভেল ব্লুপ্রিন্ট কোড (Production Blueprint)

নিচে একটি কমপ্লিট কোড দেওয়া হলো যা ইন্টারভিউতে বোর্ডে লিখে দিলে ইন্টারভিউয়ার বুঝবেন যে আপনি রিয়েল-ওয়ার্ল্ড আর্কিটেকচার জানেন:

```go
package main

import (
	"errors"
	"fmt"
)

// কাস্টম এরর কোড
const (
	ErrNotFound = 404
	ErrTimeout  = 408
)

// AppError আমাদের মেইন কাস্টম এরর স্ট্রাক্ট
type AppError struct {
	Code    int
	Message string
	Op      string // কোন ফাংশনে এরর হয়েছে (Operation)
	Err     error  // মূল এরর
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] কোড=%d মেসেজ=%s: %v", e.Op, e.Code, e.Message, e.Err)
}

// Unwrap মেথডটি থাকায় errors.Is এবং errors.As এররের ভেতরে প্রবেশ করতে পারে
func (e *AppError) Unwrap() error {
	return e.Err
}

// ডাটাবেজ লেয়ার (নকল ফাংশন)
func DBGetRecord() error {
	return fmt.Errorf("postgres driver connection dropped") // আসল এরর
}

// অ্যাপ্লিকেশন লেয়ার
func FetchUser() error {
	err := DBGetRecord()
	if err != nil {
		return &AppError{
			Code:    ErrTimeout,
			Message: "ইউজার প্রোফাইল লোড করা যায়নি",
			Op:      "UserService.FetchUser",
			Err:     err,
		}
	}
	return nil
}

func main() {
	err := FetchUser()
	if err != nil {
		var appErr *AppError
		// errors.As দিয়ে কাস্টম এররের ডেটা এক্সট্রাক্ট করা হচ্ছে
		if errors.As(err, &appErr) {
			fmt.Printf("ক্লায়েন্টকে পাঠানোর জন্য HTTP Status: %d\n", appErr.Code)
			fmt.Printf("লগিং ডিটেইলস (Operation): %s\n", appErr.Op)
			fmt.Printf("আসল মূল কারণ (Root Cause): %v\n", errors.Unwrap(err))
		}
	}
}

```
