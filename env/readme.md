### 🔹 What is `godotenv`?

godotenv is a Go library that loads environment variables from a `.env` file into your application at runtime.

Instead of hardcoding secrets like:

- DB password
- API keys
- JWT secret

You store them in a `.env` file.

---

## 🔹 Why use `godotenv`?

Because production apps must **not hardcode sensitive data**.

### ❌ Bad practice:

```go
dbPassword := "123456"
```

### ✅ Good practice:

```
DB_PASSWORD=123456
```

Then load it using `godotenv`.

---

## 🔹 Benefits

### 1. Security

Secrets are not inside code → safer for GitHub

### 2. Environment separation

Different configs for:

- local
- staging
- production

### 3. Easy deployment

Change `.env` only, no code change

### 4. Clean architecture

Config is separated from business logic

---

## 🔹 How to use `godotenv`

### 1. Install

```bash
go get github.com/joho/godotenv
```

---

### 2. Create `.env` file

```
PORT=8080
DB_HOST=localhost
DB_USER=admin
DB_PASSWORD=secret123
DB_NAME=myapp
JWT_SECRET=myjwtsecret
```

---

### 3. Load `.env` in Go

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")

	fmt.Println("Server running on port:", port)
	fmt.Println("DB User:", dbUser)
}
```

---

## 🔹 Production-Level Architecture (Important)

In real-world systems, you **don’t directly use os.Getenv everywhere**.

You create a **config package**.

---

## 🔥 Production Example (Clean Architecture)

### 📁 Project Structure

```
myapp/
 ├── config/
 │     └── config.go
 ├── main.go
 ├── .env
 └── go.mod
```

---

## 🔹 config/config.go

```go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT       string
	DB_HOST    string
	DB_USER    string
	DB_PASS    string
	DB_NAME    string
	JWT_SECRET string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env not found, using system env")
	}

	return &Config{
		PORT:       getEnv("PORT", "8080"),
		DB_HOST:    getEnv("DB_HOST", "localhost"),
		DB_USER:    getEnv("DB_USER", "root"),
		DB_PASS:    getEnv("DB_PASSWORD", ""),
		DB_NAME:    getEnv("DB_NAME", "test"),
		JWT_SECRET: getEnv("JWT_SECRET", "defaultsecret"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
```

---

## 🔹 main.go

```go
package main

import (
	"fmt"
	"myapp/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println("Server running on:", cfg.PORT)
	fmt.Println("DB Host:", cfg.DB_HOST)
}
```

---

## 🔹 Why this is production-level?

✔ Centralized config
✔ Fallback values
✔ Works without `.env` in production (Docker/CI/CD)
✔ Clean separation of concerns
✔ Easier testing

---

## 🔹 Real Production Flow

### Development:

```
.env file → godotenv → app uses config
```

### Production (Docker / Cloud):

```
Environment variables (system level)
→ godotenv skipped
→ os.Getenv reads directly
```

---

## 🔥 Important Production Tip

In production you often **DO NOT use `.env` file at all**.

Instead:

- Docker environment variables
- Kubernetes Secrets
- AWS Parameter Store / SSM
- CI/CD secrets

---

## 🔹 When NOT to use `godotenv`

❌ Large-scale Kubernetes systems
❌ Cloud-native secret managers already used
❌ When env vars are injected directly

---

## 🔥 Summary

- `godotenv` loads `.env` into Go app
- Useful for development and local setup
- Keeps secrets outside code
- Production apps use config wrapper + real env vars

---

### 🔹 `godotenv` কী?

godotenv হলো Go ভাষার একটি লাইব্রেরি, যেটা `.env` ফাইল থেকে environment variables লোড করে অ্যাপে ব্যবহার করতে সাহায্য করে।

অর্থাৎ, তুমি যেসব sensitive data (password, API key, JWT secret) কোডে লিখবে না, সেগুলো `.env` ফাইলে রাখবে।

---

## 🔹 কেন `godotenv` ব্যবহার করা হয়?

### ❌ খারাপ পদ্ধতি:

```go
dbPassword := "123456"
```

এভাবে কোডে password রাখা নিরাপদ না।

---

### ✅ ভালো পদ্ধতি:

`.env` ফাইল:

```
DB_PASSWORD=123456
```

তারপর Go অ্যাপে লোড করা হয়।

---

## 🔹 `godotenv` ব্যবহার করার সুবিধা

### 🔐 1. নিরাপত্তা

Sensitive data কোডের বাইরে থাকে।

### 🌍 2. আলাদা environment support

- development
- staging
- production

### 🚀 3. সহজ deployment

কোড পরিবর্তন না করে শুধু `.env` পরিবর্তন করলেই হয়।

### 🧹 4. clean code architecture

config আলাদা থাকে, business logic আলাদা থাকে।

---

## 🔹 কীভাবে ব্যবহার করবে

### 1. Install করো

```bash
go get github.com/joho/godotenv
```

---

### 2. `.env` ফাইল তৈরি করো

```
PORT=8080
DB_HOST=localhost
DB_USER=admin
DB_PASSWORD=secret123
DB_NAME=myapp
JWT_SECRET=myjwtsecret
```

---

### 3. Go কোডে লোড করো

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file load করা যায়নি")
	}

	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")

	fmt.Println("Server চলছে port:", port)
	fmt.Println("DB User:", dbUser)
}
```

---

## 🔥 Production-level বাস্তব ব্যবহার

বাস্তব প্রজেক্টে সরাসরি `os.Getenv()` ছড়িয়ে না দিয়ে একটা **config package** ব্যবহার করা হয়।

---

## 📁 Project Structure

```
myapp/
 ├── config/
 │     └── config.go
 ├── main.go
 ├── .env
```

---

## 🔹 config.go

```go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT       string
	DB_HOST    string
	DB_USER    string
	DB_PASS    string
	DB_NAME    string
	JWT_SECRET string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env পাওয়া যায়নি, system env ব্যবহার হবে")
	}

	return &Config{
		PORT:       getEnv("PORT", "8080"),
		DB_HOST:    getEnv("DB_HOST", "localhost"),
		DB_USER:    getEnv("DB_USER", "root"),
		DB_PASS:    getEnv("DB_PASSWORD", ""),
		DB_NAME:    getEnv("DB_NAME", "test"),
		JWT_SECRET: getEnv("JWT_SECRET", "default"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
```

---

## 🔹 main.go

```go
package main

import (
	"fmt"
	"myapp/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println("Server port:", cfg.PORT)
	fmt.Println("DB host:", cfg.DB_HOST)
}
```

---

## 🔥 বাস্তবে কীভাবে কাজ করে?

### Development:

```
.env → godotenv → Go app
```

### Production:

```
System environment variables → Go app (godotenv না থাকলেও চলবে)
```

---

## ⚠️ Production টিপস

বাস্তব production system এ অনেক সময়:

- Docker environment variables
- Kubernetes Secrets
- AWS SSM / Secret Manager

ব্যবহার করা হয় `.env` ফাইল না রেখে।

---

## 🔹 সংক্ষেপে

- `godotenv` = `.env` ফাইল থেকে env variable লোড করে
- নিরাপত্তা + clean architecture দেয়
- development এ খুব useful
- production এ অনেক সময় direct env ব্যবহার করা হয়

---

চাও তাহলে আমি তোমাকে next step দেখাতে পারি:
👉 Go + Docker production config system
👉 বা Viper vs godotenv comparison (কোনটা কখন ব্যবহার করবে)
