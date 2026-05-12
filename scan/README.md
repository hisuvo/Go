Go এ `Scan` method/user input নেওয়ার জন্য `fmt` package ব্যবহার করা হয়।

সবচেয়ে common হলো:

- `fmt.Scan()`
- `fmt.Scanln()`
- `fmt.Scanf()`

---

# 1. fmt.Scan()

স্পেস পর্যন্ত input নেয়।

```go
package main

import "fmt"

func main() {
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name)
}
```

## Input

```txt
Suvo
```

## Output

```txt
Hello Suvo
```

---

# Important → `&` কেন?

```go
fmt.Scan(&name)
```

`&name` মানে variable এর memory address পাঠানো।

কারণ `Scan()` variable এর value change করবে।

---

# Multiple Input

```go
package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scan(&name, &age)

	fmt.Println(name, age)
}
```

## Input

```txt
Suvo 22
```

---

# 2. fmt.Scanln()

`Enter` press পর্যন্ত input নেয়।

```go
package main

import "fmt"

func main() {
	var first string
	var last string

	fmt.Scanln(&first, &last)

	fmt.Println(first, last)
}
```

---

# Difference Between Scan & Scanln

| Method     | Stops At       |
| ---------- | -------------- |
| `Scan()`   | space/new line |
| `Scanln()` | new line       |

---

# 3. fmt.Scanf()

Formatted input নেওয়ার জন্য।

```go
package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scanf("%s %d", &name, &age)

	fmt.Println(name, age)
}
```

---

# `%s` and `%d`

| Symbol | Meaning |
| ------ | ------- |
| `%s`   | string  |
| `%d`   | integer |
| `%f`   | float   |

---

# Real Example

```go
package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("Enter two numbers: ")

	fmt.Scan(&num1, &num2)

	fmt.Println("Sum =", num1+num2)
}
```

---

# Return Value

`Scan()` কতগুলো value successfully read করেছে সেটা return করে।

```go
n, err := fmt.Scan(&name)
```

- `n` → কয়টা input read হয়েছে
- `err` → error আছে কিনা

Example:

```go
package main

import "fmt"

func main() {
	var age int

	n, err := fmt.Scan(&age)

	fmt.Println(n)
	fmt.Println(err)
}
```

---

# Modern Better Way

Large text/string input এর জন্য অনেক developer `bufio.NewReader()` ব্যবহার করে।

কারণ `Scan()` full sentence এ problem করতে পারে।

Example:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	fmt.Println(text)
}
```
