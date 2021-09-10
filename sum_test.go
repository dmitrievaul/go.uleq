package main

import "testing"

func Sum(numbers [2]int) int {
	sum := 0
	for i := 0; i < 2; i++ {
		sum += numbers[i]
	}
	return sum
}
func TestSum(t *testing.T) {

	numbers := [2]int{10, 10}

	got := Sum(numbers)
	want := 20

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

