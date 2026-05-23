Go তে custom error তৈরি করার সবচেয়ে common way হলো `errors.New()` অথবা `fmt.Errorf()` ব্যবহার করা।
আর advanced ভাবে `struct + Error() method` ব্যবহার করা হয়।

---

# 1. Simple Custom Error

```go
package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {
	if age < 18 {
		return errors.New("age must be 18 or older")
	}

	return nil
}

func main() {
	err := checkAge(15)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Valid age")
}
```

### Output

```bash
Error: age must be 18 or older
```

---

# 2. Dynamic Custom Error (`fmt.Errorf`)

Dynamic message বানাতে:

```go
package main

import (
	"fmt"
)

func divide(a, b int) error {
	if b == 0 {
		return fmt.Errorf("cannot divide %d by zero", a)
	}

	return nil
}

func main() {
	err := divide(10, 0)

	if err != nil {
		fmt.Println(err)
	}
}
```

---

# 3. Professional Custom Error (Struct দিয়ে)

Real backend project এ এটা বেশি use হয়।

```go
package main

import (
	"fmt"
)

type ValidationError struct {
	Field string
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return ValidationError{
			Field: "password",
			Message: "must be at least 8 characters",
		}
	}

	return nil
}

func main() {
	err := validatePassword("123")

	if err != nil {
		fmt.Println(err)
	}
}
```

### Output

```bash
password: must be at least 8 characters
```

---

# কেন Custom Error ব্যবহার হয়?

- Better debugging
- Clean backend architecture
- API response handling
- Different error type detect করতে
- Large project maintain করতে

---

# Real Backend Example

```go
type NotFoundError struct {
	Resource string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", n.Resource)
}
```

Use:

```go
return NotFoundError{
	Resource: "User",
}
```

Output:

```bash
User not found
```

---

# Interview Question

## Go তে error কি?

Go তে `error` একটি built-in interface।

```go
type error interface {
	Error() string
}
```

---

## কেন struct দিয়ে custom error বানানো হয়?

কারণ extra data রাখা যায়।

Example:

- Status code
- Field name
- Error type
- Request ID

---

## `errors.New()` vs `fmt.Errorf()`

| Feature        | errors.New | fmt.Errorf |
| -------------- | ---------- | ---------- |
| Static message | ✅         | ✅         |
| Dynamic value  | ❌         | ✅         |
| Formatting     | ❌         | ✅         |

---

# Important Concept

যে কোনো struct যদি:

```go
Error() string
```

method implement করে, তাহলে সেটা automatically `error` হয়ে যায়।

```go
type MyError struct {}

func (m MyError) Error() string {
	return "custom error"
}
```

এখন `MyError` একটি valid error type।
