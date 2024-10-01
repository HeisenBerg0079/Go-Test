package main

import "testing"

func TestPyramidPattern_PrintPyramid(t *testing.T) {
	tests := []struct {
		name     string
		p        *PyramidPattern
		expected string
	}{
		{
			name:     "4 rows pyramid",
			p:        &PyramidPattern{Rows: 4},
			expected: "   *   \n  ***  \n ***** \n   *   \n  ***  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.PrintPyramid()
		})
	}
}
