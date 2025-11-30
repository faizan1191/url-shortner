# **URL Shortener**

A simple, high-performance URL shortener built in **Go** using **Gin** and **Redis**.

* Shorten long URLs to 6-character codes
* Redirect short codes to original URLs
* Persistent storage using Redis
* Easy to swap storage (Memory or Redis)

---

## **Features**

* POST `/shorten` → create a short URL
* GET `/:code` → redirect to original URL
* Memory store for development / Redis store for persistence
* Storage-agnostic design (interface-based)

---

## **Prerequisites**

* Go 1.21+
* Redis (for persistent storage)

```bash
# macOS
brew install redis
brew services start redis
redis-cli ping  # should return PONG
```

---

## **Installation & Run**

```bash
# Clone the repo
git clone https://github.com/faizan1191/url-shortener.git
cd url-shortener

# Install dependencies
go mod tidy

# Run the server
go run .
```

* Server runs at: `http://localhost:8080`

---

## **API Usage**

### **1. Shorten URL**

```bash
curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url":"https://google.com"}'
```

**Response:**

```json
{
  "short_url": "http://localhost:8080/aB9xP2"
}
```

---

### **2. Redirect**

```bash
curl -I http://localhost:8080/aB9xP2
```

**Response Headers:**

```
HTTP/1.1 301 Moved Permanently
Location: https://google.com
```

* Open in a browser → redirects to the original URL

---

## **Storage**

* **MemoryStore** → ephemeral, for development
* **RedisStore** → persistent, recommended for production

```go
// Memory
store := storage.NewMemoryStore()

// Redis
store := storage.NewRedisStore("localhost:6379")
```

---

## **Project Structure**

```
url-shortener/
 ├── main.go
 ├── router/router.go
 ├── handlers/url.go
 ├── storage/memory.go
 ├── storage/redis.go
 └── utils/shortener.go
```

---

## **Notes**

* Redis persistence: enable RDB or AOF to survive server restarts
* Gin runs in debug mode by default (`Ctrl+C` to stop, `go run .` to restart)
* Extensible: add analytics, custom codes, expiration, or DB support

---

## **License**

MIT License
