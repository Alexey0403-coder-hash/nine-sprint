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
		return []int{}
	}
	data := make([]int, size)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(size) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxInt := data[0]
	for _, i := range data {
		if i > maxInt {
			maxInt = i
		}
	}
	return maxInt
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup
	chunkSize := len(data) / CHUNKS
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)

		go func(start int, end int, i int) {
			defer wg.Done()
			maxValues[i] = maximum(data[start:end])
		}(start, end, i)
	}
	wg.Wait()
	maxS := maximum(maxValues)
	return maxS
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()

	max := maximum(data)

	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()

	max = maxChunks(data)

	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
