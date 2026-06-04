Go Backend Industry-তে একটাই "official standard" project structure নেই। তবে বড় কোম্পানি, SaaS, Microservice এবং Enterprise Go Application-এ সাধারণত **Feature-Based + Clean Architecture inspired structure** বেশি দেখা যায়।

## Industry Standard Structure

```text
bloodmap/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── auth/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   ├── model.go
│   │   └── routes.go
│   │
│   ├── user/
│   ├── donor/
│   ├── bloodrequest/
│   └── hospital/
│
├── pkg/
│   ├── logger/
│   ├── validator/
│   └── jwt/
│
├── configs/
│
├── migrations/
│
├── docs/
│
├── scripts/
│
├── tests/
│
├── .env
├── go.mod
└── go.sum
```

---

# 1. cmd/

```text
cmd/
 └── api/
      └── main.go
```

### কাজ

Application শুরু করার Entry Point।

```go
func main() {
    app.Start()
}
```

### কেন?

একটি Project-এ একাধিক Executable থাকতে পারে।

```text
cmd/
├── api/
├── worker/
└── cron/
```

উদাহরণ:

- api → REST API Server
- worker → Background Job
- cron → Scheduled Task

Run:

```bash
go run ./cmd/api
```

---

# 2. internal/

সবচেয়ে গুরুত্বপূর্ণ Folder।

```text
internal/
├── auth/
├── user/
├── donor/
└── hospital/
```

### কাজ

Business Logic রাখা।

---

## auth Feature

```text
auth/
├── handler.go
├── service.go
├── repository.go
├── model.go
└── routes.go
```

### handler.go

HTTP Request Handle করে।

```go
func Login(c echo.Context) error
```

---

### service.go

Business Logic।

```go
func Authenticate(email, password string)
```

---

### repository.go

Database Query।

```go
func FindUserByEmail(email string)
```

---

### model.go

Struct Definition।

```go
type User struct {
    ID int
    Name string
}
```

---

### routes.go

Route Registration।

```go
func RegisterRoutes(e *echo.Echo)
```

---

## কেন internal?

Go-এর Special Folder।

```go
import "bloodmap/internal/auth"
```

শুধু নিজের Project থেকে Import করা যাবে।

অন্য Project Import করতে পারবে না।

### Benefit

- Encapsulation
- Private Code
- Better Security
- Better Maintainability

---

# 3. pkg/

```text
pkg/
├── logger/
├── validator/
└── jwt/
```

### কাজ

Reusable Component।

---

Logger Example

```go
pkg/logger/logger.go
```

```go
logger.Info("Server Started")
```

---

Validator Example

```go
validator.Validate(user)
```

---

### কেন?

অনেক Feature একই Code ব্যবহার করবে।

যেমন:

```text
auth/
user/
hospital/
```

সব Module Logger ব্যবহার করবে।

তাই Shared Code `pkg` এ রাখা হয়।

---

# 4. configs/

```text
configs/
├── dev.yaml
├── prod.yaml
└── test.yaml
```

### কাজ

Configuration Store করা।

উদাহরণ:

```yaml
server:
  port: 8080
```

---

### কেন?

Hardcoded Value Avoid করতে।

খারাপ:

```go
dbHost := "localhost"
```

ভালো:

```yaml
database:
  host: localhost
```

---

# 5. migrations/

```text
migrations/
├── 001_users.up.sql
├── 001_users.down.sql
```

### কাজ

Database Schema Manage করা।

---

Example

```sql
CREATE TABLE users(
    id SERIAL PRIMARY KEY
);
```

---

### কেন?

Database Change Track করার জন্য।

Git যেমন Code Version Control করে।

Migration Database Version Control করে।

---

# 6. docs/

```text
docs/
```

### কাজ

API Documentation।

উদাহরণ:

```text
swagger.yaml
openapi.json
```

---

### কেন?

Frontend Developer, Mobile Developer, QA Team API বুঝতে পারে।

---

# 7. scripts/

```text
scripts/
├── build.sh
├── deploy.sh
└── migrate.sh
```

### কাজ

Automation।

---

Build

```bash
./scripts/build.sh
```

Deploy

```bash
./scripts/deploy.sh
```

---

### কেন?

বারবার একই Command Manual না লিখে Automation করা।

---

# 8. tests/

```text
tests/
├── auth_test.go
└── user_test.go
```

### কাজ

Integration Test এবং E2E Test।

---

Example

```go
func TestLogin(t *testing.T)
```

---

### কেন?

Production Release-এর আগে Bug ধরার জন্য।

---

# Industry-তে সবচেয়ে জনপ্রিয় Pattern

বর্তমানে Go Industry-তে:

```text
internal/
 ├── auth/
 ├── user/
 ├── order/
 └── payment/
```

এই **Feature-Based Structure** সবচেয়ে বেশি ব্যবহৃত হয়।

কারণ:

❌ Layer Based

```text
handlers/
services/
repositories/
models/
```

বড় Project-এ 100+ File হয়ে যায়।

---

✅ Feature Based

```text
auth/
user/
payment/
```

প্রতিটি Feature-এর সব Code এক Folder-এ থাকে।

এটি maintain, debug এবং scale করা অনেক সহজ।

### Go Backend Developer হিসেবে তোমার জন্য Recommendation

যদি BloodMap, E-commerce, CMS বা SaaS Project বানাও:

```text
cmd/
internal/
pkg/
configs/
migrations/
docs/
scripts/
tests/
```

এই Structure ব্যবহার করো। এটি Mid-Level থেকে Enterprise-Level Go Backend Project-এ খুবই সাধারণ এবং Production Ready।

---

**Standard Go Project Layout** (স্ট্যান্ডার্ড গো প্রজেক্ট লেআউট) মূলত বড় বা মাঝারি আকারের অ্যাপ্লিকেশনের ক্ষেত্রে গো (Go/Golang) কমিউনিটিতে বহুল ব্যবহৃত একটি প্রজেক্ট স্ট্রাকচার। গো নিজস্ব কোনো কড়া নিয়ম চাপিয়ে না দিলেও, কোড গোছানো রাখার জন্য এই স্ট্রাকচারটি সবাই অনুসরণ করে।

---

## কেন এই স্ট্রাকচারটি ব্যবহার করবেন?

- **কাজের আলাদা বিভাজন (Separation of Concerns):** এটি কনফিগারেশন, মূল লজিক, এপিআই (API), এবং টেস্ট কোডকে একদম আলাদা আলাদা রাখে।
- **সার্কুলার ডিপেন্ডেন্সি রোধ (Prevents Circular Dependencies):** গো-তে এক প্যাকেজ অপর প্যাকেজকে এবং অপর প্যাকেজ আবার প্রথম প্যাকেজকে ইম্পোর্ট করলে (Circular import) এরর দেখায়। এই স্ট্রাকচারটি কোডকে এমনভাবে সাজাতে সাহায্য করে যাতে এই সমস্যা না হয়।
- **কোড সুরক্ষিত রাখা:** `internal/` ডিরেক্টরির কারণে অন্য কোনো এক্সটার্নাল প্রজেক্ট ভুলবশত আপনার প্রাইভেট কোড ইম্পোর্ট করতে পারে না।
- **সহজে বড় করার সুবিধা (Scalability):** একই প্রজেক্টে যদি একাধিক মেইন ফাইল থাকে (যেমন: একটা ওয়েব সার্ভার, একটা CLI টুল), তবে তারা সবাই মূল লজিক শেয়ার করে একসাথে থাকতে পারে।

---

## প্রতিটি ডিরেক্টরির কাজ ও সুবিধা

### `cmd/`

- **কাজ:** এটি আপনার অ্যাপ্লিকেশনের এন্ট্রি পয়েন্ট (যেখান থেকে প্রোগ্রাম চলা শুরু করে)। এর ভেতরে অ্যাপের নাম দিয়ে ফোল্ডার থাকে, যেমন: `cmd/api/main.go` বা `cmd/cli/main.go`।
- **সুবিধা:** মেইন ডিরেক্টরি পরিষ্কার থাকে। এখানে কোনো জটিল লজিক থাকে না—শুধু কনফিগারেশন রিড করা এবং অ্যাপ রান করার কোড থাকে।

### `internal/`

- **কাজ:** এই ফোল্ডারের কোডগুলো সম্পূর্ণ প্রাইভেট। গো কম্পাইলারের নিয়ম অনুযায়ী, এই প্রজেক্টের বাইরের অন্য কোনো প্রজেক্ট বা মডিউল `internal/` এর ভেতরের কোড ইম্পোর্ট করতে পারবে না।
- **সুবিধা:** **সুরক্ষা ও স্বাধীনতা।** আপনার মূল বিজনেস লজিক, ডাটাবেজ কোড সব এখানে থাকে। আপনি নিশ্চিন্তে এই কোড রিফ্যাক্টর বা পরিবর্তন করতে পারেন, কারণ আপনি জানেন যে বাইরের কেউ এটার ওপর নির্ভর করে বসে নেই।

### `pkg/`

- **কাজ:** এখানে এমন কোড রাখা হয় যা অন্য যেকোনো এক্সটার্নাল প্রজেক্ট বা থার্ড-পার্টি লাইব্রেরি চাইলে ইম্পোর্ট করে ব্যবহার করতে পারবে।
- **সুবিধা:** **পুনর্ব্যবহারযোগ্যতা (Reusability)।** যেমন, আপনি যদি দারুণ কোনো কাস্টম লগার (Logger) বা এপিআই ক্লায়েন্ট তৈরি করেন যা আপনার অফিসের অন্য টিমও ব্যবহার করতে চায়, তবে তা এখানে রাখা হয়।

> _নোট:_ যদি আপনার কোড অন্য কাউকে শেয়ার করার ইচ্ছা না থাকে, তবে আধুনিক গো ডেভেলপাররা `pkg/` বাদ দিয়ে সবকিছু `internal/` এ রাখেন।

### `configs/`

- **কাজ:** আপনার অ্যাপ্লিকেশনের কনফিগারেশন ফাইলের টেমপ্লেট বা ডিফল্ট কনফিগ ফাইল (যেমন: `config.yaml`, `config.json`, `.env.example`) এখানে থাকে।
- **সুবিধা:** কেন্দ্রীয় নিয়ন্ত্রণ। কোনো নতুন ডেভেলপার প্রজেক্টে আসলে সে এই ফোল্ডার দেখেই বুঝে যাবে অ্যাপটি চালাতে কী কী এনভায়রনমেন্ট ভ্যারিয়েবল বা কনফিগারেশন লাগবে।

### `migrations/`

- **কাজ:** ডাটাবেজের টেবিল তৈরি বা পরিবর্তনের SQL স্ক্রিপ্ট বা মাইগ্রেশন ফাইলগুলো (যেমন: `golang-migrate` বা `goose` এর ফাইল) এখানে থাকে।
- **সুবিধা:** ডাটাবেজের ট্র্যাক রাখা। আপনার কোড পরিবর্তনের সাথে সাথে ডাটাবেজের চেহারা কেমন পরিবর্তন হচ্ছে, তা সহজেই ভার্সন কন্ট্রোলে (যেমন Git) ট্র্যাকিং করা যায়।

### `docs/`

- **কাজ:** ডিজাইন ডকুমেন্ট, আর্কিটেকচার নোট এবং API ডকুমেন্টেশন (যেমন: Swagger/OpenAPI ফাইল) এখানে রাখা হয়।
- **সুবিধা:** যেকোনো ডেভেলপার কোডের পাশাপাশি ডকুমেন্টেশন পেয়ে যায়, ফলে কোড বুঝতে সহজ হয়।

### `scripts/`

- **কাজ:** প্রজেক্ট বিল্ড করা, ইনস্টল করা বা কোনো অটোমেটেড কাজ করার জন্য বিভিন্ন স্ক্রিপ্ট (যেমন: Bash scripts, Makefiles, Docker entrypoints) এখানে থাকে।
- **সুবিধা:** **অটোমেশন।** মেইন ফোল্ডারে হিজিবিজি স্ক্রিপ্ট না রেখে এক জায়গায় গোছানো থাকে, যা CI/CD পাইপলাইনেও সহজে ব্যবহার করা যায়।

### `tests/`

- **কাজ:** অতিরিক্ত বড় এক্সটার্নাল ইন্টিগ্রেশন টেস্ট (Integration tests) এবং টেস্টের ফেক ডাটা (Mock data) এখানে রাখা হয়।
- **সুবিধা:** সাধারণত গো-তে ইউনিট টেস্টের ফাইলগুলো মূল কোডের পাশেই (`*_test.go`) থাকে। কিন্তু বড় বা স্লো ইন্টিগ্রেশন টেস্টগুলো আলাদা রাখার জন্য এই ফোল্ডার ব্যবহার করা হয়, যাতে কোডের মূল অংশ হিজিবিজি না দেখায়।

---

একটি প্রজেক্টকে সহজে বিল্ড, রান এবং টেস্ট করার জন্য একটি `Makefile` যোগ করা দারুণ বুদ্ধি। নিচে স্ট্যান্ডার্ড গো প্রজেক্ট স্ট্রাকচারের সাথে মানানসই একটি প্রজেক্ট-রেডি `Makefile` দেওয়া হলো।

আপনার প্রজেক্টের রুট ডিরেক্টরিতে (সবচেয়ে বাইরের ফোল্ডারে) `Makefile` নামে একটি ফাইল তৈরি করুন এবং নিচের কোডটি পেস্ট করুন:

```makefile
# ভ্যারিয়েবল সেটআপ
BINARY_NAME=myapp
CMD_DIR=./cmd/api
MIGRATIONS_DIR=./migrations
DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable

.PHONY: all build run test clean migrate-up migrate-down docs-gen help

# ডিফল্ট কমান্ড (শুধু 'make' লিখলে এটি রান হবে)
all: test build

## build: অ্যাপ্লিকেশনটি বিল্ড করার জন্য
build:
	@echo "Building binary..."
	go build -o bin/$(BINARY_NAME) $(CMD_DIR)/main.go

## run: অ্যাপ্লিকেশনটি সরাসরি রান করার জন্য
run:
	@echo "Running application..."
	go run $(CMD_DIR)/main.go

## test: সব ইউনিট টেস্ট রান করার জন্য
test:
	@echo "Running tests..."
	go test -v -race ./...

## clean: বিল্ড করা ফাইল মুছে ফেলার জন্য
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

## deps: ডিপেন্ডেন্সি ডাউনলোড এবং টাইডি করার জন্য
deps:
	@echo "Tidying up Go modules..."
	go mod tidy
	go mod download

## migrate-up: ডাটাবেজ মাইগ্রেশন রান করার জন্য (golang-migrate ব্যবহার করলে)
migrate-up:
	@echo "Running database migrations up..."
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

## migrate-down: ডাটাবেজ মাইগ্রেশন রোলব্যাক করার জন্য
migrate-down:
	@echo "Running database migrations down..."
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

## docs-gen: সোয়াগার বা এপিআই ডকুমেন্টেশন জেনারেট করার জন্য (swag CLI ব্যবহার করলে)
docs-gen:
	@echo "Generating API documentation..."
	swag init -g $(CMD_DIR)/main.go -o ./docs

## help: কি কি কমান্ড আছে তা দেখার জন্য
help:
	@echo "Available commands:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

```

---

## কেন এই Makefile ব্যবহার করবেন এবং এর সুবিধা:

- **কম্পাইলিং সহজ করা:** গো-তে বারবার `go build -o bin/myapp ./cmd/api/main.go` লেখার চেয়ে শুধু `make build` লেখা অনেক সহজ ও সময় সাশ্রয়ী।
- **ভুল কমানো:** ডাটাবেজ মাইগ্রেশনের মতো বড় কমান্ড (`migrate -path ./migrations...`) মুখস্থ রাখা কঠিন। `make migrate-up` দিয়ে এটি এক সেকেন্ডে করা সম্ভব।
- **নতুন ডেভেলপারদের সুবিধা:** কোনো নতুন ডেভেলপার আপনার প্রজেক্টে কাজ করতে আসলে সে শুধু `make help` লিখলেই বুঝে যাবে প্রজেক্টটি রান, টেস্ট বা বিল্ড করার উপায় কী।
- **CI/CD ফ্রেন্ডলি:** গিটহাব অ্যাকশন বা যেকোনো CI/CD পাইপলাইনে প্রজেক্ট টেস্ট ও বিল্ড করার জন্য শুধুমাত্র `make test` এবং `make build` কমান্ড দুটি লিখে দিলেই চলে।

> **জরুরি নোট:** `Makefile`-এ প্রতিটি কমান্ডের ভেতরের লাইনে (যেমন `@echo` বা `go build` এর শুরুতে) অবশ্যই **একটি Tab** ব্যবহার করতে হবে, স্পেস (Space) দিলে Makefile কাজ করবে না।

---
