package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	input := 10
	expected := 10

	result := generateRandomElements(input)

	if len(result) != expected {
		t.Errorf("Ожидали длину %d, но получили %d", expected, len(result))
	}
}

func TestMaximum(t *testing.T) {
	slice := []int{10, 20, 50, 30}
	result := maximum(slice)

	if result != 50 {
		t.Errorf("Ожидали  50, но получили %d", result)
	}
}

func TestMaxChunks(t *testing.T) {
	slice := []int{10, 20, 50, 30}
	result := maxChunks(slice)
	if result != 50 {
		t.Errorf("Ожидали 50, но получили %d", result)
	}
}
