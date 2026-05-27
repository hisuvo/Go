# Go Backend Parameters Handling (Proper Guide)

Go backend এ request data সাধারণত ৩ভাবে আসে:

1. Query Parameters → `?search=go`
2. Path Variables → `/users/12`
3. Request Body (JSON) → `{ "name": "suvo" }`

---

# 1. Query Parameters

👉 URL এর `?` এর পরে থাকে

Example:

```
/search?name=go&page=2
```

---

## ✅ Using net/http

```go
func searchHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	page := r.URL.Query().Get("page")

	fmt.Fprintf(w, "Name: %s, Page: %s", name, page)
}
```

### Important Points:

- `r.URL.Query()` → সব query map আকারে আনে
- `.Get()` → প্রথম value দেয়
- Missing হলে empty string return করে

---

## ⚡ Extra Tip (Multiple values)

```
/search?tag=go&tag=backend
```

```go
tags := r.URL.Query()["tag"]
```

---

## Using Gin

```go
name := c.Query("name")
page := c.DefaultQuery("page", "1")
```

### Important:

- `Query()` → optional
- `DefaultQuery()` → default value দেয়

---

# 2. Path Variables

👉 URL path এর ভিতরের dynamic value

Example:

```
/users/12
/users/12/posts/5
```

---

## ⚡ Go (net/http - Go 1.22+)

```go
mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	fmt.Fprintf(w, "User ID: %s", id)
})
```

### Multiple params:

```go
mux.HandleFunc("/users/{userId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {

	userId := r.PathValue("userId")
	postId := r.PathValue("postId")

	fmt.Println(userId, postId)
})
```

### Important:

- Go 1.22+ লাগবে
- `http.NewServeMux()` ব্যবহার করতে হবে

---

## Using Gin

```go
router.GET("/users/:id", func(c *gin.Context) {

	id := c.Param("id")

	c.JSON(200, gin.H{"id": id})
})
```

---

# 3. Request Body (JSON)

👉 POST/PUT এ data পাঠানো হয়

Example:

```json
{
  "name": "suvo",
  "age": 24
}
```

---

## Struct define

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

---

## Using net/http

```go
func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	fmt.Fprintf(w, "User: %s", user.Name)
}
```

### Important Points:

- `r.Body` = raw stream
- `json.NewDecoder` memory efficient
- `r.Body.Close()` auto handled by Go server but good practice in advanced use

---

## Using Gin

```go
func CreateUser(c *gin.Context) {

	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"name": user.Name})
}
```

---

# ⚡ Quick Comparison Table

| Type  | Example   | net/http            | Gin                |
| ----- | --------- | ------------------- | ------------------ |
| Query | ?name=go  | `r.URL.Query()`     | `c.Query()`        |
| Path  | /users/12 | `r.PathValue()`     | `c.Param()`        |
| Body  | JSON      | `json.NewDecoder()` | `ShouldBindJSON()` |

---

# 🔥 Important Extra Things (Interview Level)

## 1. Always validate input

```go
if name == "" {
	http.Error(w, "name required", 400)
	return
}
```

---

## 2. JSON decode error must handle

```go
if err != nil {
	return
}
```

---

## 3. Prefer Decoder over ReadAll

✔ Better memory
✔ Streaming support
❌ `io.ReadAll` heavy for large data

---

## 4. PathValue only works in Go 1.22+

Old Go → manual parsing needed

---

# 🎯 Final Summary

- Query → filtering/search
- Path → resource identify
- Body → data send/create/update

---
