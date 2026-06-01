## Variable Shadowing কী?

যখন কোনো inner scope (block, function, if, for, switch ইত্যাদি) এর মধ্যে একই নামের নতুন variable declare করা হয়, তখন outer scope-এর variable সাময়িকভাবে hidden হয়ে যায়। এটাকেই **Variable Shadowing** বলে।

তোমার কোড:

```go
package main

import "fmt"

var a = 10

func main() {
	age := 30

	if age > 120 {
		a := 47
		fmt.Println("Condition is true ->", a)
	} else {
		a := 40
		fmt.Println("Condition is false ->", a)
	}

	fmt.Println("Condition outSide ->", a)
}
```

### Memory View

Global Scope

```go
a = 10
```

If Block

```go
a := 47
```

Else Block

```go
a := 40
```

এখানে `:=` নতুন variable তৈরি করে।

---

## Output

```text
Condition is false -> 40
Condition outSide -> 10
```

### কেন 10 প্রিন্ট হলো?

কারণ:

```go
a := 40
```

এখানে global `a` update হয়নি।

বরং নতুন local `a` তৈরি হয়েছে।

Block শেষ হলে local `a` destroy হয়ে যায়।

তাই বাইরে আবার global `a = 10` দেখা যায়।

---

# Shadowing Visualization

```go
var a = 10

func main() {

	a := 20

	fmt.Println(a)
}
```

Visualization:

```text
Global Scope
a = 10

Main Scope
a = 20
```

Output:

```text
20
```

Main function-এর ভিতরে global `a` shadow হয়ে গেছে।

---

# Shadowing vs Assignment

## Shadowing

```go
a := 10

{
	a := 20
	fmt.Println(a)
}

fmt.Println(a)
```

Output

```text
20
10
```

নতুন variable তৈরি হয়েছে।

---

## Assignment

```go
a := 10

{
	a = 20
	fmt.Println(a)
}

fmt.Println(a)
```

Output

```text
20
20
```

নতুন variable তৈরি হয়নি।

আগের variable-এর value change হয়েছে।

---

# Dangerous Example

```go
func main() {

	x := 10

	if true {
		x := 20
		fmt.Println(x)
	}

	fmt.Println(x)
}
```

Output

```text
20
10
```

অনেক Go developer ভুল করে মনে করে:

```go
x := 20
```

এটা x update করছে।

আসলে নতুন x তৈরি করছে।

---

# Error Handling Shadowing

Interview-এ খুব common।

### Bad Example

```go
file, err := os.Open("data.txt")

if err != nil {
	return
}

data, err := io.ReadAll(file)
```

এখানে সমস্যা নেই।

কিন্তু:

```go
err := doSomething()

if err != nil {

	err := doAnotherThing()

	fmt.Println(err)
}

fmt.Println(err)
```

এখানে ভিতরের `err` বাইরের `err`-কে shadow করেছে।

---

# For Loop Shadowing

```go
x := 100

for i := 0; i < 3; i++ {
	x := i
	fmt.Println(x)
}

fmt.Println(x)
```

Output

```text
0
1
2
100
```

---

# Function Parameter Shadowing

```go
var name = "Global"

func printName(name string) {
	fmt.Println(name)
}
```

Function parameter `name` global `name`-কে shadow করেছে।

---

# Named Return Shadowing

```go
func test() (err error) {

	if true {
		err := fmt.Errorf("something wrong")
		fmt.Println(err)
	}

	return
}
```

এখানে named return variable `err` shadow হয়ে গেছে।

অনেক bug এই কারণে হয়।

---

# How to Avoid Shadowing

### Wrong

```go
count := 10

if true {
	count := 20
}
```

### Correct

```go
count := 10

if true {
	count = 20
}
```

`:=` এর বদলে `=` ব্যবহার করো।

---

# Interview Questions

### 1. Variable Shadowing কী?

**Answer:**
Inner scope-এ একই নামের নতুন variable declare করলে outer variable hidden হয়ে যায়। এটাকে Variable Shadowing বলে।

---

### 2. Output কী হবে?

```go
x := 10

{
	x := 20
	fmt.Println(x)
}

fmt.Println(x)
```

Output:

```text
20
10
```

---

### 3. `:=` এবং `=` এর পার্থক্য কী?

| Operator | Meaning                       |
| -------- | ----------------------------- |
| `:=`     | New variable declare + assign |
| `=`      | Existing variable update      |

---

### 4. Shadowing কেন dangerous?

কারণ developer মনে করতে পারে variable update হচ্ছে, কিন্তু আসলে নতুন variable তৈরি হচ্ছে।

ফলে unexpected output বা bug হয়।

---

### 5. Global Variable Shadowing Example?

```go
var count = 10

func main() {
	count := 20
	fmt.Println(count)
}
```

Output:

```text
20
```

Global `count` shadow হয়েছে।

---

### 6. তোমার কোডের শেষ লাইনে `10` কেন প্রিন্ট হবে?

কারণ:

```go
a := 40
```

এবং

```go
a := 47
```

দুটিই local variable তৈরি করেছে।

Global `a = 10` পরিবর্তন হয়নি।

তাই:

```go
fmt.Println("Condition outSide ->", a)
```

প্রিন্ট করবে:

```text
Condition outSide -> 10
```

Go interview-এ shadowing নিয়ে সবচেয়ে বেশি জিজ্ঞেস করা হয়:

- `:=` vs `=`
- `err` shadowing
- loop variable shadowing
- global variable shadowing
- named return shadowing
- scope rules (block scope, package scope, function scope)
