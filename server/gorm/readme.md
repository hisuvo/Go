# GORM কী?

**GORM** হলো Go language এর সবচেয়ে popular ORM (Object Relational Mapper) লাইব্রেরি।

👉 Full form: **Go Object Relational Mapping**

---

# 🧠 সহজভাবে বুঝলে

Normally database query:

```sql
SELECT * FROM users WHERE id = 1;
```

GORM দিয়ে:

```go
db.First(&user, 1)
```

👉 মানে SQL না লিখেও Go struct দিয়ে database handle করা যায়।

---

# 🔥 GORM কী করে?

GORM তোমার জন্য করে দেয়:

- CRUD (Create, Read, Update, Delete)
- Table auto migration
- Relationship (has one, has many)
- Query builder
- Transactions

---

# 📦 Install

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

(বা MySQL)

---

# ⚡ Basic Example

## 1. Model (Table)

```go
type User struct {
	ID   uint
	Name string
	Age  int
}
```

---

## 2. DB Connection

```go
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=123 dbname=test port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect db")
	}

	fmt.Println("Database connected")
}
```

---

## 3. Auto Table Create

```go
db.AutoMigrate(&User{})
```

👉 এটা SQL ছাড়া table বানিয়ে দেয়

---

## 4. Insert Data (Create)

```go
user := User{Name: "Suvo", Age: 24}

db.Create(&user)
```

---

## 5. Read Data (Get)

```go
var user User

db.First(&user, 1) // id = 1
```

---

## 6. Update Data

```go
db.Model(&User{}).Where("id = ?", 1).Update("age", 25)
```

---

## 7. Delete Data

```go
db.Delete(&User{}, 1)
```

---

# 🔥 GORM vs Raw SQL

| Feature        | Raw SQL | GORM   |
| -------------- | ------- | ------ |
| Code length    | Long    | Short  |
| Safety         | Low     | High   |
| Learning curve | Easy    | Medium |
| Productivity   | Low     | High   |

---

# 🧠 Advanced Features

## 1. Relationships

### One-to-Many

```go
type User struct {
	ID    uint
	Posts []Post
}

type Post struct {
	ID     uint
	Title  string
	UserID uint
}
```

---

## 2. Transaction

```go
db.Transaction(func(tx *gorm.DB) error {

	tx.Create(&User{Name: "A"})
	tx.Create(&User{Name: "B"})

	return nil
})
```

---

## 3. Auto Migration

```go
db.AutoMigrate(&User{}, &Post{})
```

---

# ⚡ Why GORM use করা হয়?

- SQL avoid করা যায়
- Fast development
- Clean code
- Backend scalable হয়

---

# 🎯 Simple Definition (Interview)

👉 GORM is an ORM library for Go that allows developers to interact with databases using Go structs instead of writing raw SQL queries.

---

# 💡 Real Backend Flow

```text
HTTP Request → Go Handler → GORM → Database → Response
```

---

If you want next step, I can show:

- GORM full CRUD project (real API)
- MongoDB vs GORM comparison
- or JWT + GORM authentication system

Just বলো 👍

---

## ORM library কী? (সহজভাবে + example)

ORM মানে:
👉 **Object Relational Mapping**

এটা একটা tool/library যেটা দিয়ে তুমি **database (SQL)** কে **Go struct/object** দিয়ে handle করতে পারো।

---

# 🧠 সহজ ভাষায় বুঝো

Normal SQL এ তুমি লিখো:

```sql id="sql1"
SELECT * FROM users WHERE id = 1;
```

👉 কিন্তু ORM দিয়ে তুমি লিখো:

```go id="go1"
db.First(&user, 1)
```

👉 মানে: SQL না লিখে Go code দিয়েই database কাজ করা।

---

# 📦 Example (Real Life Concept)

ধরো তোমার database table:

## users table

| id  | name | age |
| --- | ---- | --- |
| 1   | Suvo | 24  |

---

## Without ORM (Raw SQL)

```go id="raw1"
rows, _ := db.Query("SELECT name, age FROM users WHERE id = ?", 1)
```

👉 এখানে তুমি SQL নিজে লিখছো

---

## With ORM (GORM)

```go id="orm1"
type User struct {
	ID   uint
	Name string
	Age  int
}

var user User
db.First(&user, 1)
```

👉 এখানে SQL নাই
👉 শুধু struct + function

---

# 🔥 ORM কীভাবে কাজ করে?

ORM এই কাজগুলো করে দেয়:

### 1. Struct → Table mapping

```go id="map1"
type User struct {
	ID   uint
	Name string
}
```

👉 এটা database এ convert হয়:

```sql id="map2"
users table
```

---

### 2. Object → Row insert

```go id="ins1"
user := User{Name: "Suvo", Age: 24}
db.Create(&user)
```

👉 SQL internally:

```sql id="ins2"
INSERT INTO users (name, age) VALUES ('Suvo', 24);
```

---

### 3. Row → Object read

```go id="read1"
db.First(&user, 1)
```

👉 SQL internally:

```sql id="read2"
SELECT * FROM users WHERE id = 1;
```

---

# 🧠 Simple Analogy

ORM = Translator 🧑‍💻

👉 তুমি বলো Go language এ
👉 ORM সেটা SQL এ convert করে দেয়

---

# ⚡ Real Life Example

## Without ORM:

👉 তুমি electrician কে directly wire connect করছো (manual SQL)

## With ORM:

👉 তুমি switch চাপ দাও
👉 system নিজে wiring handle করে

---

# 🔥 Why ORM use করা হয়?

### Advantages:

- SQL লিখতে হয় না
- Faster development
- Less error
- Clean code

---

# ❌ Disadvantages:

- Slow (sometimes)
- Complex query hard
- Less control than raw SQL

---

# 🎯 Final Simple Definition

👉 ORM is a tool that allows you to work with database using programming language objects instead of writing raw SQL queries.

---

# 💡 One Line Memory Trick

👉 **ORM = SQL কে Go object দিয়ে control করার system**

---

If you want, I can next show:

- GORM real CRUD mini project
- or ORM vs Raw SQL deep comparison
- or database design for backend (important)

Just বলো 👍
