package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, value := range data {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	results := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {

		var start int
		var end int
		start = i * chunkSize
		end = start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(partSlice []int, idx int) {
			localmax := maximum(partSlice)
			results[idx] = localmax
			wg.Done()
		}(data[start:end], i)
	}
	wg.Wait()
	return maximum(results)
}

func main() {
	var maxVal int
	var elapsed int64
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxVal = maximum(elements)
	elapsed = time.Since(startTime).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxVal = maxChunks(elements)
	elapsed = time.Since(startTime).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)
}
