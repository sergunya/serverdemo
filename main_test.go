package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIntMin(t *testing.T) {
	ans := IntMin(2, -2)
	fmt.Println(os.Hostname())
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMin1(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 2, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
