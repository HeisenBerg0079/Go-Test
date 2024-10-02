package main

import "testing"

func TestPyramidPattern_PrintPyramid(t *testing.T) {
	tests := []struct {
		name     string
		p        *PyramidPattern
		expected string
	}{
		{
			name:     "5 rows pyramid",
			p:        &PyramidPattern{Rows: 5},
			expected: "      *       \n    * * *     \n  * * * * *   \n* * * * * * * \n     *      \n    * * *     ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.PrintPyramid()
		})
	}
}
