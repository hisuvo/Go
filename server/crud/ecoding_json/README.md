Go-এর `encoding/json` প্যাকেজে ৩টা জিনিস খুব গুরুত্বপূর্ণ:

- `json.Marshal`
- `json.NewEncoder`
- `json.NewDecoder`

এরা একই কাজ করে না, বরং different use-case এ use হয়।

---

# 1. json.Marshal (Encode to []byte)

## কাজ:

Go struct → JSON `[]byte`

## Example:

```go
data, err := json.Marshal(user)
```

## Output:

```json
[]byte → {"name":"suvo","age":24}
```

---

## Key Points

- পুরো data memory তে রাখে
- `[]byte` return করে
- small/medium data এর জন্য ভালো
- manual write লাগে (w.Write)

---

## Problem

```go
json.Marshal(largeData)
```

➡️ বেশি memory use করে

---

# 2. json.NewEncoder (Write directly to stream)

## কাজ:

Go struct → HTTP response / file / stream

---

## Example:

```go
json.NewEncoder(w).Encode(user)
```

---

## Key Points

- direct stream এ write করে
- extra memory buffer লাগে না
- best for API response
- production standard

---

## Flow

```text
struct → encoder → http.ResponseWriter
```

---

## Advantage

- fast
- memory efficient
- clean code

---

# 3. json.NewDecoder (Read from stream)

## কাজ:

JSON → Go struct

---

## Example:

```go
json.NewDecoder(r.Body).Decode(&user)
```

---

## Key Points

- request body directly read করে
- streaming input handle করে
- large JSON safe
- no full memory load

---

## Flow

```text
http.Request.Body → decoder → struct
```

---

# Comparison Table

| Feature     | Marshal           | NewEncoder      | NewDecoder      |
| ----------- | ----------------- | --------------- | --------------- |
| Direction   | struct → JSON     | struct → stream | JSON → struct   |
| Output      | []byte            | direct write    | struct          |
| Memory      | high              | low             | low             |
| Use case    | manual processing | API response    | request parsing |
| Performance | medium            | high            | high            |

---

# Real World Use Case

## 1. GET API (Send data)

```go
json.NewEncoder(w).Encode(users)
```

---

## 2. POST API (Receive data)

```go
json.NewDecoder(r.Body).Decode(&user)
```

---

## 3. Marshal (when needed manual control)

```go
data, _ := json.Marshal(user)
fmt.Println(string(data))
```

---

# Important Concept (Interview Gold)

## ❌ Wrong approach

```go
data, _ := json.Marshal(r.Body)
```

👉 কারণ r.Body JSON না, এটা stream

---

## ❌ Wrong in API

```go
w.Write(json.Marshal(user))
```

👉 inefficient + error handling weak

---

## ✅ Best Practice

```go
json.NewEncoder(w).Encode(user)
```

```go
json.NewDecoder(r.Body).Decode(&user)
```

---

# Memory Concept

## Marshal

```text
Full JSON stored in memory ([]byte)
```

## Encoder/Decoder

```text
Stream-based processing (no full load)
```

---

# Performance Idea

| Scenario        | Best choice |
| --------------- | ----------- |
| API response    | Encoder     |
| API request     | Decoder     |
| logging/debug   | Marshal     |
| file small JSON | Marshal     |

---

# Interview Questions (Very Important)

## Q1: Difference between Marshal and Encoder?

**Answer:**
Marshal returns JSON as byte slice, Encoder writes directly to output stream (like HTTP response), so Encoder is more memory efficient.

---

## Q2: Why NewDecoder is used with r.Body?

**Answer:**
Because r.Body is a stream, not full JSON data. Decoder reads it efficiently without loading entire payload into memory.

---

## Q3: Which is best for REST API?

**Answer:**

- POST request → `json.NewDecoder`
- GET response → `json.NewEncoder`

---

## Q4: Why Marshal is not preferred in API?

**Answer:**
Because it creates full JSON in memory before sending, which is inefficient for large data.

---

## Q5: What is stream processing in encoding/json?

**Answer:**
Processing data piece by piece instead of loading full data into memory (used by Encoder/Decoder).

---

# Final Short Summary

- `Marshal` → struct → JSON bytes
- `Encoder` → struct → HTTP response (stream)
- `Decoder` → request JSON → struct (stream)

---

If you want, next আমি তোমাকে **real CRUD REST API structure (best architecture)** Go দিয়ে বানিয়ে দেখাতে পারি।
