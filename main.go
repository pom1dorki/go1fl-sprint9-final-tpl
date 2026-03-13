package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	data := make([]int, size)
	for i := range data {
		data[i] = r.Intn(1_000_000_000) + 1 // 1..1_000_000_000
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(i, start, end int) {
			defer wg.Done()
			if start >= end {
				maxes[i] = 0
				return
			}

			localMax := data[start]
			for _, v := range data[start+1 : end] {
				if v > localMax {
					localMax = v
				}
			}
			maxes[i] = localMax
		}(i, start, end)
	}

	wg.Wait()
	return maximum(maxes)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", max, elapsed)
}
