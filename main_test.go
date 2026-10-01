package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	zeroElements := generateRandomElements(0)
	assert.Equal(t, []int{}, zeroElements)

	negElements := generateRandomElements(-3)
	assert.Equal(t, []int{}, negElements)

	elements := generateRandomElements(10)
	assert.Equal(t, 10, len(elements))
}

func TestMaximum(t *testing.T) {
	zeroV := maximum([]int{})
	assert.Equal(t, 0, zeroV)

	data := []int{1, 2, 3, 4, 5}
	maxV := maximum(data)
	assert.Equal(t, 5, maxV)

	data = []int{3}
	oneV := maximum(data)
	assert.Equal(t, 3, oneV)

	data = []int{3, 4, 5, 8, 8}
	maxV = maximum(data)
	assert.Equal(t, 8, maxV)
}

func TestMaxChunks(t *testing.T) {
	zeroV := maxChunks([]int{})
	assert.Equal(t, 0, zeroV)

	data := []int{2, 8, 9, 10, 4}
	maxV := maxChunks(data)
	assert.Equal(t, 10, maxV)

	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}
	maxV = maxChunks(data)
	assert.Equal(t, 100, maxV)
}
