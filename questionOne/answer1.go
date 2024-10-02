package main

import "fmt"

type PyramidPattern struct {
	Rows int
}

func NewPyramidPattern(rows int) *PyramidPattern {
	return &PyramidPattern{Rows: rows}
}

func (p *PyramidPattern) PrintPyramid() {
	//Top
	for i := 1; i <= p.Rows-1; i++ {
		// Print spaces
		for j := 1; j <= p.Rows-(i+1); j++ {
			fmt.Print("  ")
		}
		// Print stars
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	//Bottom
	for i := 1; i <= 2; i++ {
		// Print spaces
		for j := 1; j <= p.Rows-(i+1); j++ {
			fmt.Print("  ")
		}
		// Print stars
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	pyramid := NewPyramidPattern(4)
	fmt.Println("This is the pyramid pattern:")
	pyramid.PrintPyramid()
}
