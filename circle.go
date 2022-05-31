package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (c Circle) CalcArea() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
