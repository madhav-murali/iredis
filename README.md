# iredis

A high-performance Redis clone written in **Go**, focused on low-latency execution and the **RESP (Redis Serialization Protocol)**. 

## 🚀 Overview
**iredis** is an in-memory data store built from the ground up to explore the internals of distributed systems, networking, and concurrent data access. It achieves **$O(1)$** time complexity for core operations and supports a variety of data types, including Strings and Lists.

### Key Features
* **RESP v2 Protocol:** Hand-rolled parser for bulk strings, arrays, and integers.
* **Concurrent Networking:** Leverages Go's goroutines to handle multiple TCP client connections simultaneously.
* **Data Structures:** * **Strings:** `GET`, `SET`, `ECHO`.
    * **Lists:** Full support for linked-list operations including `LPUSH`, `RPUSH`, `LPOP`, and `LRANGE`.
* **Performance:** Designed for $O(1)$ lookup and insertion for primary data storage.

---

## 🛠 Tech Stack
* **Language:** Go (Golang)
* **Transport:** TCP Sockets
* **Serialization:** RESP (Redis Serialization Protocol)

---

## 🏗 Architecture
The project is modularized to separate the networking logic from the command execution engine.

* **Command Handler:** A central switchboard that routes RESP-parsed commands to the appropriate storage logic.
* **Storage Engine:** Thread-safe storage using Go's internal maps and custom list implementations to ensure data integrity during concurrent access.
* **RESP Formatter:** A dedicated utility package to ensure all outbound data strictly adheres to the Redis protocol.

### Implementation Snippet (Command Dispatcher)
```go
case "RPUSH":
    length := lst.RPUSH(elements[1], elements[2:])
    handleWrite(*Writer, ":"+strconv.Itoa(length)+"\r\n")
case "LRANGE":
    handleWrite(*Writer, resp.ArrayRESP(lst.LRANGE(elements[1], elements[2], elements[3])))
```

---

## 🚦 Getting Started

### Installation
```bash
git clone https://github.com/your-username/iredis.git
cd iredis
```

### Running the Server
```bash
go run main.go
```

### Testing
You can interact with **iredis** using the standard `redis-cli`:
```bash
redis-cli LPUSH mylist "world"
redis-cli LPUSH mylist "hello"
redis-cli LRANGE mylist 0 -1
# Output: 1) "hello" 2) "world"
```

---

## 📈 Roadmap
- [x] Basic RESP Parsing
- [x] Key-Value (String) support
- [x] List support (`LPUSH`, `RPUSH`, `LPOP`, `LLEN`)
- [ ] Set/Hash support
- [ ] Persistence (RDB/AOF)
- [ ] Pub/Sub Architecture

