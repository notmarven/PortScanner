# 🔍 Port Scanner in Go

Hey there! 👋  
This is my **first ever project in Golang**, and I decided to build a basic but functional **port scanner** something I’ve always wanted to understand better.

---

## 💡 What does it do?

It scans a range of TCP ports on any host you enter and tells you whether each one is open or closed.  
It runs super fast thanks to Go's **goroutines** (basically threads, but better), and I learned a ton about:

- Go basics (variables, input, types)
- Structs and slices
- Concurrency (WaitGroups, Mutex)
- Sorting data cleanly

---

## 🛠️ How to use it

### ✅ Requirements

- You need Go installed. [Get it here](https://go.dev/dl/)

---

### ▶️ Running the scanner

```bash
go run main.go
```

### 🧠 Why I built this

I’m new to Go, and I wanted something practical and network-related to help me learn how Go handles:

Threads (goroutines)

Shared memory (mutex)

Timing

Input/output

Turns out writing a port scanner touches on all of that and it was actually fun to build!


🚧 What I might add next
 - Add color to open/closed ports (green/red)
 - Module to grab (MOTD, version, player count)

