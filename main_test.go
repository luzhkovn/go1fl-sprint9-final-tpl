package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	input := 10
	expected := 10

	result := generateRandomElements(input)

	assert.Equal(t, expected, len(result))
}

func TestMaximum1(t *testing.T) {
	slice := []int{10, 20, 50, 30}
	result := maximum(slice)

	if result != 50 {
		t.Errorf("Ожидали  50, но получили %d", result)
	}
}

func TestMaxChunks1(t *testing.T) {
	slice := []int{10, 20, 50, 30}
	result := maxChunks(slice)
	if result != 50 {
		t.Errorf("Ожидали 50, но получили %d", result)
	}
}

func TestMaximum2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "обычный слайс чисел",
			input:    []int{10, 20, 50, 30},
			expected: 50,
		},
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := maximum(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
func TestMaxChunks2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "обычный слайс чисел",
			input:    []int{10, 20, 50, 30},
			expected: 50,
		},
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := maxChunks(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
