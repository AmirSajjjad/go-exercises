# 01 - Concurrent Line Counter

This Go program reads a text file and counts the number of lines **concurrently** using multiple goroutines.  
It demonstrates concepts such as **worker pools**, **context timeouts**, **channel-based semaphores**, and **synchronization with WaitGroups**.

---

## 🧠 Concepts Covered

### 1. File Reading and Buffering
- Uses `bufio.NewReader` and a fixed chunk size (`4 KB`) to read the file progressively.
- The code validates file existence, regular type, and reasonable size limits (<100 MB).

### 2. Worker Pool & Concurrency Control
- The constant `maxActiveWorkers = 40` limits the number of concurrent goroutines.
- A **channel semaphore** (`tokens chan struct{}`) controls concurrency by ensuring only a fixed number of goroutines are active at once:
  ```go
  tokens := make(chan struct{}, maxActiveWorkers)
  tokens <- struct{}{} // acquire slot
  <-tokens              // release slot
  ```
- This prevents memory overload or excessive CPU scheduling.

### 3. Context with Timeout
- The program uses `context.WithTimeout` to automatically cancel goroutines if processing takes too long:
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
  defer cancel()
  ```
- If the timeout triggers, all worker goroutines stop gracefully.

### 4. Synchronization with WaitGroup
- A `sync.WaitGroup` is used to track all worker goroutines.
- When all goroutines finish, the result channel (`ch`) is closed to signal completion.

### 5. Channels for Communication
- Results (line counts per chunk) are sent via a buffered channel.
- Once all workers complete, the main goroutine sums up all counts to get the total number of lines.

---

## ⚙️ How It Works

1. Open and validate the target file.
2. Create a reader and a limited-size buffer.
3. For each chunk read:
   - Start a worker goroutine that counts `\n` in the chunk.
   - Limit concurrency using the semaphore (`tokens`).
4. Wait for all workers to finish and aggregate the total line count.

---

## 🧩 Key Constants

| Constant | Meaning | Default |
|-----------|----------|----------|
| `maxActiveWorkers` | Max number of active goroutines | `40` |
| `chunkSize` | File chunk size per read | `4 KB` |
| `contextTimeOut` | Time limit for processing | `2s` |

---

## 🧪 Example Output

```bash
📦 File size: 32418 bytes, buffer size: 80
Number of lines: 912
Number of Go Routines: 80
```

---

## 🪄 Educational Notes

- **Semaphores with channels** are an elegant idiomatic way to limit concurrency in Go — better than using raw mutexes for this case.  
- **Context** is essential for safe goroutine cancellation in long-running operations.
- **Buffered channels** improve performance when collecting results concurrently.
- Reading the file in chunks keeps memory usage low even for larger files.
- `strings.Count(chunk, "\n")` is a simple but effective method for line counting.

---

## 🚀 Run the Program

```bash
go run concurrency_line_counter.go
```

Or initialize a Go module for this exercise first:

```bash
go mod init github.com/AmirSajjjad/go-exercises/exercises/01-line-counter-concurrent
go run concurrency_line_counter.go
```
