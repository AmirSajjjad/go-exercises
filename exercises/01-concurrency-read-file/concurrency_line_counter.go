package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

const maxActiveWorkers = 40
const chunkSize = 4 * 1024 // 4 KB
const contextTimeOut = 2 * time.Second
const filename = "asdf.txt"

func ReadFile(filename string) (*os.File, int, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, 0, fmt.Errorf("file %s does not exist", filename)
	}
	if !info.Mode().IsRegular() {
		return nil, 0, fmt.Errorf("not a regular file")
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, 0, fmt.Errorf("error opening the file: %w", err)
	}

	fileSize := info.Size()
	if fileSize > 100*1024*1024 {
		return nil, 0, fmt.Errorf("file size too big")
	}

	bufferSize := int(fileSize / chunkSize)
	if bufferSize < maxActiveWorkers*2 {
		bufferSize = maxActiveWorkers * 2
	}

	fmt.Printf("📦 File size: %d bytes, buffer size: %d\n", fileSize, bufferSize)
	return f, bufferSize, nil
}

func LineCounterWorker(ctx context.Context, chunk string, resultCh chan<- int, wg *sync.WaitGroup, tokens chan struct{}) {
	defer wg.Done()
	defer func() { <-tokens }()

	select {
	case <-ctx.Done():
		return
	case resultCh <- strings.Count(chunk, "\n"):
		return
	}
}

func main() {
	f, bufferSize, err := ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeOut)
	defer cancel()

	var wg sync.WaitGroup
	ch := make(chan int, bufferSize)
	tokens := make(chan struct{}, maxActiveWorkers)
	goRoutineCounter := 0

	reader := bufio.NewReader(f)
	buf := make([]byte, chunkSize)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
			tokens <- struct{}{}
			chunk := string(buf[:n])
			wg.Add(1)
			go LineCounterWorker(ctx, chunk, ch, &wg, tokens)
			goRoutineCounter++
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file:", err)
			break
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalLines := 0
	for count := range ch {
		totalLines += count
	}

	fmt.Println("Number of lines:", totalLines)
	fmt.Println("Number of Go Routines:", goRoutineCounter)
}
