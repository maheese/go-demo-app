package main

import "testing"

func TestOne(t *testing.T) {
	got := 1 + 1
	if got != 2 {
		t.Errorf("1 + 1 = %d; want 2", got)
	}
}

func TestTwo(t *testing.T) {
	got := 2 * 2
	if got != 4 {
		t.Errorf("2 * 2 = %d; want 4", got)
	}
}
