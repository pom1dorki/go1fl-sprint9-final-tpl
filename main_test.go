package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"zero size", 0, 0},
		{"positive size", 5, 5},
		{"negative size", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("generateRandomElements() len = %v, want %v", len(got), tt.want)
			}
			for _, v := range got {
				if v <= 0 {
					t.Errorf("generateRandomElements() got non-positive value %v", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"multiple elements", []int{1, 3, 2, 5, 4}, 5},
		{"all equal", []int{2, 2, 2}, 2},
		{"negative numbers", []int{-1, -3, -2}, -1},
		{"mixed signs", []int{-1, 3, 0}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty", []int{}, 0},
		{"one element", []int{42}, 42},
		{"small less than chunks", []int{1, 5, 3}, 5},
		{"multiple elements", []int{1, 3, 2, 8, 5, 7, 4}, 8},
		{"all equal", []int{10, 10, 10, 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
